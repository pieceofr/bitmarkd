// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package p2pannounce

import (
	"encoding/hex"
	"sync"
	"time"

	"github.com/bitmark-inc/bitmarkd/avl"
	"github.com/bitmark-inc/bitmarkd/background"
	"github.com/bitmark-inc/bitmarkd/fault"
	"github.com/bitmark-inc/bitmarkd/util"
	"github.com/bitmark-inc/logger"
	proto "github.com/golang/protobuf/proto"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// type of listener
const (
	TypeRPC  = iota
	TypePeer = iota
)

// type for SHA3 fingerprints
type fingerprintType [32]byte

// RPC entries
type rpcEntry struct {
	address     util.PackedConnection // packed addresses
	fingerprint fingerprintType       // SHA3-256(certificate)
	timestamp   time.Time             // creation time
	local       bool                  // true => never expires
}

// globals for background process
type announcerData struct {
	sync.RWMutex // to allow locking

	// logger
	log *logger.L

	// this node's packed annoucements
	peerID      peerlib.ID
	listeners   []ma.Multiaddr
	fingerprint fingerprintType
	rpcs        []byte
	peerSet     bool
	rpcsSet     bool

	// tree of nodes available
	peerTree    *avl.Tree
	thisNode    *avl.Node // this node's position in the tree
	treeChanged bool      // tree was changed
	peerFile    string

	// database of all RPCs
	rpcIndex map[fingerprintType]int // index to find rpc entry
	rpcList  []*rpcEntry             // array of RPCs

	// data for thread
	ann announcer

	nodesLookup nodesLookup

	// for background
	background *background.T

	// set once during initialise
	initialised bool
}

// global data
var globalData announcerData

// format for timestamps
const timeFormat = "2006-01-02 15:04:05"

// Initialise - set up the announcement system
// pass a fully qualified domain for root node list
// or empty string for no root nodes
func Initialise(nodesDomain, peerFile string) error {

	globalData.Lock()
	defer globalData.Unlock()

	// no need to start if already started
	if globalData.initialised {
		return fault.AlreadyInitialised
	}

	globalData.log = logger.New("p2pannounce")
	globalData.log.Info("p2pannounce starting…")

	globalData.peerTree = avl.New()
	globalData.thisNode = nil
	globalData.treeChanged = false

	globalData.rpcIndex = make(map[fingerprintType]int, 1000)
	globalData.rpcList = make([]*rpcEntry, 0, 1000)

	globalData.peerSet = false
	globalData.rpcsSet = false
	globalData.peerFile = peerFile

	globalData.log.Info("start restoring peer data…")
	if _, err := restorePeers(globalData.peerFile); err != nil {

		globalData.log.Errorf("fail to restore peer data: %s", err.Error())
	}

	if err := globalData.nodesLookup.initialise(nodesDomain); nil != err {
		return err
	}

	if err := globalData.ann.initialise(); nil != err {
		return err
	}

	// all data initialised
	globalData.initialised = true

	// start background processes
	globalData.log.Info("start background…")

	processes := background.Processes{
		&globalData.nodesLookup, &globalData.ann,
	}

	globalData.background = background.Start(processes, globalData.log)

	return nil
}

// Finalise - stop all background tasks
func Finalise() error {

	if !globalData.initialised {
		return fault.NotInitialised
	}

	globalData.log.Info("shutting down…")
	globalData.log.Flush()

	// stop background
	globalData.background.Stop()

	globalData.log.Info("start backing up peer data…")
	if err := storePeers(globalData.peerFile); err != nil {
		globalData.log.Errorf("fail to backup peer data: %s", err.Error())
	}

	// finally...
	globalData.initialised = false

	globalData.log.Info("finished")
	globalData.log.Flush()

	return nil
}

// MarshalText - convert fingerprint to little endian hex text
func (fingerprint fingerprintType) MarshalText() ([]byte, error) {
	size := hex.EncodedLen(len(fingerprint))
	buffer := make([]byte, size)
	hex.Encode(buffer, fingerprint[:])
	return buffer, nil
}

func printBinaryAddrs(addrs []byte) string {
	maAddrs := Addrs{}
	err := proto.Unmarshal(addrs, &maAddrs)
	if err != nil {
		return ""
	}
	printAddrs := util.PrintMaAddrs(util.GetMultiAddrsFromBytes(maAddrs.Address))
	return printAddrs
}
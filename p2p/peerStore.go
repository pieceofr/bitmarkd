package p2p

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

//TODO: This function add address into the peer with the same id. Needs to take care of  IP changes
// addPeer to PeerStore
func (n *Node) addPeerAddrs(info peerlib.AddrInfo) {
	n.Lock()
	for _, addr := range info.Addrs {
		n.Host.Peerstore().AddAddr(info.ID, addr, peerstore.ConnectedAddrTTL)
		n.log.Infof("add peerstore:%s", info.String())
	}
	n.Unlock()
}

func (n *Node) addPeerAddr(id peer.ID, peerAddr ma.Multiaddr) {
	n.Lock()
	n.log.Infof("add peerstore:%s", id.String())
	n.Host.Peerstore().AddAddr(id, peerAddr, peerstore.ConnectedAddrTTL)
	n.Unlock()
}

func (n *Node) printPeerStore() {
	if len(n.Host.Peerstore().PeersWithAddrs()) == 0 {
		n.log.Warn("no peers in peerstore")
		return
	}
	for index, id := range n.Host.Peerstore().PeersWithAddrs() {
		infoAddrs := n.Host.Peerstore().Addrs(id)
		addrsString := ""
		for _, addr := range infoAddrs {
			addrsString = fmt.Sprintf("%s-%s", addrsString, addr)
		}
		n.log.Infof("Peerstore[%d]: ID:%v  time:%d \nAddrs:%s", index, id, time.Now().Unix(), addrsString)
	}
}
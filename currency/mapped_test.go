// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package currency_test

import (
	"encoding/json"
	"testing"

	"github.com/bitmark-inc/bitmarkd/currency"
	"github.com/bitmark-inc/bitmarkd/fault"
)

// test the currency.Map packer

func TestMapPack(t *testing.T) {

	testData := []struct {
		m currency.Map
		s currency.Set
		j string
	}{
		{
			m: currency.Map{
				currency.Bitcoin: "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			s: currency.MakeSet(currency.Bitcoin),
			j: `{"BTC":"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn"}`,
		},
		{
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
			},
			s: currency.MakeSet(currency.Litecoin),
			j: `{"LTC":"mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6"}`,
		},
		{
			m: currency.Map{
				currency.Bitcoin:  "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
			},
			s: currency.MakeSet(currency.Bitcoin, currency.Litecoin),
			j: `{"BTC":"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn","LTC":"mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6"}`,
		},
		{
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
				currency.Bitcoin:  "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			s: currency.MakeSet(currency.Bitcoin, currency.Litecoin),
			j: `{"BTC":"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn","LTC":"mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6"}`,
		},
	}

	for i, item := range testData {

		buffer, err := item.m.Pack(true)
		if nil != err {
			t.Fatalf("%d: pack error: %s", i, err)
		}

		mr, cs, err := currency.UnpackMap(buffer, true)
		if nil != err {
			t.Fatalf("%d: unpack error: %s", i, err)
		}

		for c, a := range item.m {
			if a != mr[c] {
				t.Fatalf("%d: unpack actual: %+v  expected: %+v", i, mr, item.m)
			}
		}
		if item.s != cs {
			t.Errorf("%d: actual set: %v  expected: %v", i, cs, item.s)
		}

		j, err := json.Marshal(item.m)
		if nil != err {
			t.Fatalf("%d: marshal JSON error: %s", i, err)
		}
		if string(j) != item.j {
			t.Errorf("%d: actual: `%s`  expected: %s", i, j, item.j)
		}
	}
}

func TestMapUnpackBuffer(t *testing.T) {

	testData := []struct {
		b []byte
		m currency.Map
		s currency.Set
	}{
		{
			b: []byte{
				0x02, 0x22, 0x6d, 0x6d, 0x43, 0x4b, 0x5a, 0x53,
				0x37, 0x74, 0x6f, 0x45, 0x36, 0x39, 0x51, 0x67,
				0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a, 0x63, 0x6a,
				0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38, 0x4c, 0x66,
				0x55, 0x62, 0x7a, 0x36,
			},
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
			},
			s: currency.MakeSet(currency.Litecoin),
		},
		{
			b: []byte{
				0x01, 0x22, 0x6d, 0x69, 0x70, 0x63, 0x42, 0x62,
				0x46, 0x67, 0x39, 0x67, 0x4d, 0x69, 0x43, 0x68,
				0x38, 0x31, 0x4b, 0x6a, 0x38, 0x74, 0x71, 0x71,
				0x64, 0x67, 0x6f, 0x5a, 0x75, 0x62, 0x31, 0x5a,
				0x4a, 0x52, 0x66, 0x6e,
			},
			m: currency.Map{
				currency.Bitcoin: "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			s: currency.MakeSet(currency.Bitcoin),
		},
		{
			b: []byte{
				0x02, 0x22, 0x6d, 0x6d, 0x43, 0x4b, 0x5a, 0x53,
				0x37, 0x74, 0x6f, 0x45, 0x36, 0x39, 0x51, 0x67,
				0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a, 0x63, 0x6a,
				0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38, 0x4c, 0x66,
				0x55, 0x62, 0x7a, 0x36, 0x01, 0x22, 0x6d, 0x69,
				0x70, 0x63, 0x42, 0x62, 0x46, 0x67, 0x39, 0x67,
				0x4d, 0x69, 0x43, 0x68, 0x38, 0x31, 0x4b, 0x6a,
				0x38, 0x74, 0x71, 0x71, 0x64, 0x67, 0x6f, 0x5a,
				0x75, 0x62, 0x31, 0x5a, 0x4a, 0x52, 0x66, 0x6e,
			},
			m: currency.Map{
				currency.Bitcoin:  "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
			},
			s: currency.MakeSet(currency.Bitcoin, currency.Litecoin),
		},

		{
			b: []byte{
				0x01, 0x22, 0x6d, 0x69, 0x70, 0x63, 0x42, 0x62,
				0x46, 0x67, 0x39, 0x67, 0x4d, 0x69, 0x43, 0x68,
				0x38, 0x31, 0x4b, 0x6a, 0x38, 0x74, 0x71, 0x71,
				0x64, 0x67, 0x6f, 0x5a, 0x75, 0x62, 0x31, 0x5a,
				0x4a, 0x52, 0x66, 0x6e, 0x02, 0x22, 0x6d, 0x6d,
				0x43, 0x4b, 0x5a, 0x53, 0x37, 0x74, 0x6f, 0x45,
				0x36, 0x39, 0x51, 0x67, 0x58, 0x4e, 0x73, 0x31,
				0x4a, 0x5a, 0x63, 0x6a, 0x57, 0x36, 0x4c, 0x46,
				0x6a, 0x38, 0x4c, 0x66, 0x55, 0x62, 0x7a, 0x36,
			},
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
				currency.Bitcoin:  "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			s: currency.MakeSet(currency.Bitcoin, currency.Litecoin),
		},
	}

	for i, item := range testData {

		mr, cs, err := currency.UnpackMap(item.b, true)
		if nil != err {
			t.Fatalf("%d: unpack error: %s", i, err)
		}

		for c, a := range item.m {
			if a != mr[c] {
				t.Fatalf("%d: unpack actual: %+v  expected: %+v", i, mr, item.m)
			}
		}
		if item.s != cs {
			t.Logf("%d: actual set: %v  expected: %v", i, cs, item.s)
		}
	}
}

func TestMapPackInvalid(t *testing.T) {

	testData := []struct {
		m   currency.Map
		err error
	}{
		{
			m: currency.Map{
				currency.Nothing: "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			err: fault.ErrInvalidCurrency,
		},
		{
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUbz6",
				currency.Bitcoin:  "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
				currency.Last + 1: "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			},
			err: fault.ErrInvalidCurrency,
		},
		{
			m: currency.Map{
				currency.Litecoin: "mmCKZS7toE69QgXNs1JZcjW6LFj8LfUb6z",
			},
			err: fault.ErrInvalidLitecoinAddress,
		},
	}

	for i, item := range testData {

		_, err := item.m.Pack(true)
		if item.err != err {
			t.Fatalf("%d: error: %s  expected: %s", i, err, item.err)
		}
	}
}

func TestMapUnpackInvalid(t *testing.T) {

	testData := []struct {
		b   []byte
		err error
	}{
		{
			b: []byte{
				0x00, 0x22, 0x6d, 0x6d, 0x43, 0x4b, 0x5a, 0x53,
				0x37, 0x74, 0x6f, 0x45, 0x36, 0x39, 0x51, 0x67,
				0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a, 0x63, 0x6a,
				0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38, 0x4c, 0x66,
				0x55, 0x62, 0x7a, 0x36, 0x01, 0x22, 0x6d, 0x69,
				0x70, 0x63, 0x42, 0x62, 0x46, 0x67, 0x39, 0x67,
				0x4d, 0x69, 0x43, 0x68, 0x38, 0x31, 0x4b, 0x6a,
				0x38, 0x74, 0x71, 0x71, 0x64, 0x67, 0x6f, 0x5a,
				0x75, 0x62, 0x31, 0x5a, 0x4a, 0x52, 0x66, 0x6e,
			},
			err: fault.ErrInvalidCurrency,
		},
		{
			b: []byte{
				0xff, 0xff, 0x7f, 0x22, 0x6d, 0x6d, 0x43, 0x4b,
				0x5a, 0x53, 0x37, 0x74, 0x6f, 0x45, 0x36, 0x39,
				0x51, 0x67, 0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a,
				0x63, 0x6a, 0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38,
				0x4c, 0x66, 0x55, 0x62, 0x7a, 0x36, 0x01, 0x22,
				0x6d, 0x69, 0x70, 0x63, 0x42, 0x62, 0x46, 0x67,
				0x39, 0x67, 0x4d, 0x69, 0x43, 0x68, 0x38, 0x31,
				0x4b, 0x6a, 0x38, 0x74, 0x71, 0x71, 0x64, 0x67,
				0x6f, 0x5a, 0x75, 0x62, 0x31, 0x5a, 0x4a, 0x52,
				0x66, 0x6e,
			},
			err: fault.ErrInvalidCurrency,
		},
		{
			b: []byte{
				0x02, 0x22, 0x6d, 0x6d, 0x43, 0x4b, 0x5a, 0x53,
				0x37, 0x74, 0x6f, 0x45, 0x36, 0x39, 0x51, 0x67,
				0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a, 0x63, 0x6a,
				0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38, 0x4c, 0x66,
				0x55, 0x62, 0x36, 0x71, 0x01, 0x22, 0x6d, 0x69,
				0x70, 0x63, 0x42, 0x62, 0x46, 0x67, 0x39, 0x67,
				0x4d, 0x69, 0x43, 0x68, 0x38, 0x31, 0x4b, 0x6a,
				0x38, 0x74, 0x71, 0x71, 0x64, 0x67, 0x6f, 0x5a,
				0x75, 0x62, 0x31, 0x5a, 0x4a, 0x52, 0x66, 0x6e,
			},
			err: fault.ErrInvalidLitecoinAddress,
		},
		{
			b: []byte{
				0x02, 0xff, 0x7f, 0x6d, 0x6d, 0x43, 0x4b, 0x5a,
				0x53, 0x37, 0x74, 0x6f, 0x45, 0x36, 0x39, 0x51,
				0x67, 0x58, 0x4e, 0x73, 0x31, 0x4a, 0x5a, 0x63,
				0x6a, 0x57, 0x36, 0x4c, 0x46, 0x6a, 0x38, 0x4c,
				0x66, 0x55, 0x62, 0x36, 0x71, 0x01, 0x22, 0x6d,
				0x69, 0x70, 0x63, 0x42, 0x62, 0x46, 0x67, 0x39,
				0x67, 0x4d, 0x69, 0x43, 0x68, 0x38, 0x31, 0x4b,
				0x6a, 0x38, 0x74, 0x71, 0x71, 0x64, 0x67, 0x6f,
				0x5a, 0x75, 0x62, 0x31, 0x5a, 0x4a, 0x52, 0x66,
				0x6e,
			},
			err: fault.ErrInvalidCount,
		},
	}

	for i, item := range testData {

		_, _, err := currency.UnpackMap(item.b, true)
		if item.err != err {
			t.Fatalf("%d: error: %s  expected: %s", i, err, item.err)
		}
	}
}

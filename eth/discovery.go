// Copyright 2019 The eutherum Authors
// This file is part of the eutherum library.
//
// The eutherum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The eutherum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the eutherum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"github.com/eutherum/eutherum/core"
	"github.com/eutherum/eutherum/core/forkid"
	"github.com/eutherum/eutherum/p2p/enode"
	"github.com/eutherum/eutherum/rlp"
)

// ethEntry is the "eth" ENR entry which advertises eth protocol
// on the discovery network.
type ethEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e ethEntry) ENRKey() string {
	return "eth"
}

// startEthEntryUpdate starts the ENR updater loop.
func (eth *Eutherum) startEthEntryUpdate(ln *enode.LocalNode) {
	var newHead = make(chan core.ChainHeadEvent, 10)
	sub := eth.blockchain.SubscribeChainHeadEvent(newHead)

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case <-newHead:
				ln.Set(eth.currentEthEntry())
			case <-sub.Err():
				// Would be nice to sync with eth.Stop, but there is no
				// good way to do that.
				return
			}
		}
	}()
}

func (eth *Eutherum) currentEthEntry() *ethEntry {
	return &ethEntry{ForkID: forkid.NewID(eth.blockchain.Config(), eth.blockchain.Genesis().Hash(),
		eth.blockchain.CurrentHeader().Number.Uint64())}
}

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

// Package checkpointeuracle is a an on-chain light client checkpoint euracle.
package checkpointeuracle

//go:generate abigen --sol contract/euracle.sol --pkg contract --out contract/euracle.go

import (
	"errors"
	"math/big"

	"github.com/eutherum/eutherum/accounts/abi/bind"
	"github.com/eutherum/eutherum/common"
	"github.com/eutherum/eutherum/contracts/checkpointeuracle/contract"
	"github.com/eutherum/eutherum/core/types"
)

// CheckpointEuracle is a Go wrapper around an on-chain checkpoint euracle contract.
type CheckpointEuracle struct {
	address  common.Address
	contract *contract.CheckpointEuracle
}

// NewCheckpointEuracle binds checkpoint contract and returns a registrar instance.
func NewCheckpointEuracle(contractAddr common.Address, backend bind.ContractBackend) (*CheckpointEuracle, error) {
	c, err := contract.NewCheckpointEuracle(contractAddr, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointEuracle{address: contractAddr, contract: c}, nil
}

// ContractAddr returns the address of contract.
func (euracle *CheckpointEuracle) ContractAddr() common.Address {
	return euracle.address
}

// Contract returns the underlying contract instance.
func (euracle *CheckpointEuracle) Contract() *contract.CheckpointEuracle {
	return euracle.contract
}

// LookupCheckpointEvents searches checkpoint event for specific section in the
// given log batches.
func (euracle *CheckpointEuracle) LookupCheckpointEvents(blockLogs [][]*types.Log, section uint64, hash common.Hash) []*contract.CheckpointEuracleNewCheckpointVote {
	var votes []*contract.CheckpointEuracleNewCheckpointVote

	for _, logs := range blockLogs {
		for _, log := range logs {
			event, err := euracle.contract.ParseNewCheckpointVote(*log)
			if err != nil {
				continue
			}
			if event.Index == section && event.CheckpointHash == hash {
				votes = append(votes, event)
			}
		}
	}
	return votes
}

// RegisterCheckpoint registers the checkpoint with a batch of associated signatures
// that are collected off-chain and sorted by lexicographical order.
//
// Notably all signatures given should be transformed to "eutherum style" which transforms
// v from 0/1 to 27/28 according to the yellow paper.
func (euracle *CheckpointEuracle) RegisterCheckpoint(opts *bind.TransactOpts, index uint64, hash []byte, rnum *big.Int, rhash [32]byte, sigs [][]byte) (*types.Transaction, error) {
	var (
		r [][32]byte
		s [][32]byte
		v []uint8
	)
	for i := 0; i < len(sigs); i++ {
		if len(sigs[i]) != 65 {
			return nil, errors.New("invalid signature")
		}
		r = append(r, common.BytesToHash(sigs[i][:32]))
		s = append(s, common.BytesToHash(sigs[i][32:64]))
		v = append(v, sigs[i][64])
	}
	return euracle.contract.SetCheckpoint(opts, rnum, rhash, common.BytesToHash(hash), index, v, r, s)
}

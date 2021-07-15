// Copyright 2019 The Tomochain Authors
// Copyright (c) 2021 Sdxchain
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package posv

import (
	"math/big"

	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/consensus"
	"github.com/69th-byte/SmartDex-Chain/core/types"
	"github.com/69th-byte/SmartDex-Chain/rpc"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain consensus.ChainReader
	posv  *Posv
}
type NetworkInformation struct {
	NetworkId                  *big.Int
	SdxValidatorAddress        common.Address
	RelayerRegistrationAddress common.Address
	SdxXListingAddress         common.Address
	SdxZAddress                common.Address
	LendingAddress             common.Address
}

// GetSnapshot retrieves the state snapshot at a given block.
func (api *API) GetSnapshot(number *rpc.BlockNumber) (*Snapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.posv.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.posv.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.posv.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.GetSigners(), nil
}

// GetSignersAtHash retrieves the state snapshot at a given block.
func (api *API) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.posv.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil)
	if err != nil {
		return nil, err
	}
	return snap.GetSigners(), nil
}

func (api *API) NetworkInformation() NetworkInformation {
	api.posv.lock.RLock()
	defer api.posv.lock.RUnlock()
	info := NetworkInformation{}
	info.NetworkId = api.chain.Config().ChainId
	info.SdxValidatorAddress = common.HexToAddress(common.MasternodeVotingSMC)
	if common.IsTestnet {
		info.LendingAddress = common.HexToAddress(common.LendingRegistrationSMCTestnet)
		info.RelayerRegistrationAddress = common.HexToAddress(common.RelayerRegistrationSMCTestnet)
		info.SdxXListingAddress = common.SdxXListingSMCTestNet
		info.SdxZAddress = common.SRC21IssuerSMCTestNet
	} else {
		info.LendingAddress = common.HexToAddress(common.LendingRegistrationSMC)
		info.RelayerRegistrationAddress = common.HexToAddress(common.RelayerRegistrationSMC)
		info.SdxXListingAddress = common.SdxXListingSMC
		info.SdxZAddress = common.SRC21IssuerSMC
	}
	return info
}

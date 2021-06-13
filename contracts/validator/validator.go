// Copyright (c) 2021 SDXchain
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

package validator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/validator/contract"
)

type Validator struct {
	*contract.SDXValidatorSession
	contractBackend bind.ContractBackend
}

func NewValidator(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Validator, error) {
	validator, err := contract.NewSDXValidator(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Validator{
		&contract.SDXValidatorSession{
			Contract:     validator,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployValidator(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, validatorAddress []common.Address, caps []*big.Int, ownerAddress common.Address) (common.Address, *Validator, error) {
	minDeposit := new(big.Int)
	minDeposit.SetString("10000000000000000000000", 10)
	minVoterCap := new(big.Int)
	minVoterCap.SetString("10000000000000000000", 10)
	// Deposit 10K SDX
	// Min Voter Cap 10 SDX
	// 500 masternodes
	// Candidate Delay Withdraw 30 days = 1296000 blocks
	// Voter Delay Withdraw 10 days = 432000 blocks
	validatorAddr, _, _, err := contract.DeploySDXValidator(transactOpts, contractBackend, validatorAddress, caps, ownerAddress, minDeposit, minVoterCap, big.NewInt(1500), big.NewInt(1296000), big.NewInt(432000))
	if err != nil {
		return validatorAddr, nil, err
	}

	validator, err := NewValidator(transactOpts, validatorAddr, contractBackend)
	if err != nil {
		return validatorAddr, nil, err
	}

	return validatorAddr, validator, nil
}

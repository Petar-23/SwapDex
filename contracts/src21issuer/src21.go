package src21issuer

import (
	"math/big"

	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/contracts/src21issuer/contract"
)

type MySRC21 struct {
	*contract.MySRC21Session
	contractBackend bind.ContractBackend
}

func NewSRC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MySRC21, error) {
	smartContract, err := contract.NewMySRC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MySRC21{
		&contract.MySRC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeploySRC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, name string, symbol string, decimals uint8, cap, fee *big.Int) (common.Address, *MySRC21, error) {
	contractAddr, _, _, err := contract.DeployMySRC21(transactOpts, contractBackend, name, symbol, decimals, cap, fee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewSRC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}

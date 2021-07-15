package sdxx

import (
	"math/big"

	"github.com/69th-byte/SmartDex-Chain/accounts/abi/bind"
	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/contracts/sdxx/contract"
)

type SRC21Issuer struct {
	*contract.SRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewSRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*SRC21Issuer, error) {
	contractObject, err := contract.NewSRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &SRC21Issuer{
		&contract.SRC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeploySRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *SRC21Issuer, error) {
	contractAddr, _, _, err := contract.DeploySRC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewSRC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}

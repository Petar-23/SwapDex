package sdxx

import (
	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/contracts/sdxx/contract"
)

type SDXXListing struct {
	*contract.SDXXListingSession
	contractBackend bind.ContractBackend
}

func NewMySDXXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*SDXXListing, error) {
	smartContract, err := contract.NewSDXXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &SDXXListing{
		&contract.SDXXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeploySDXXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *SDXXListing, error) {
	contractAddr, _, _, err := contract.DeploySDXXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMySDXXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}

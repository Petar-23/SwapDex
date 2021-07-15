// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/69th-byte/SmartDex-Chain/accounts/abi"
	"github.com/69th-byte/SmartDex-Chain/accounts/abi/bind"
	"github.com/69th-byte/SmartDex-Chain/common"
	"github.com/69th-byte/SmartDex-Chain/core/types"
)

// SDXXListingABI is the input ABI used to generate the binding from.
const SDXXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// SDXXListingBin is the compiled bytecode used for deploying new contracts.
const SDXXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeploySDXXListing deploys a new Ethereum contract, binding an instance of SDXXListing to it.
func DeploySDXXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SDXXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(SDXXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SDXXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SDXXListing{SDXXListingCaller: SDXXListingCaller{contract: contract}, SDXXListingTransactor: SDXXListingTransactor{contract: contract}, SDXXListingFilterer: SDXXListingFilterer{contract: contract}}, nil
}

// SDXXListing is an auto generated Go binding around an Ethereum contract.
type SDXXListing struct {
	SDXXListingCaller     // Read-only binding to the contract
	SDXXListingTransactor // Write-only binding to the contract
	SDXXListingFilterer   // Log filterer for contract events
}

// SDXXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SDXXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDXXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SDXXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDXXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SDXXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SDXXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SDXXListingSession struct {
	Contract     *SDXXListing      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SDXXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SDXXListingCallerSession struct {
	Contract *SDXXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SDXXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SDXXListingTransactorSession struct {
	Contract     *SDXXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SDXXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SDXXListingRaw struct {
	Contract *SDXXListing // Generic contract binding to access the raw methods on
}

// SDXXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SDXXListingCallerRaw struct {
	Contract *SDXXListingCaller // Generic read-only contract binding to access the raw methods on
}

// SDXXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SDXXListingTransactorRaw struct {
	Contract *SDXXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSDXXListing creates a new instance of SDXXListing, bound to a specific deployed contract.
func NewSDXXListing(address common.Address, backend bind.ContractBackend) (*SDXXListing, error) {
	contract, err := bindSDXXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SDXXListing{SDXXListingCaller: SDXXListingCaller{contract: contract}, SDXXListingTransactor: SDXXListingTransactor{contract: contract}, SDXXListingFilterer: SDXXListingFilterer{contract: contract}}, nil
}

// NewSDXXListingCaller creates a new read-only instance of SDXXListing, bound to a specific deployed contract.
func NewSDXXListingCaller(address common.Address, caller bind.ContractCaller) (*SDXXListingCaller, error) {
	contract, err := bindSDXXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SDXXListingCaller{contract: contract}, nil
}

// NewSDXXListingTransactor creates a new write-only instance of SDXXListing, bound to a specific deployed contract.
func NewSDXXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*SDXXListingTransactor, error) {
	contract, err := bindSDXXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SDXXListingTransactor{contract: contract}, nil
}

// NewSDXXListingFilterer creates a new log filterer instance of SDXXListing, bound to a specific deployed contract.
func NewSDXXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*SDXXListingFilterer, error) {
	contract, err := bindSDXXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SDXXListingFilterer{contract: contract}, nil
}

// bindSDXXListing binds a generic wrapper to an already deployed contract.
func bindSDXXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SDXXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDXXListing *SDXXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SDXXListing.Contract.SDXXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDXXListing *SDXXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDXXListing.Contract.SDXXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDXXListing *SDXXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDXXListing.Contract.SDXXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SDXXListing *SDXXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SDXXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SDXXListing *SDXXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SDXXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SDXXListing *SDXXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SDXXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_SDXXListing *SDXXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SDXXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_SDXXListing *SDXXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _SDXXListing.Contract.GetTokenStatus(&_SDXXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_SDXXListing *SDXXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _SDXXListing.Contract.GetTokenStatus(&_SDXXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_SDXXListing *SDXXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _SDXXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_SDXXListing *SDXXListingSession) Tokens() ([]common.Address, error) {
	return _SDXXListing.Contract.Tokens(&_SDXXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_SDXXListing *SDXXListingCallerSession) Tokens() ([]common.Address, error) {
	return _SDXXListing.Contract.Tokens(&_SDXXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_SDXXListing *SDXXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SDXXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_SDXXListing *SDXXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _SDXXListing.Contract.Apply(&_SDXXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_SDXXListing *SDXXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _SDXXListing.Contract.Apply(&_SDXXListing.TransactOpts, token)
}

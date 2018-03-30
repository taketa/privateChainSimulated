// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SimpleStorageABI is the input ABI used to generate the binding from.
const SimpleStorageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"changeName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"changeAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAge\",\"outputs\":[{\"name\":\"_age\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimpleStorageBin is the compiled bytecode used for deploying new contracts.
const SimpleStorageBin = `0x6060604052341561000f57600080fd5b6103cd8061001e6000396000f3006060604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317d7de7c81146100665780635353a2d8146100f05780635e439f6614610143578063967e6e6514610159575b600080fd5b341561007157600080fd5b61007961017e565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100b557808201518382015260200161009d565b50505050905090810190601f1680156100e25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100fb57600080fd5b61014160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061026695505050505050565b005b341561014e57600080fd5b6101416004356102a0565b341561016457600080fd5b61016c6102cb565b60405190815260200160405180910390f35b6101866102f7565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561025b5780601f106102305761010080835404028352916020019161025b565b820191906000526020600020905b81548152906001019060200180831161023e57829003601f168201915b505050505090505b90565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208190526040902081805161029c929160200190610309565b5050565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260208190526040902060010155565b73ffffffffffffffffffffffffffffffffffffffff331660009081526020819052604090206001015490565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061034a57805160ff1916838001178555610377565b82800160010185558215610377579182015b8281111561037757825182559160200191906001019061035c565b50610383929150610387565b5090565b61026391905b80821115610383576000815560010161038d5600a165627a7a72305820675c900b2498c944e21c89cd18bc07128eb861df6de2bbb3c68060dad91431550029`

// DeploySimpleStorage deploys a new Ethereum contract, binding an instance of SimpleStorage to it.
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimpleStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// GetAge is a free data retrieval call binding the contract method 0x967e6e65.
//
// Solidity: function getAge() constant returns(_age uint256)
func (_SimpleStorage *SimpleStorageCaller) GetAge(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimpleStorage.contract.Call(opts, out, "getAge")
	return *ret0, err
}

// GetAge is a free data retrieval call binding the contract method 0x967e6e65.
//
// Solidity: function getAge() constant returns(_age uint256)
func (_SimpleStorage *SimpleStorageSession) GetAge() (*big.Int, error) {
	return _SimpleStorage.Contract.GetAge(&_SimpleStorage.CallOpts)
}

// GetAge is a free data retrieval call binding the contract method 0x967e6e65.
//
// Solidity: function getAge() constant returns(_age uint256)
func (_SimpleStorage *SimpleStorageCallerSession) GetAge() (*big.Int, error) {
	return _SimpleStorage.Contract.GetAge(&_SimpleStorage.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(_name string)
func (_SimpleStorage *SimpleStorageCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SimpleStorage.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(_name string)
func (_SimpleStorage *SimpleStorageSession) GetName() (string, error) {
	return _SimpleStorage.Contract.GetName(&_SimpleStorage.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(_name string)
func (_SimpleStorage *SimpleStorageCallerSession) GetName() (string, error) {
	return _SimpleStorage.Contract.GetName(&_SimpleStorage.CallOpts)
}

// ChangeAge is a paid mutator transaction binding the contract method 0x5e439f66.
//
// Solidity: function changeAge(_age uint256) returns()
func (_SimpleStorage *SimpleStorageTransactor) ChangeAge(opts *bind.TransactOpts, _age *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "changeAge", _age)
}

// ChangeAge is a paid mutator transaction binding the contract method 0x5e439f66.
//
// Solidity: function changeAge(_age uint256) returns()
func (_SimpleStorage *SimpleStorageSession) ChangeAge(_age *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.ChangeAge(&_SimpleStorage.TransactOpts, _age)
}

// ChangeAge is a paid mutator transaction binding the contract method 0x5e439f66.
//
// Solidity: function changeAge(_age uint256) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) ChangeAge(_age *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.ChangeAge(&_SimpleStorage.TransactOpts, _age)
}

// ChangeName is a paid mutator transaction binding the contract method 0x5353a2d8.
//
// Solidity: function changeName(_name string) returns()
func (_SimpleStorage *SimpleStorageTransactor) ChangeName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "changeName", _name)
}

// ChangeName is a paid mutator transaction binding the contract method 0x5353a2d8.
//
// Solidity: function changeName(_name string) returns()
func (_SimpleStorage *SimpleStorageSession) ChangeName(_name string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.ChangeName(&_SimpleStorage.TransactOpts, _name)
}

// ChangeName is a paid mutator transaction binding the contract method 0x5353a2d8.
//
// Solidity: function changeName(_name string) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) ChangeName(_name string) (*types.Transaction, error) {
	return _SimpleStorage.Contract.ChangeName(&_SimpleStorage.TransactOpts, _name)
}

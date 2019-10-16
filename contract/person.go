// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PersonABI is the input ABI used to generate the binding from.
const PersonABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Person is an auto generated Go binding around an Ethereum contract.
type Person struct {
	PersonCaller     // Read-only binding to the contract
	PersonTransactor // Write-only binding to the contract
	PersonFilterer   // Log filterer for contract events
}

// PersonCaller is an auto generated read-only Go binding around an Ethereum contract.
type PersonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PersonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PersonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PersonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PersonSession struct {
	Contract     *Person           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PersonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PersonCallerSession struct {
	Contract *PersonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PersonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PersonTransactorSession struct {
	Contract     *PersonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PersonRaw is an auto generated low-level Go binding around an Ethereum contract.
type PersonRaw struct {
	Contract *Person // Generic contract binding to access the raw methods on
}

// PersonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PersonCallerRaw struct {
	Contract *PersonCaller // Generic read-only contract binding to access the raw methods on
}

// PersonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PersonTransactorRaw struct {
	Contract *PersonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerson creates a new instance of Person, bound to a specific deployed contract.
func NewPerson(address common.Address, backend bind.ContractBackend) (*Person, error) {
	contract, err := bindPerson(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Person{PersonCaller: PersonCaller{contract: contract}, PersonTransactor: PersonTransactor{contract: contract}, PersonFilterer: PersonFilterer{contract: contract}}, nil
}

// NewPersonCaller creates a new read-only instance of Person, bound to a specific deployed contract.
func NewPersonCaller(address common.Address, caller bind.ContractCaller) (*PersonCaller, error) {
	contract, err := bindPerson(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PersonCaller{contract: contract}, nil
}

// NewPersonTransactor creates a new write-only instance of Person, bound to a specific deployed contract.
func NewPersonTransactor(address common.Address, transactor bind.ContractTransactor) (*PersonTransactor, error) {
	contract, err := bindPerson(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PersonTransactor{contract: contract}, nil
}

// NewPersonFilterer creates a new log filterer instance of Person, bound to a specific deployed contract.
func NewPersonFilterer(address common.Address, filterer bind.ContractFilterer) (*PersonFilterer, error) {
	contract, err := bindPerson(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PersonFilterer{contract: contract}, nil
}

// bindPerson binds a generic wrapper to an already deployed contract.
func bindPerson(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PersonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Person *PersonRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Person.Contract.PersonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Person *PersonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Person.Contract.PersonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Person *PersonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Person.Contract.PersonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Person *PersonCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Person.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Person *PersonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Person.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Person *PersonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Person.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Person *PersonCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Person.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Person *PersonSession) Name() (string, error) {
	return _Person.Contract.Name(&_Person.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Person *PersonCallerSession) Name() (string, error) {
	return _Person.Contract.Name(&_Person.CallOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Person *PersonTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Person.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Person *PersonSession) SetName(_name string) (*types.Transaction, error) {
	return _Person.Contract.SetName(&_Person.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Person *PersonTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Person.Contract.SetName(&_Person.TransactOpts, _name)
}

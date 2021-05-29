// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demo

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DemoABI is the input ABI used to generate the binding from.
const DemoABI = "[{\"inputs\":[],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DemoBin is the compiled bytecode used for deploying new contracts.
var DemoBin = "0x60806040526000805534801561001457600080fd5b5061010e806100246000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80634f2be91f14602d575b600080fd5b60336035565b005b5b600a60005410156059576000808154809291906050906065565b91905055506036565b565b6000819050919050565b6000606e82605b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415609e57609d60a9565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220207011a882081d2a6f05adaf6df1d3aa8cd9a53c076433c026011d9e035ab6d564736f6c63430008040033"

// DeployDemo deploys a new Ethereum contract, binding an instance of Demo to it.
func DeployDemo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Demo, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DemoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Demo{DemoCaller: DemoCaller{contract: contract}, DemoTransactor: DemoTransactor{contract: contract}, DemoFilterer: DemoFilterer{contract: contract}}, nil
}

// Demo is an auto generated Go binding around an Ethereum contract.
type Demo struct {
	DemoCaller     // Read-only binding to the contract
	DemoTransactor // Write-only binding to the contract
	DemoFilterer   // Log filterer for contract events
}

// DemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DemoSession struct {
	Contract     *Demo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DemoCallerSession struct {
	Contract *DemoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DemoTransactorSession struct {
	Contract     *DemoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DemoRaw struct {
	Contract *Demo // Generic contract binding to access the raw methods on
}

// DemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DemoCallerRaw struct {
	Contract *DemoCaller // Generic read-only contract binding to access the raw methods on
}

// DemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DemoTransactorRaw struct {
	Contract *DemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDemo creates a new instance of Demo, bound to a specific deployed contract.
func NewDemo(address common.Address, backend bind.ContractBackend) (*Demo, error) {
	contract, err := bindDemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Demo{DemoCaller: DemoCaller{contract: contract}, DemoTransactor: DemoTransactor{contract: contract}, DemoFilterer: DemoFilterer{contract: contract}}, nil
}

// NewDemoCaller creates a new read-only instance of Demo, bound to a specific deployed contract.
func NewDemoCaller(address common.Address, caller bind.ContractCaller) (*DemoCaller, error) {
	contract, err := bindDemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DemoCaller{contract: contract}, nil
}

// NewDemoTransactor creates a new write-only instance of Demo, bound to a specific deployed contract.
func NewDemoTransactor(address common.Address, transactor bind.ContractTransactor) (*DemoTransactor, error) {
	contract, err := bindDemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DemoTransactor{contract: contract}, nil
}

// NewDemoFilterer creates a new log filterer instance of Demo, bound to a specific deployed contract.
func NewDemoFilterer(address common.Address, filterer bind.ContractFilterer) (*DemoFilterer, error) {
	contract, err := bindDemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DemoFilterer{contract: contract}, nil
}

// bindDemo binds a generic wrapper to an already deployed contract.
func bindDemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demo *DemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demo.Contract.DemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demo *DemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demo.Contract.DemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demo *DemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demo.Contract.DemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Demo *DemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Demo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Demo *DemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Demo *DemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Demo.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Demo *DemoTransactor) Add(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Demo.contract.Transact(opts, "add")
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Demo *DemoSession) Add() (*types.Transaction, error) {
	return _Demo.Contract.Add(&_Demo.TransactOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4f2be91f.
//
// Solidity: function add() returns()
func (_Demo *DemoTransactorSession) Add() (*types.Transaction, error) {
	return _Demo.Contract.Add(&_Demo.TransactOpts)
}

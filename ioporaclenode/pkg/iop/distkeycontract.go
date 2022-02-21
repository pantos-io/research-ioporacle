// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iop

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DistKeyContractMetaData contains all meta data concerning the DistKeyContract contract.
var DistKeyContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DistKeyGenerationLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KEY_FINAL_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEY_GEN_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_publicKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfValidators\",\"type\":\"uint256\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"name\":\"setRegistryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DistKeyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DistKeyContractMetaData.ABI instead.
var DistKeyContractABI = DistKeyContractMetaData.ABI

// DistKeyContract is an auto generated Go binding around an Ethereum contract.
type DistKeyContract struct {
	DistKeyContractCaller     // Read-only binding to the contract
	DistKeyContractTransactor // Write-only binding to the contract
	DistKeyContractFilterer   // Log filterer for contract events
}

// DistKeyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistKeyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistKeyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistKeyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistKeyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistKeyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistKeyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistKeyContractSession struct {
	Contract     *DistKeyContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistKeyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistKeyContractCallerSession struct {
	Contract *DistKeyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DistKeyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistKeyContractTransactorSession struct {
	Contract     *DistKeyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DistKeyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistKeyContractRaw struct {
	Contract *DistKeyContract // Generic contract binding to access the raw methods on
}

// DistKeyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistKeyContractCallerRaw struct {
	Contract *DistKeyContractCaller // Generic read-only contract binding to access the raw methods on
}

// DistKeyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistKeyContractTransactorRaw struct {
	Contract *DistKeyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistKeyContract creates a new instance of DistKeyContract, bound to a specific deployed contract.
func NewDistKeyContract(address common.Address, backend bind.ContractBackend) (*DistKeyContract, error) {
	contract, err := bindDistKeyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistKeyContract{DistKeyContractCaller: DistKeyContractCaller{contract: contract}, DistKeyContractTransactor: DistKeyContractTransactor{contract: contract}, DistKeyContractFilterer: DistKeyContractFilterer{contract: contract}}, nil
}

// NewDistKeyContractCaller creates a new read-only instance of DistKeyContract, bound to a specific deployed contract.
func NewDistKeyContractCaller(address common.Address, caller bind.ContractCaller) (*DistKeyContractCaller, error) {
	contract, err := bindDistKeyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistKeyContractCaller{contract: contract}, nil
}

// NewDistKeyContractTransactor creates a new write-only instance of DistKeyContract, bound to a specific deployed contract.
func NewDistKeyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DistKeyContractTransactor, error) {
	contract, err := bindDistKeyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistKeyContractTransactor{contract: contract}, nil
}

// NewDistKeyContractFilterer creates a new log filterer instance of DistKeyContract, bound to a specific deployed contract.
func NewDistKeyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DistKeyContractFilterer, error) {
	contract, err := bindDistKeyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistKeyContractFilterer{contract: contract}, nil
}

// bindDistKeyContract binds a generic wrapper to an already deployed contract.
func bindDistKeyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistKeyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistKeyContract *DistKeyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistKeyContract.Contract.DistKeyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistKeyContract *DistKeyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistKeyContract.Contract.DistKeyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistKeyContract *DistKeyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistKeyContract.Contract.DistKeyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistKeyContract *DistKeyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistKeyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistKeyContract *DistKeyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistKeyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistKeyContract *DistKeyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistKeyContract.Contract.contract.Transact(opts, method, params...)
}

// KEYFINALTIME is a free data retrieval call binding the contract method 0xc09cd231.
//
// Solidity: function KEY_FINAL_TIME() view returns(uint256)
func (_DistKeyContract *DistKeyContractCaller) KEYFINALTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistKeyContract.contract.Call(opts, &out, "KEY_FINAL_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KEYFINALTIME is a free data retrieval call binding the contract method 0xc09cd231.
//
// Solidity: function KEY_FINAL_TIME() view returns(uint256)
func (_DistKeyContract *DistKeyContractSession) KEYFINALTIME() (*big.Int, error) {
	return _DistKeyContract.Contract.KEYFINALTIME(&_DistKeyContract.CallOpts)
}

// KEYFINALTIME is a free data retrieval call binding the contract method 0xc09cd231.
//
// Solidity: function KEY_FINAL_TIME() view returns(uint256)
func (_DistKeyContract *DistKeyContractCallerSession) KEYFINALTIME() (*big.Int, error) {
	return _DistKeyContract.Contract.KEYFINALTIME(&_DistKeyContract.CallOpts)
}

// KEYGENINTERVAL is a free data retrieval call binding the contract method 0xfcb0801e.
//
// Solidity: function KEY_GEN_INTERVAL() view returns(uint256)
func (_DistKeyContract *DistKeyContractCaller) KEYGENINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistKeyContract.contract.Call(opts, &out, "KEY_GEN_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KEYGENINTERVAL is a free data retrieval call binding the contract method 0xfcb0801e.
//
// Solidity: function KEY_GEN_INTERVAL() view returns(uint256)
func (_DistKeyContract *DistKeyContractSession) KEYGENINTERVAL() (*big.Int, error) {
	return _DistKeyContract.Contract.KEYGENINTERVAL(&_DistKeyContract.CallOpts)
}

// KEYGENINTERVAL is a free data retrieval call binding the contract method 0xfcb0801e.
//
// Solidity: function KEY_GEN_INTERVAL() view returns(uint256)
func (_DistKeyContract *DistKeyContractCallerSession) KEYGENINTERVAL() (*big.Int, error) {
	return _DistKeyContract.Contract.KEYGENINTERVAL(&_DistKeyContract.CallOpts)
}

// GetNumberOfValidators is a free data retrieval call binding the contract method 0x9031d913.
//
// Solidity: function getNumberOfValidators() view returns(uint256)
func (_DistKeyContract *DistKeyContractCaller) GetNumberOfValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistKeyContract.contract.Call(opts, &out, "getNumberOfValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfValidators is a free data retrieval call binding the contract method 0x9031d913.
//
// Solidity: function getNumberOfValidators() view returns(uint256)
func (_DistKeyContract *DistKeyContractSession) GetNumberOfValidators() (*big.Int, error) {
	return _DistKeyContract.Contract.GetNumberOfValidators(&_DistKeyContract.CallOpts)
}

// GetNumberOfValidators is a free data retrieval call binding the contract method 0x9031d913.
//
// Solidity: function getNumberOfValidators() view returns(uint256)
func (_DistKeyContract *DistKeyContractCallerSession) GetNumberOfValidators() (*big.Int, error) {
	return _DistKeyContract.Contract.GetNumberOfValidators(&_DistKeyContract.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(uint256[4])
func (_DistKeyContract *DistKeyContractCaller) GetPublicKey(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _DistKeyContract.contract.Call(opts, &out, "getPublicKey")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(uint256[4])
func (_DistKeyContract *DistKeyContractSession) GetPublicKey() ([4]*big.Int, error) {
	return _DistKeyContract.Contract.GetPublicKey(&_DistKeyContract.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x2e334452.
//
// Solidity: function getPublicKey() view returns(uint256[4])
func (_DistKeyContract *DistKeyContractCallerSession) GetPublicKey() ([4]*big.Int, error) {
	return _DistKeyContract.Contract.GetPublicKey(&_DistKeyContract.CallOpts)
}

// DisputePublicKey is a paid mutator transaction binding the contract method 0x3a12141d.
//
// Solidity: function disputePublicKey() returns()
func (_DistKeyContract *DistKeyContractTransactor) DisputePublicKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistKeyContract.contract.Transact(opts, "disputePublicKey")
}

// DisputePublicKey is a paid mutator transaction binding the contract method 0x3a12141d.
//
// Solidity: function disputePublicKey() returns()
func (_DistKeyContract *DistKeyContractSession) DisputePublicKey() (*types.Transaction, error) {
	return _DistKeyContract.Contract.DisputePublicKey(&_DistKeyContract.TransactOpts)
}

// DisputePublicKey is a paid mutator transaction binding the contract method 0x3a12141d.
//
// Solidity: function disputePublicKey() returns()
func (_DistKeyContract *DistKeyContractTransactorSession) DisputePublicKey() (*types.Transaction, error) {
	return _DistKeyContract.Contract.DisputePublicKey(&_DistKeyContract.TransactOpts)
}

// Generate is a paid mutator transaction binding the contract method 0x2a1bbc34.
//
// Solidity: function generate() returns()
func (_DistKeyContract *DistKeyContractTransactor) Generate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistKeyContract.contract.Transact(opts, "generate")
}

// Generate is a paid mutator transaction binding the contract method 0x2a1bbc34.
//
// Solidity: function generate() returns()
func (_DistKeyContract *DistKeyContractSession) Generate() (*types.Transaction, error) {
	return _DistKeyContract.Contract.Generate(&_DistKeyContract.TransactOpts)
}

// Generate is a paid mutator transaction binding the contract method 0x2a1bbc34.
//
// Solidity: function generate() returns()
func (_DistKeyContract *DistKeyContractTransactorSession) Generate() (*types.Transaction, error) {
	return _DistKeyContract.Contract.Generate(&_DistKeyContract.TransactOpts)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xf4cc18a3.
//
// Solidity: function setPublicKey(uint256[4] _publicKey, uint256 _numberOfValidators) returns()
func (_DistKeyContract *DistKeyContractTransactor) SetPublicKey(opts *bind.TransactOpts, _publicKey [4]*big.Int, _numberOfValidators *big.Int) (*types.Transaction, error) {
	return _DistKeyContract.contract.Transact(opts, "setPublicKey", _publicKey, _numberOfValidators)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xf4cc18a3.
//
// Solidity: function setPublicKey(uint256[4] _publicKey, uint256 _numberOfValidators) returns()
func (_DistKeyContract *DistKeyContractSession) SetPublicKey(_publicKey [4]*big.Int, _numberOfValidators *big.Int) (*types.Transaction, error) {
	return _DistKeyContract.Contract.SetPublicKey(&_DistKeyContract.TransactOpts, _publicKey, _numberOfValidators)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xf4cc18a3.
//
// Solidity: function setPublicKey(uint256[4] _publicKey, uint256 _numberOfValidators) returns()
func (_DistKeyContract *DistKeyContractTransactorSession) SetPublicKey(_publicKey [4]*big.Int, _numberOfValidators *big.Int) (*types.Transaction, error) {
	return _DistKeyContract.Contract.SetPublicKey(&_DistKeyContract.TransactOpts, _publicKey, _numberOfValidators)
}

// SetRegistryContract is a paid mutator transaction binding the contract method 0x028ebc44.
//
// Solidity: function setRegistryContract(address _registryContract) returns()
func (_DistKeyContract *DistKeyContractTransactor) SetRegistryContract(opts *bind.TransactOpts, _registryContract common.Address) (*types.Transaction, error) {
	return _DistKeyContract.contract.Transact(opts, "setRegistryContract", _registryContract)
}

// SetRegistryContract is a paid mutator transaction binding the contract method 0x028ebc44.
//
// Solidity: function setRegistryContract(address _registryContract) returns()
func (_DistKeyContract *DistKeyContractSession) SetRegistryContract(_registryContract common.Address) (*types.Transaction, error) {
	return _DistKeyContract.Contract.SetRegistryContract(&_DistKeyContract.TransactOpts, _registryContract)
}

// SetRegistryContract is a paid mutator transaction binding the contract method 0x028ebc44.
//
// Solidity: function setRegistryContract(address _registryContract) returns()
func (_DistKeyContract *DistKeyContractTransactorSession) SetRegistryContract(_registryContract common.Address) (*types.Transaction, error) {
	return _DistKeyContract.Contract.SetRegistryContract(&_DistKeyContract.TransactOpts, _registryContract)
}

// DistKeyContractDistKeyGenerationLogIterator is returned from FilterDistKeyGenerationLog and is used to iterate over the raw logs and unpacked data for DistKeyGenerationLog events raised by the DistKeyContract contract.
type DistKeyContractDistKeyGenerationLogIterator struct {
	Event *DistKeyContractDistKeyGenerationLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DistKeyContractDistKeyGenerationLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistKeyContractDistKeyGenerationLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DistKeyContractDistKeyGenerationLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DistKeyContractDistKeyGenerationLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistKeyContractDistKeyGenerationLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistKeyContractDistKeyGenerationLog represents a DistKeyGenerationLog event raised by the DistKeyContract contract.
type DistKeyContractDistKeyGenerationLog struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistKeyGenerationLog is a free log retrieval operation binding the contract event 0xd5bc20dfd67f9905e22927db6247633f2bf9ca553e4dd2f850259c9414893127.
//
// Solidity: event DistKeyGenerationLog(uint256 threshold)
func (_DistKeyContract *DistKeyContractFilterer) FilterDistKeyGenerationLog(opts *bind.FilterOpts) (*DistKeyContractDistKeyGenerationLogIterator, error) {

	logs, sub, err := _DistKeyContract.contract.FilterLogs(opts, "DistKeyGenerationLog")
	if err != nil {
		return nil, err
	}
	return &DistKeyContractDistKeyGenerationLogIterator{contract: _DistKeyContract.contract, event: "DistKeyGenerationLog", logs: logs, sub: sub}, nil
}

// WatchDistKeyGenerationLog is a free log subscription operation binding the contract event 0xd5bc20dfd67f9905e22927db6247633f2bf9ca553e4dd2f850259c9414893127.
//
// Solidity: event DistKeyGenerationLog(uint256 threshold)
func (_DistKeyContract *DistKeyContractFilterer) WatchDistKeyGenerationLog(opts *bind.WatchOpts, sink chan<- *DistKeyContractDistKeyGenerationLog) (event.Subscription, error) {

	logs, sub, err := _DistKeyContract.contract.WatchLogs(opts, "DistKeyGenerationLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistKeyContractDistKeyGenerationLog)
				if err := _DistKeyContract.contract.UnpackLog(event, "DistKeyGenerationLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistKeyGenerationLog is a log parse operation binding the contract event 0xd5bc20dfd67f9905e22927db6247633f2bf9ca553e4dd2f850259c9414893127.
//
// Solidity: event DistKeyGenerationLog(uint256 threshold)
func (_DistKeyContract *DistKeyContractFilterer) ParseDistKeyGenerationLog(log types.Log) (*DistKeyContractDistKeyGenerationLog, error) {
	event := new(DistKeyContractDistKeyGenerationLog)
	if err := _DistKeyContract.contract.UnpackLog(event, "DistKeyGenerationLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

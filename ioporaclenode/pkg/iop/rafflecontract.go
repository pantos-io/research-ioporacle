// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iop

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

// RaffleContractABI is the input ABI used to generate the binding from.
const RaffleContractABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceiveLog\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true}]"

// RaffleContract is an auto generated Go binding around an Ethereum contract.
type RaffleContract struct {
	RaffleContractCaller     // Read-only binding to the contract
	RaffleContractTransactor // Write-only binding to the contract
	RaffleContractFilterer   // Log filterer for contract events
}

// RaffleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaffleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaffleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaffleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaffleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaffleContractSession struct {
	Contract     *RaffleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaffleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaffleContractCallerSession struct {
	Contract *RaffleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RaffleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaffleContractTransactorSession struct {
	Contract     *RaffleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RaffleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaffleContractRaw struct {
	Contract *RaffleContract // Generic contract binding to access the raw methods on
}

// RaffleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaffleContractCallerRaw struct {
	Contract *RaffleContractCaller // Generic read-only contract binding to access the raw methods on
}

// RaffleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaffleContractTransactorRaw struct {
	Contract *RaffleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaffleContract creates a new instance of RaffleContract, bound to a specific deployed contract.
func NewRaffleContract(address common.Address, backend bind.ContractBackend) (*RaffleContract, error) {
	contract, err := bindRaffleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RaffleContract{RaffleContractCaller: RaffleContractCaller{contract: contract}, RaffleContractTransactor: RaffleContractTransactor{contract: contract}, RaffleContractFilterer: RaffleContractFilterer{contract: contract}}, nil
}

// NewRaffleContractCaller creates a new read-only instance of RaffleContract, bound to a specific deployed contract.
func NewRaffleContractCaller(address common.Address, caller bind.ContractCaller) (*RaffleContractCaller, error) {
	contract, err := bindRaffleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleContractCaller{contract: contract}, nil
}

// NewRaffleContractTransactor creates a new write-only instance of RaffleContract, bound to a specific deployed contract.
func NewRaffleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RaffleContractTransactor, error) {
	contract, err := bindRaffleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaffleContractTransactor{contract: contract}, nil
}

// NewRaffleContractFilterer creates a new log filterer instance of RaffleContract, bound to a specific deployed contract.
func NewRaffleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RaffleContractFilterer, error) {
	contract, err := bindRaffleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaffleContractFilterer{contract: contract}, nil
}

// bindRaffleContract binds a generic wrapper to an already deployed contract.
func bindRaffleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RaffleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleContract *RaffleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleContract.Contract.RaffleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleContract *RaffleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleContract.Contract.RaffleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleContract *RaffleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleContract.Contract.RaffleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaffleContract *RaffleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaffleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaffleContract *RaffleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaffleContract *RaffleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaffleContract.Contract.contract.Transact(opts, method, params...)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RaffleContract *RaffleContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaffleContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RaffleContract *RaffleContractSession) Receive() (*types.Transaction, error) {
	return _RaffleContract.Contract.Receive(&_RaffleContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RaffleContract *RaffleContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RaffleContract.Contract.Receive(&_RaffleContract.TransactOpts)
}

// RaffleContractReceiveLogIterator is returned from FilterReceiveLog and is used to iterate over the raw logs and unpacked data for ReceiveLog events raised by the RaffleContract contract.
type RaffleContractReceiveLogIterator struct {
	Event *RaffleContractReceiveLog // Event containing the contract specifics and raw log

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
func (it *RaffleContractReceiveLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaffleContractReceiveLog)
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
		it.Event = new(RaffleContractReceiveLog)
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
func (it *RaffleContractReceiveLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaffleContractReceiveLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaffleContractReceiveLog represents a ReceiveLog event raised by the RaffleContract contract.
type RaffleContractReceiveLog struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveLog is a free log retrieval operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 amount)
func (_RaffleContract *RaffleContractFilterer) FilterReceiveLog(opts *bind.FilterOpts, sender []common.Address) (*RaffleContractReceiveLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RaffleContract.contract.FilterLogs(opts, "ReceiveLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &RaffleContractReceiveLogIterator{contract: _RaffleContract.contract, event: "ReceiveLog", logs: logs, sub: sub}, nil
}

// WatchReceiveLog is a free log subscription operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 amount)
func (_RaffleContract *RaffleContractFilterer) WatchReceiveLog(opts *bind.WatchOpts, sink chan<- *RaffleContractReceiveLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RaffleContract.contract.WatchLogs(opts, "ReceiveLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaffleContractReceiveLog)
				if err := _RaffleContract.contract.UnpackLog(event, "ReceiveLog", log); err != nil {
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

// ParseReceiveLog is a log parse operation binding the contract event 0x3001d122178ee9da2e7224593b8383a606c707c305951086f284673d62757cf5.
//
// Solidity: event ReceiveLog(address indexed sender, uint256 amount)
func (_RaffleContract *RaffleContractFilterer) ParseReceiveLog(log types.Log) (*RaffleContractReceiveLog, error) {
	event := new(RaffleContractReceiveLog)
	if err := _RaffleContract.contract.UnpackLog(event, "ReceiveLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

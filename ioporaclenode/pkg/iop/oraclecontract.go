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

// OracleContractValidationResult is an auto generated low-level Go binding around an user-defined struct.
type OracleContractValidationResult struct {
	Id     *big.Int
	Result bool
}

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distKeyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"SubmitValidationResultLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"ValidateTransactionLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"findValidationResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.ValidationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"validationResultExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// FindValidationResult is a free data retrieval call binding the contract method 0x953bde29.
//
// Solidity: function findValidationResult(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractCaller) FindValidationResult(opts *bind.CallOpts, _id *big.Int) (OracleContractValidationResult, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findValidationResult", _id)

	if err != nil {
		return *new(OracleContractValidationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractValidationResult)).(*OracleContractValidationResult)

	return out0, err

}

// FindValidationResult is a free data retrieval call binding the contract method 0x953bde29.
//
// Solidity: function findValidationResult(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractSession) FindValidationResult(_id *big.Int) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindValidationResult(&_OracleContract.CallOpts, _id)
}

// FindValidationResult is a free data retrieval call binding the contract method 0x953bde29.
//
// Solidity: function findValidationResult(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractCallerSession) FindValidationResult(_id *big.Int) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindValidationResult(&_OracleContract.CallOpts, _id)
}

// ValidationResultExists is a free data retrieval call binding the contract method 0xef915125.
//
// Solidity: function validationResultExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractCaller) ValidationResultExists(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "validationResultExists", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidationResultExists is a free data retrieval call binding the contract method 0xef915125.
//
// Solidity: function validationResultExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractSession) ValidationResultExists(_id *big.Int) (bool, error) {
	return _OracleContract.Contract.ValidationResultExists(&_OracleContract.CallOpts, _id)
}

// ValidationResultExists is a free data retrieval call binding the contract method 0xef915125.
//
// Solidity: function validationResultExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractCallerSession) ValidationResultExists(_id *big.Int) (bool, error) {
	return _OracleContract.Contract.ValidationResultExists(&_OracleContract.CallOpts, _id)
}

// SubmitValidationResult is a paid mutator transaction binding the contract method 0x4c281081.
//
// Solidity: function submitValidationResult(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitValidationResult(opts *bind.TransactOpts, _id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitValidationResult", _id, _result, _signature)
}

// SubmitValidationResult is a paid mutator transaction binding the contract method 0x4c281081.
//
// Solidity: function submitValidationResult(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitValidationResult(_id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitValidationResult(&_OracleContract.TransactOpts, _id, _result, _signature)
}

// SubmitValidationResult is a paid mutator transaction binding the contract method 0x4c281081.
//
// Solidity: function submitValidationResult(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitValidationResult(_id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitValidationResult(&_OracleContract.TransactOpts, _id, _result, _signature)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x5f6c680d.
//
// Solidity: function validateTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _hash, _confirmations)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x5f6c680d.
//
// Solidity: function validateTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x5f6c680d.
//
// Solidity: function validateTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// OracleContractSubmitValidationResultLogIterator is returned from FilterSubmitValidationResultLog and is used to iterate over the raw logs and unpacked data for SubmitValidationResultLog events raised by the OracleContract contract.
type OracleContractSubmitValidationResultLogIterator struct {
	Event *OracleContractSubmitValidationResultLog // Event containing the contract specifics and raw log

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
func (it *OracleContractSubmitValidationResultLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractSubmitValidationResultLog)
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
		it.Event = new(OracleContractSubmitValidationResultLog)
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
func (it *OracleContractSubmitValidationResultLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractSubmitValidationResultLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractSubmitValidationResultLog represents a SubmitValidationResultLog event raised by the OracleContract contract.
type OracleContractSubmitValidationResultLog struct {
	Sender common.Address
	Id     *big.Int
	Result bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitValidationResultLog is a free log retrieval operation binding the contract event 0x36cf801fa1e8739e3bf8426bba2c30eca735b965ef2ab047727480e5bf49847a.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) FilterSubmitValidationResultLog(opts *bind.FilterOpts, sender []common.Address) (*OracleContractSubmitValidationResultLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "SubmitValidationResultLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractSubmitValidationResultLogIterator{contract: _OracleContract.contract, event: "SubmitValidationResultLog", logs: logs, sub: sub}, nil
}

// WatchSubmitValidationResultLog is a free log subscription operation binding the contract event 0x36cf801fa1e8739e3bf8426bba2c30eca735b965ef2ab047727480e5bf49847a.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) WatchSubmitValidationResultLog(opts *bind.WatchOpts, sink chan<- *OracleContractSubmitValidationResultLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "SubmitValidationResultLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractSubmitValidationResultLog)
				if err := _OracleContract.contract.UnpackLog(event, "SubmitValidationResultLog", log); err != nil {
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

// ParseSubmitValidationResultLog is a log parse operation binding the contract event 0x36cf801fa1e8739e3bf8426bba2c30eca735b965ef2ab047727480e5bf49847a.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) ParseSubmitValidationResultLog(log types.Log) (*OracleContractSubmitValidationResultLog, error) {
	event := new(OracleContractSubmitValidationResultLog)
	if err := _OracleContract.contract.UnpackLog(event, "SubmitValidationResultLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidateTransactionLogIterator is returned from FilterValidateTransactionLog and is used to iterate over the raw logs and unpacked data for ValidateTransactionLog events raised by the OracleContract contract.
type OracleContractValidateTransactionLogIterator struct {
	Event *OracleContractValidateTransactionLog // Event containing the contract specifics and raw log

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
func (it *OracleContractValidateTransactionLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidateTransactionLog)
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
		it.Event = new(OracleContractValidateTransactionLog)
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
func (it *OracleContractValidateTransactionLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidateTransactionLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidateTransactionLog represents a ValidateTransactionLog event raised by the OracleContract contract.
type OracleContractValidateTransactionLog struct {
	Sender        common.Address
	Id            *big.Int
	Hash          [32]byte
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidateTransactionLog is a free log retrieval operation binding the contract event 0xdb3924693edfb1380a509834c8de41c8d57794e3b056a3cc6b0bac025ba57da4.
//
// Solidity: event ValidateTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) FilterValidateTransactionLog(opts *bind.FilterOpts, sender []common.Address, id []*big.Int) (*OracleContractValidateTransactionLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidateTransactionLog", senderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidateTransactionLogIterator{contract: _OracleContract.contract, event: "ValidateTransactionLog", logs: logs, sub: sub}, nil
}

// WatchValidateTransactionLog is a free log subscription operation binding the contract event 0xdb3924693edfb1380a509834c8de41c8d57794e3b056a3cc6b0bac025ba57da4.
//
// Solidity: event ValidateTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) WatchValidateTransactionLog(opts *bind.WatchOpts, sink chan<- *OracleContractValidateTransactionLog, sender []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidateTransactionLog", senderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidateTransactionLog)
				if err := _OracleContract.contract.UnpackLog(event, "ValidateTransactionLog", log); err != nil {
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

// ParseValidateTransactionLog is a log parse operation binding the contract event 0xdb3924693edfb1380a509834c8de41c8d57794e3b056a3cc6b0bac025ba57da4.
//
// Solidity: event ValidateTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) ParseValidateTransactionLog(log types.Log) (*OracleContractValidateTransactionLog, error) {
	event := new(OracleContractValidateTransactionLog)
	if err := _OracleContract.contract.UnpackLog(event, "ValidateTransactionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

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

// ECDSAOracleContractVerificationResult is an auto generated low-level Go binding around an user-defined struct.
type ECDSAOracleContractVerificationResult struct {
	Id      *big.Int
	Request *big.Int
	Result  bool
}

// ECDSAOracleContractABI is the input ABI used to generate the binding from.
const ECDSAOracleContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"SubmitVerificationLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"findVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"internalType\":\"structECDSAOracleContract.VerificationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"verificationExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"encodeResult\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ECDSAOracleContract is an auto generated Go binding around an Ethereum contract.
type ECDSAOracleContract struct {
	ECDSAOracleContractCaller     // Read-only binding to the contract
	ECDSAOracleContractTransactor // Write-only binding to the contract
	ECDSAOracleContractFilterer   // Log filterer for contract events
}

// ECDSAOracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSAOracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAOracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSAOracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAOracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAOracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAOracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSAOracleContractSession struct {
	Contract     *ECDSAOracleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ECDSAOracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSAOracleContractCallerSession struct {
	Contract *ECDSAOracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ECDSAOracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSAOracleContractTransactorSession struct {
	Contract     *ECDSAOracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ECDSAOracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSAOracleContractRaw struct {
	Contract *ECDSAOracleContract // Generic contract binding to access the raw methods on
}

// ECDSAOracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSAOracleContractCallerRaw struct {
	Contract *ECDSAOracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSAOracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSAOracleContractTransactorRaw struct {
	Contract *ECDSAOracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSAOracleContract creates a new instance of ECDSAOracleContract, bound to a specific deployed contract.
func NewECDSAOracleContract(address common.Address, backend bind.ContractBackend) (*ECDSAOracleContract, error) {
	contract, err := bindECDSAOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContract{ECDSAOracleContractCaller: ECDSAOracleContractCaller{contract: contract}, ECDSAOracleContractTransactor: ECDSAOracleContractTransactor{contract: contract}, ECDSAOracleContractFilterer: ECDSAOracleContractFilterer{contract: contract}}, nil
}

// NewECDSAOracleContractCaller creates a new read-only instance of ECDSAOracleContract, bound to a specific deployed contract.
func NewECDSAOracleContractCaller(address common.Address, caller bind.ContractCaller) (*ECDSAOracleContractCaller, error) {
	contract, err := bindECDSAOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContractCaller{contract: contract}, nil
}

// NewECDSAOracleContractTransactor creates a new write-only instance of ECDSAOracleContract, bound to a specific deployed contract.
func NewECDSAOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSAOracleContractTransactor, error) {
	contract, err := bindECDSAOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContractTransactor{contract: contract}, nil
}

// NewECDSAOracleContractFilterer creates a new log filterer instance of ECDSAOracleContract, bound to a specific deployed contract.
func NewECDSAOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAOracleContractFilterer, error) {
	contract, err := bindECDSAOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContractFilterer{contract: contract}, nil
}

// bindECDSAOracleContract binds a generic wrapper to an already deployed contract.
func bindECDSAOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAOracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAOracleContract *ECDSAOracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAOracleContract.Contract.ECDSAOracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAOracleContract *ECDSAOracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.ECDSAOracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAOracleContract *ECDSAOracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.ECDSAOracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAOracleContract *ECDSAOracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAOracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAOracleContract *ECDSAOracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAOracleContract *ECDSAOracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.contract.Transact(opts, method, params...)
}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
func (_ECDSAOracleContract *ECDSAOracleContractCaller) EncodeResult(opts *bind.CallOpts, _id *big.Int, _result bool) ([]byte, error) {
	var out []interface{}
	err := _ECDSAOracleContract.contract.Call(opts, &out, "encodeResult", _id, _result)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
func (_ECDSAOracleContract *ECDSAOracleContractSession) EncodeResult(_id *big.Int, _result bool) ([]byte, error) {
	return _ECDSAOracleContract.Contract.EncodeResult(&_ECDSAOracleContract.CallOpts, _id, _result)
}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
func (_ECDSAOracleContract *ECDSAOracleContractCallerSession) EncodeResult(_id *big.Int, _result bool) ([]byte, error) {
	return _ECDSAOracleContract.Contract.EncodeResult(&_ECDSAOracleContract.CallOpts, _id, _result)
}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool))
func (_ECDSAOracleContract *ECDSAOracleContractCaller) FindVerification(opts *bind.CallOpts, _id *big.Int) (ECDSAOracleContractVerificationResult, error) {
	var out []interface{}
	err := _ECDSAOracleContract.contract.Call(opts, &out, "findVerification", _id)

	if err != nil {
		return *new(ECDSAOracleContractVerificationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(ECDSAOracleContractVerificationResult)).(*ECDSAOracleContractVerificationResult)

	return out0, err

}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool))
func (_ECDSAOracleContract *ECDSAOracleContractSession) FindVerification(_id *big.Int) (ECDSAOracleContractVerificationResult, error) {
	return _ECDSAOracleContract.Contract.FindVerification(&_ECDSAOracleContract.CallOpts, _id)
}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool))
func (_ECDSAOracleContract *ECDSAOracleContractCallerSession) FindVerification(_id *big.Int) (ECDSAOracleContractVerificationResult, error) {
	return _ECDSAOracleContract.Contract.FindVerification(&_ECDSAOracleContract.CallOpts, _id)
}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_ECDSAOracleContract *ECDSAOracleContractCaller) VerificationExists(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _ECDSAOracleContract.contract.Call(opts, &out, "verificationExists", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_ECDSAOracleContract *ECDSAOracleContractSession) VerificationExists(_id *big.Int) (bool, error) {
	return _ECDSAOracleContract.Contract.VerificationExists(&_ECDSAOracleContract.CallOpts, _id)
}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_ECDSAOracleContract *ECDSAOracleContractCallerSession) VerificationExists(_id *big.Int) (bool, error) {
	return _ECDSAOracleContract.Contract.VerificationExists(&_ECDSAOracleContract.CallOpts, _id)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xa68c964b.
//
// Solidity: function submitVerification(uint256 _id, bool _result, bytes[] _signatures) returns()
func (_ECDSAOracleContract *ECDSAOracleContractTransactor) SubmitVerification(opts *bind.TransactOpts, _id *big.Int, _result bool, _signatures [][]byte) (*types.Transaction, error) {
	return _ECDSAOracleContract.contract.Transact(opts, "submitVerification", _id, _result, _signatures)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xa68c964b.
//
// Solidity: function submitVerification(uint256 _id, bool _result, bytes[] _signatures) returns()
func (_ECDSAOracleContract *ECDSAOracleContractSession) SubmitVerification(_id *big.Int, _result bool, _signatures [][]byte) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.SubmitVerification(&_ECDSAOracleContract.TransactOpts, _id, _result, _signatures)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xa68c964b.
//
// Solidity: function submitVerification(uint256 _id, bool _result, bytes[] _signatures) returns()
func (_ECDSAOracleContract *ECDSAOracleContractTransactorSession) SubmitVerification(_id *big.Int, _result bool, _signatures [][]byte) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.SubmitVerification(&_ECDSAOracleContract.TransactOpts, _id, _result, _signatures)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
func (_ECDSAOracleContract *ECDSAOracleContractTransactor) VerifyTransaction(opts *bind.TransactOpts, _hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _ECDSAOracleContract.contract.Transact(opts, "verifyTransaction", _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
func (_ECDSAOracleContract *ECDSAOracleContractSession) VerifyTransaction(_hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.VerifyTransaction(&_ECDSAOracleContract.TransactOpts, _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
func (_ECDSAOracleContract *ECDSAOracleContractTransactorSession) VerifyTransaction(_hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _ECDSAOracleContract.Contract.VerifyTransaction(&_ECDSAOracleContract.TransactOpts, _hash, _confirmations)
}

// ECDSAOracleContractSubmitVerificationLogIterator is returned from FilterSubmitVerificationLog and is used to iterate over the raw logs and unpacked data for SubmitVerificationLog events raised by the ECDSAOracleContract contract.
type ECDSAOracleContractSubmitVerificationLogIterator struct {
	Event *ECDSAOracleContractSubmitVerificationLog // Event containing the contract specifics and raw log

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
func (it *ECDSAOracleContractSubmitVerificationLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAOracleContractSubmitVerificationLog)
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
		it.Event = new(ECDSAOracleContractSubmitVerificationLog)
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
func (it *ECDSAOracleContractSubmitVerificationLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAOracleContractSubmitVerificationLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAOracleContractSubmitVerificationLog represents a SubmitVerificationLog event raised by the ECDSAOracleContract contract.
type ECDSAOracleContractSubmitVerificationLog struct {
	Sender  common.Address
	Id      *big.Int
	Request *big.Int
	Result  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitVerificationLog is a free log retrieval operation binding the contract event 0x1e23af75f17bd0353808dc4860a770c7c8a63fd767c2ee101c032d39dc3768fc.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) FilterSubmitVerificationLog(opts *bind.FilterOpts, sender []common.Address) (*ECDSAOracleContractSubmitVerificationLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ECDSAOracleContract.contract.FilterLogs(opts, "SubmitVerificationLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContractSubmitVerificationLogIterator{contract: _ECDSAOracleContract.contract, event: "SubmitVerificationLog", logs: logs, sub: sub}, nil
}

// WatchSubmitVerificationLog is a free log subscription operation binding the contract event 0x1e23af75f17bd0353808dc4860a770c7c8a63fd767c2ee101c032d39dc3768fc.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) WatchSubmitVerificationLog(opts *bind.WatchOpts, sink chan<- *ECDSAOracleContractSubmitVerificationLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ECDSAOracleContract.contract.WatchLogs(opts, "SubmitVerificationLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAOracleContractSubmitVerificationLog)
				if err := _ECDSAOracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
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

// ParseSubmitVerificationLog is a log parse operation binding the contract event 0x1e23af75f17bd0353808dc4860a770c7c8a63fd767c2ee101c032d39dc3768fc.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) ParseSubmitVerificationLog(log types.Log) (*ECDSAOracleContractSubmitVerificationLog, error) {
	event := new(ECDSAOracleContractSubmitVerificationLog)
	if err := _ECDSAOracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAOracleContractVerifyTransactionLogIterator is returned from FilterVerifyTransactionLog and is used to iterate over the raw logs and unpacked data for VerifyTransactionLog events raised by the ECDSAOracleContract contract.
type ECDSAOracleContractVerifyTransactionLogIterator struct {
	Event *ECDSAOracleContractVerifyTransactionLog // Event containing the contract specifics and raw log

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
func (it *ECDSAOracleContractVerifyTransactionLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAOracleContractVerifyTransactionLog)
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
		it.Event = new(ECDSAOracleContractVerifyTransactionLog)
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
func (it *ECDSAOracleContractVerifyTransactionLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAOracleContractVerifyTransactionLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAOracleContractVerifyTransactionLog represents a VerifyTransactionLog event raised by the ECDSAOracleContract contract.
type ECDSAOracleContractVerifyTransactionLog struct {
	Id            *big.Int
	Hash          string
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVerifyTransactionLog is a free log retrieval operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) FilterVerifyTransactionLog(opts *bind.FilterOpts) (*ECDSAOracleContractVerifyTransactionLogIterator, error) {

	logs, sub, err := _ECDSAOracleContract.contract.FilterLogs(opts, "VerifyTransactionLog")
	if err != nil {
		return nil, err
	}
	return &ECDSAOracleContractVerifyTransactionLogIterator{contract: _ECDSAOracleContract.contract, event: "VerifyTransactionLog", logs: logs, sub: sub}, nil
}

// WatchVerifyTransactionLog is a free log subscription operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) WatchVerifyTransactionLog(opts *bind.WatchOpts, sink chan<- *ECDSAOracleContractVerifyTransactionLog) (event.Subscription, error) {

	logs, sub, err := _ECDSAOracleContract.contract.WatchLogs(opts, "VerifyTransactionLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAOracleContractVerifyTransactionLog)
				if err := _ECDSAOracleContract.contract.UnpackLog(event, "VerifyTransactionLog", log); err != nil {
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

// ParseVerifyTransactionLog is a log parse operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_ECDSAOracleContract *ECDSAOracleContractFilterer) ParseVerifyTransactionLog(log types.Log) (*ECDSAOracleContractVerifyTransactionLog, error) {
	event := new(ECDSAOracleContractVerifyTransactionLog)
	if err := _ECDSAOracleContract.contract.UnpackLog(event, "VerifyTransactionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

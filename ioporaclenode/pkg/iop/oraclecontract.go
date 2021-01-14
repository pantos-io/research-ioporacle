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

// OracleContractVerificationResult is an auto generated low-level Go binding around an user-defined struct.
type OracleContractVerificationResult struct {
	Id     *big.Int
	Result bool
}

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"SubmitVerificationLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PUB_KEY_X_IM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"PUB_KEY_X_RE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"PUB_KEY_Y_IM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"PUB_KEY_Y_RE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"findVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.VerificationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"verificationExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

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

// PUBKEYXIM is a free data retrieval call binding the contract method 0xc5c8df21.
//
// Solidity: function PUB_KEY_X_IM() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYXIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUB_KEY_X_IM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYXIM is a free data retrieval call binding the contract method 0xc5c8df21.
//
// Solidity: function PUB_KEY_X_IM() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYXIM() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYXIM(&_OracleContract.CallOpts)
}

// PUBKEYXIM is a free data retrieval call binding the contract method 0xc5c8df21.
//
// Solidity: function PUB_KEY_X_IM() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYXIM() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYXIM(&_OracleContract.CallOpts)
}

// PUBKEYXRE is a free data retrieval call binding the contract method 0x4401e888.
//
// Solidity: function PUB_KEY_X_RE() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYXRE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUB_KEY_X_RE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYXRE is a free data retrieval call binding the contract method 0x4401e888.
//
// Solidity: function PUB_KEY_X_RE() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYXRE() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYXRE(&_OracleContract.CallOpts)
}

// PUBKEYXRE is a free data retrieval call binding the contract method 0x4401e888.
//
// Solidity: function PUB_KEY_X_RE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYXRE() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYXRE(&_OracleContract.CallOpts)
}

// PUBKEYYIM is a free data retrieval call binding the contract method 0x52c41d22.
//
// Solidity: function PUB_KEY_Y_IM() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYYIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUB_KEY_Y_IM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYYIM is a free data retrieval call binding the contract method 0x52c41d22.
//
// Solidity: function PUB_KEY_Y_IM() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYYIM() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYYIM(&_OracleContract.CallOpts)
}

// PUBKEYYIM is a free data retrieval call binding the contract method 0x52c41d22.
//
// Solidity: function PUB_KEY_Y_IM() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYYIM() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYYIM(&_OracleContract.CallOpts)
}

// PUBKEYYRE is a free data retrieval call binding the contract method 0x24815ca5.
//
// Solidity: function PUB_KEY_Y_RE() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYYRE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUB_KEY_Y_RE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYYRE is a free data retrieval call binding the contract method 0x24815ca5.
//
// Solidity: function PUB_KEY_Y_RE() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYYRE() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYYRE(&_OracleContract.CallOpts)
}

// PUBKEYYRE is a free data retrieval call binding the contract method 0x24815ca5.
//
// Solidity: function PUB_KEY_Y_RE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYYRE() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYYRE(&_OracleContract.CallOpts)
}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractCaller) FindVerification(opts *bind.CallOpts, _id *big.Int) (OracleContractVerificationResult, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findVerification", _id)

	if err != nil {
		return *new(OracleContractVerificationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractVerificationResult)).(*OracleContractVerificationResult)

	return out0, err

}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractSession) FindVerification(_id *big.Int) (OracleContractVerificationResult, error) {
	return _OracleContract.Contract.FindVerification(&_OracleContract.CallOpts, _id)
}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,bool))
func (_OracleContract *OracleContractCallerSession) FindVerification(_id *big.Int) (OracleContractVerificationResult, error) {
	return _OracleContract.Contract.FindVerification(&_OracleContract.CallOpts, _id)
}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractCaller) VerificationExists(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "verificationExists", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractSession) VerificationExists(_id *big.Int) (bool, error) {
	return _OracleContract.Contract.VerificationExists(&_OracleContract.CallOpts, _id)
}

// VerificationExists is a free data retrieval call binding the contract method 0xc6704288.
//
// Solidity: function verificationExists(uint256 _id) view returns(bool)
func (_OracleContract *OracleContractCallerSession) VerificationExists(_id *big.Int) (bool, error) {
	return _OracleContract.Contract.VerificationExists(&_OracleContract.CallOpts, _id)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xb1f98c1f.
//
// Solidity: function submitVerification(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitVerification(opts *bind.TransactOpts, _id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitVerification", _id, _result, _signature)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xb1f98c1f.
//
// Solidity: function submitVerification(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitVerification(_id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitVerification(&_OracleContract.TransactOpts, _id, _result, _signature)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0xb1f98c1f.
//
// Solidity: function submitVerification(uint256 _id, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitVerification(_id *big.Int, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitVerification(&_OracleContract.TransactOpts, _id, _result, _signature)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xd0508fd6.
//
// Solidity: function verifyTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractTransactor) VerifyTransaction(opts *bind.TransactOpts, _hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyTransaction", _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xd0508fd6.
//
// Solidity: function verifyTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractSession) VerifyTransaction(_hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xd0508fd6.
//
// Solidity: function verifyTransaction(bytes32 _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractTransactorSession) VerifyTransaction(_hash [32]byte, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// OracleContractSubmitVerificationLogIterator is returned from FilterSubmitVerificationLog and is used to iterate over the raw logs and unpacked data for SubmitVerificationLog events raised by the OracleContract contract.
type OracleContractSubmitVerificationLogIterator struct {
	Event *OracleContractSubmitVerificationLog // Event containing the contract specifics and raw log

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
func (it *OracleContractSubmitVerificationLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractSubmitVerificationLog)
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
		it.Event = new(OracleContractSubmitVerificationLog)
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
func (it *OracleContractSubmitVerificationLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractSubmitVerificationLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractSubmitVerificationLog represents a SubmitVerificationLog event raised by the OracleContract contract.
type OracleContractSubmitVerificationLog struct {
	Sender common.Address
	Id     *big.Int
	Result bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitVerificationLog is a free log retrieval operation binding the contract event 0x8bac4ca8b4a878693f815b8e8e244f09b02c960dcd4937470482ca16ab12692b.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) FilterSubmitVerificationLog(opts *bind.FilterOpts, sender []common.Address) (*OracleContractSubmitVerificationLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "SubmitVerificationLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractSubmitVerificationLogIterator{contract: _OracleContract.contract, event: "SubmitVerificationLog", logs: logs, sub: sub}, nil
}

// WatchSubmitVerificationLog is a free log subscription operation binding the contract event 0x8bac4ca8b4a878693f815b8e8e244f09b02c960dcd4937470482ca16ab12692b.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) WatchSubmitVerificationLog(opts *bind.WatchOpts, sink chan<- *OracleContractSubmitVerificationLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "SubmitVerificationLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractSubmitVerificationLog)
				if err := _OracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
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

// ParseSubmitVerificationLog is a log parse operation binding the contract event 0x8bac4ca8b4a878693f815b8e8e244f09b02c960dcd4937470482ca16ab12692b.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, bool result)
func (_OracleContract *OracleContractFilterer) ParseSubmitVerificationLog(log types.Log) (*OracleContractSubmitVerificationLog, error) {
	event := new(OracleContractSubmitVerificationLog)
	if err := _OracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractVerifyTransactionLogIterator is returned from FilterVerifyTransactionLog and is used to iterate over the raw logs and unpacked data for VerifyTransactionLog events raised by the OracleContract contract.
type OracleContractVerifyTransactionLogIterator struct {
	Event *OracleContractVerifyTransactionLog // Event containing the contract specifics and raw log

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
func (it *OracleContractVerifyTransactionLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractVerifyTransactionLog)
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
		it.Event = new(OracleContractVerifyTransactionLog)
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
func (it *OracleContractVerifyTransactionLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractVerifyTransactionLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractVerifyTransactionLog represents a VerifyTransactionLog event raised by the OracleContract contract.
type OracleContractVerifyTransactionLog struct {
	Sender        common.Address
	Id            *big.Int
	Hash          [32]byte
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVerifyTransactionLog is a free log retrieval operation binding the contract event 0x7273af0961d6b722075a949b6dddbe85d4688ee4c18ab11a8d68b34a137d53d6.
//
// Solidity: event VerifyTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) FilterVerifyTransactionLog(opts *bind.FilterOpts, sender []common.Address, id []*big.Int) (*OracleContractVerifyTransactionLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "VerifyTransactionLog", senderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractVerifyTransactionLogIterator{contract: _OracleContract.contract, event: "VerifyTransactionLog", logs: logs, sub: sub}, nil
}

// WatchVerifyTransactionLog is a free log subscription operation binding the contract event 0x7273af0961d6b722075a949b6dddbe85d4688ee4c18ab11a8d68b34a137d53d6.
//
// Solidity: event VerifyTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) WatchVerifyTransactionLog(opts *bind.WatchOpts, sink chan<- *OracleContractVerifyTransactionLog, sender []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "VerifyTransactionLog", senderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractVerifyTransactionLog)
				if err := _OracleContract.contract.UnpackLog(event, "VerifyTransactionLog", log); err != nil {
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

// ParseVerifyTransactionLog is a log parse operation binding the contract event 0x7273af0961d6b722075a949b6dddbe85d4688ee4c18ab11a8d68b34a137d53d6.
//
// Solidity: event VerifyTransactionLog(address indexed sender, uint256 indexed id, bytes32 hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) ParseVerifyTransactionLog(log types.Log) (*OracleContractVerifyTransactionLog, error) {
	event := new(OracleContractVerifyTransactionLog)
	if err := _OracleContract.contract.UnpackLog(event, "VerifyTransactionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

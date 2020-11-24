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

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"countIopNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findIopNodeByAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findIopNodeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"iopNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerIopNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"verifyTransactionResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleContractFuncSigs maps the 4-byte function signature to its string representation.
var OracleContractFuncSigs = map[string]string{
	"f1551da7": "countIopNodes()",
	"76ac6e0b": "findIopNodeByAddress(address)",
	"7206c984": "findIopNodeByIndex(uint256)",
	"338780e0": "iopNodeIsRegistered(address)",
	"6f1daed3": "registerIopNode(string)",
	"ff84d3e9": "verifyTransaction(string,uint256)",
	"fd9f77ec": "verifyTransactionResult(bool)",
}

// OracleContractBin is the compiled bytecode used for deploying new contracts.
var OracleContractBin = "0x608060405234801561001057600080fd5b506107d3806100206000396000f3fe6080604052600436106100705760003560e01c806376ac6e0b1161004e57806376ac6e0b146101e9578063f1551da71461021c578063fd9f77ec14610243578063ff84d3e91461026f57610070565b8063338780e0146100755780636f1daed3146100bc5780637206c9841461012e575b600080fd5b34801561008157600080fd5b506100a86004803603602081101561009857600080fd5b50356001600160a01b03166102ec565b604080519115158252519081900360200190f35b61012c600480360360208110156100d257600080fd5b8101906020810181356401000000008111156100ed57600080fd5b8201836020820111156100ff57600080fd5b8035906020019184600183028401116401000000008311171561012157600080fd5b509092509050610347565b005b34801561013a57600080fd5b506101586004803603602081101561015157600080fd5b5035610422565b60405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101ad578181015183820152602001610195565b50505050905090810190601f1680156101da5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156101f557600080fd5b506101586004803603602081101561020c57600080fd5b50356001600160a01b031661056d565b34801561022857600080fd5b5061023161064c565b60408051918252519081900360200190f35b34801561024f57600080fd5b5061012c6004803603602081101561026657600080fd5b50351515610653565b34801561027b57600080fd5b5061012c6004803603604081101561029257600080fd5b8101906020810181356401000000008111156102ad57600080fd5b8201836020820111156102bf57600080fd5b803590602001918460018302840111640100000000831117156102e157600080fd5b919350915035610666565b6001546000906102fe57506000610342565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061032957fe5b6000918252602090912001546001600160a01b03161490505b919050565b610350336102ec565b15610397576040805162461bcd60e51b8152602060048201526012602482015271185b1c9958591e481c9959da5cdd195c995960721b604482015290519081900360640190fd5b33600081815260208190526040902080546001600160a01b03191690911781556103c56001820184846106dc565b5080546001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810180546001600160a01b0319166001600160a01b03909316929092179091556002909101555050565b600060606001548310610468576040805162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b604482015290519081900360640190fd5b61047061075a565b6000806001868154811061048057fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156105485780601f1061051d57610100808354040283529160200191610548565b820191906000526020600020905b81548152906001019060200180831161052b57829003601f168201915b5050509183525050600291909101546020918201528151910151909350915050915091565b6000606061057a836102ec565b6105b7576040805162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b604482015290519081900360640190fd5b6105bf61075a565b6001600160a01b0384811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156105485780601f1061051d57610100808354040283529160200191610548565b6001545b90565b6002805460ff1916911515919091179055565b7fd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28600184848460405180858152602001806020018381526020018281038252858582818152602001925080828437600083820152604051601f909101601f191690920182900397509095505050505050a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061071d5782800160ff1982351617855561074a565b8280016001018555821561074a579182015b8281111561074a57823582559160200191906001019061072f565b50610756929150610784565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b61065091905b80821115610756576000815560010161078a56fea265627a7a72315820356f1f30eb1604f7b6882d2ecbe61f938e59523ab6092d385cc9d71ef6d69a6c64736f6c63430005100032"

// DeployOracleContract deploys a new Ethereum contract, binding an instance of OracleContract to it.
func DeployOracleContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OracleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

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

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_OracleContract *OracleContractCaller) CountIopNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "countIopNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_OracleContract *OracleContractSession) CountIopNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountIopNodes(&_OracleContract.CallOpts)
}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CountIopNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountIopNodes(&_OracleContract.CallOpts)
}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns(address, string)
func (_OracleContract *OracleContractCaller) FindIopNodeByAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, string, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findIopNodeByAddress", _addr)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns(address, string)
func (_OracleContract *OracleContractSession) FindIopNodeByAddress(_addr common.Address) (common.Address, string, error) {
	return _OracleContract.Contract.FindIopNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns(address, string)
func (_OracleContract *OracleContractCallerSession) FindIopNodeByAddress(_addr common.Address) (common.Address, string, error) {
	return _OracleContract.Contract.FindIopNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns(address, string)
func (_OracleContract *OracleContractCaller) FindIopNodeByIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, string, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findIopNodeByIndex", _index)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns(address, string)
func (_OracleContract *OracleContractSession) FindIopNodeByIndex(_index *big.Int) (common.Address, string, error) {
	return _OracleContract.Contract.FindIopNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns(address, string)
func (_OracleContract *OracleContractCallerSession) FindIopNodeByIndex(_index *big.Int) (common.Address, string, error) {
	return _OracleContract.Contract.FindIopNodeByIndex(&_OracleContract.CallOpts, _index)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) IopNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "iopNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.IopNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.IopNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractTransactor) RegisterIopNode(opts *bind.TransactOpts, _ipAddr string) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "registerIopNode", _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterIopNode(&_OracleContract.TransactOpts, _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractTransactorSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterIopNode(&_OracleContract.TransactOpts, _ipAddr)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) returns()
func (_OracleContract *OracleContractTransactor) VerifyTransaction(opts *bind.TransactOpts, _hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyTransaction", _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) returns()
func (_OracleContract *OracleContractSession) VerifyTransaction(_hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) returns()
func (_OracleContract *OracleContractTransactorSession) VerifyTransaction(_hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// VerifyTransactionResult is a paid mutator transaction binding the contract method 0xfd9f77ec.
//
// Solidity: function verifyTransactionResult(bool _result) returns()
func (_OracleContract *OracleContractTransactor) VerifyTransactionResult(opts *bind.TransactOpts, _result bool) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyTransactionResult", _result)
}

// VerifyTransactionResult is a paid mutator transaction binding the contract method 0xfd9f77ec.
//
// Solidity: function verifyTransactionResult(bool _result) returns()
func (_OracleContract *OracleContractSession) VerifyTransactionResult(_result bool) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransactionResult(&_OracleContract.TransactOpts, _result)
}

// VerifyTransactionResult is a paid mutator transaction binding the contract method 0xfd9f77ec.
//
// Solidity: function verifyTransactionResult(bool _result) returns()
func (_OracleContract *OracleContractTransactorSession) VerifyTransactionResult(_result bool) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransactionResult(&_OracleContract.TransactOpts, _result)
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
	Id            *big.Int
	Hash          string
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVerifyTransactionLog is a free log retrieval operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) FilterVerifyTransactionLog(opts *bind.FilterOpts) (*OracleContractVerifyTransactionLogIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "VerifyTransactionLog")
	if err != nil {
		return nil, err
	}
	return &OracleContractVerifyTransactionLogIterator{contract: _OracleContract.contract, event: "VerifyTransactionLog", logs: logs, sub: sub}, nil
}

// WatchVerifyTransactionLog is a free log subscription operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) WatchVerifyTransactionLog(opts *bind.WatchOpts, sink chan<- *OracleContractVerifyTransactionLog) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "VerifyTransactionLog")
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

// ParseVerifyTransactionLog is a log parse operation binding the contract event 0xd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28.
//
// Solidity: event VerifyTransactionLog(uint256 id, string hash, uint256 confirmations)
func (_OracleContract *OracleContractFilterer) ParseVerifyTransactionLog(log types.Log) (*OracleContractVerifyTransactionLog, error) {
	event := new(OracleContractVerifyTransactionLog)
	if err := _OracleContract.contract.UnpackLog(event, "VerifyTransactionLog", log); err != nil {
		return nil, err
	}
	return event, nil
}

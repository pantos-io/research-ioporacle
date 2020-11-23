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

// IopOracleContractABI is the input ABI used to generate the binding from.
const IopOracleContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countIopNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findIopNodeByAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findIopNodeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"iopNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerIopNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IopOracleContractFuncSigs maps the 4-byte function signature to its string representation.
var IopOracleContractFuncSigs = map[string]string{
	"f1551da7": "countIopNodes()",
	"76ac6e0b": "findIopNodeByAddress(address)",
	"7206c984": "findIopNodeByIndex(uint256)",
	"338780e0": "iopNodeIsRegistered(address)",
	"6f1daed3": "registerIopNode(string)",
}

// IopOracleContractBin is the compiled bytecode used for deploying new contracts.
var IopOracleContractBin = "0x608060405234801561001057600080fd5b5061067b806100206000396000f3fe60806040526004361061004a5760003560e01c8063338780e01461004f5780636f1daed3146100965780637206c9841461010857806376ac6e0b146101c3578063f1551da7146101f6575b600080fd5b34801561005b57600080fd5b506100826004803603602081101561007257600080fd5b50356001600160a01b031661021d565b604080519115158252519081900360200190f35b610106600480360360208110156100ac57600080fd5b8101906020810181356401000000008111156100c757600080fd5b8201836020820111156100d957600080fd5b803590602001918460018302840111640100000000831117156100fb57600080fd5b509092509050610278565b005b34801561011457600080fd5b506101326004803603602081101561012b57600080fd5b5035610353565b60405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561018757818101518382015260200161016f565b50505050905090810190601f1680156101b45780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156101cf57600080fd5b50610132600480360360208110156101e657600080fd5b50356001600160a01b031661049e565b34801561020257600080fd5b5061020b61057d565b60408051918252519081900360200190f35b60015460009061022f57506000610273565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061025a57fe5b6000918252602090912001546001600160a01b03161490505b919050565b6102813361021d565b156102c8576040805162461bcd60e51b8152602060048201526012602482015271185b1c9958591e481c9959da5cdd195c995960721b604482015290519081900360640190fd5b33600081815260208190526040902080546001600160a01b03191690911781556102f6600182018484610584565b5080546001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810180546001600160a01b0319166001600160a01b03909316929092179091556002909101555050565b600060606001548310610399576040805162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b604482015290519081900360640190fd5b6103a1610602565b600080600186815481106103b157fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156104795780601f1061044e57610100808354040283529160200191610479565b820191906000526020600020905b81548152906001019060200180831161045c57829003601f168201915b5050509183525050600291909101546020918201528151910151909350915050915091565b600060606104ab8361021d565b6104e8576040805162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b604482015290519081900360640190fd5b6104f0610602565b6001600160a01b0384811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156104795780601f1061044e57610100808354040283529160200191610479565b6001545b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105c55782800160ff198235161785556105f2565b828001600101855582156105f2579182015b828111156105f25782358255916020019190600101906105d7565b506105fe92915061062c565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b61058191905b808211156105fe576000815560010161063256fea265627a7a7231582039efa683d6ba3a337256375f871e8a4f870c04d217a77666d06984830881aca764736f6c63430005100032"

// DeployIopOracleContract deploys a new Ethereum contract, binding an instance of IopOracleContract to it.
func DeployIopOracleContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IopOracleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IopOracleContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IopOracleContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IopOracleContract{IopOracleContractCaller: IopOracleContractCaller{contract: contract}, IopOracleContractTransactor: IopOracleContractTransactor{contract: contract}, IopOracleContractFilterer: IopOracleContractFilterer{contract: contract}}, nil
}

// IopOracleContract is an auto generated Go binding around an Ethereum contract.
type IopOracleContract struct {
	IopOracleContractCaller     // Read-only binding to the contract
	IopOracleContractTransactor // Write-only binding to the contract
	IopOracleContractFilterer   // Log filterer for contract events
}

// IopOracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type IopOracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IopOracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IopOracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IopOracleContractSession struct {
	Contract     *IopOracleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IopOracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IopOracleContractCallerSession struct {
	Contract *IopOracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IopOracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IopOracleContractTransactorSession struct {
	Contract     *IopOracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IopOracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type IopOracleContractRaw struct {
	Contract *IopOracleContract // Generic contract binding to access the raw methods on
}

// IopOracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IopOracleContractCallerRaw struct {
	Contract *IopOracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// IopOracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IopOracleContractTransactorRaw struct {
	Contract *IopOracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIopOracleContract creates a new instance of IopOracleContract, bound to a specific deployed contract.
func NewIopOracleContract(address common.Address, backend bind.ContractBackend) (*IopOracleContract, error) {
	contract, err := bindIopOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IopOracleContract{IopOracleContractCaller: IopOracleContractCaller{contract: contract}, IopOracleContractTransactor: IopOracleContractTransactor{contract: contract}, IopOracleContractFilterer: IopOracleContractFilterer{contract: contract}}, nil
}

// NewIopOracleContractCaller creates a new read-only instance of IopOracleContract, bound to a specific deployed contract.
func NewIopOracleContractCaller(address common.Address, caller bind.ContractCaller) (*IopOracleContractCaller, error) {
	contract, err := bindIopOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IopOracleContractCaller{contract: contract}, nil
}

// NewIopOracleContractTransactor creates a new write-only instance of IopOracleContract, bound to a specific deployed contract.
func NewIopOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*IopOracleContractTransactor, error) {
	contract, err := bindIopOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IopOracleContractTransactor{contract: contract}, nil
}

// NewIopOracleContractFilterer creates a new log filterer instance of IopOracleContract, bound to a specific deployed contract.
func NewIopOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*IopOracleContractFilterer, error) {
	contract, err := bindIopOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IopOracleContractFilterer{contract: contract}, nil
}

// bindIopOracleContract binds a generic wrapper to an already deployed contract.
func bindIopOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IopOracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IopOracleContract *IopOracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IopOracleContract.Contract.IopOracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IopOracleContract *IopOracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IopOracleContract.Contract.IopOracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IopOracleContract *IopOracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IopOracleContract.Contract.IopOracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IopOracleContract *IopOracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IopOracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IopOracleContract *IopOracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IopOracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IopOracleContract *IopOracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IopOracleContract.Contract.contract.Transact(opts, method, params...)
}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracleContract *IopOracleContractCaller) CountIopNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IopOracleContract.contract.Call(opts, &out, "countIopNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracleContract *IopOracleContractSession) CountIopNodes() (*big.Int, error) {
	return _IopOracleContract.Contract.CountIopNodes(&_IopOracleContract.CallOpts)
}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracleContract *IopOracleContractCallerSession) CountIopNodes() (*big.Int, error) {
	return _IopOracleContract.Contract.CountIopNodes(&_IopOracleContract.CallOpts)
}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns(address, string)
func (_IopOracleContract *IopOracleContractCaller) FindIopNodeByAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, string, error) {
	var out []interface{}
	err := _IopOracleContract.contract.Call(opts, &out, "findIopNodeByAddress", _addr)

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
func (_IopOracleContract *IopOracleContractSession) FindIopNodeByAddress(_addr common.Address) (common.Address, string, error) {
	return _IopOracleContract.Contract.FindIopNodeByAddress(&_IopOracleContract.CallOpts, _addr)
}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns(address, string)
func (_IopOracleContract *IopOracleContractCallerSession) FindIopNodeByAddress(_addr common.Address) (common.Address, string, error) {
	return _IopOracleContract.Contract.FindIopNodeByAddress(&_IopOracleContract.CallOpts, _addr)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns(address, string)
func (_IopOracleContract *IopOracleContractCaller) FindIopNodeByIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, string, error) {
	var out []interface{}
	err := _IopOracleContract.contract.Call(opts, &out, "findIopNodeByIndex", _index)

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
func (_IopOracleContract *IopOracleContractSession) FindIopNodeByIndex(_index *big.Int) (common.Address, string, error) {
	return _IopOracleContract.Contract.FindIopNodeByIndex(&_IopOracleContract.CallOpts, _index)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns(address, string)
func (_IopOracleContract *IopOracleContractCallerSession) FindIopNodeByIndex(_index *big.Int) (common.Address, string, error) {
	return _IopOracleContract.Contract.FindIopNodeByIndex(&_IopOracleContract.CallOpts, _index)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracleContract *IopOracleContractCaller) IopNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _IopOracleContract.contract.Call(opts, &out, "iopNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracleContract *IopOracleContractSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _IopOracleContract.Contract.IopNodeIsRegistered(&_IopOracleContract.CallOpts, _addr)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracleContract *IopOracleContractCallerSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _IopOracleContract.Contract.IopNodeIsRegistered(&_IopOracleContract.CallOpts, _addr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracleContract *IopOracleContractTransactor) RegisterIopNode(opts *bind.TransactOpts, _ipAddr string) (*types.Transaction, error) {
	return _IopOracleContract.contract.Transact(opts, "registerIopNode", _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracleContract *IopOracleContractSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _IopOracleContract.Contract.RegisterIopNode(&_IopOracleContract.TransactOpts, _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracleContract *IopOracleContractTransactorSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _IopOracleContract.Contract.RegisterIopNode(&_IopOracleContract.TransactOpts, _ipAddr)
}

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

// IopOracleABI is the input ABI used to generate the binding from.
const IopOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countIopNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findIopNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"iopNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerIopNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IopOracleFuncSigs maps the 4-byte function signature to its string representation.
var IopOracleFuncSigs = map[string]string{
	"f1551da7": "countIopNodes()",
	"40830b23": "findIopNode(address)",
	"338780e0": "iopNodeIsRegistered(address)",
	"6f1daed3": "registerIopNode(string)",
}

// IopOracleBin is the compiled bytecode used for deploying new contracts.
var IopOracleBin = "0x608060405234801561001057600080fd5b5061054b806100206000396000f3fe60806040526004361061003f5760003560e01c8063338780e01461004457806340830b231461008b5780636f1daed31461014f578063f1551da7146101c1575b600080fd5b34801561005057600080fd5b506100776004803603602081101561006757600080fd5b50356001600160a01b03166101e8565b604080519115158252519081900360200190f35b34801561009757600080fd5b506100be600480360360208110156100ae57600080fd5b50356001600160a01b0316610243565b60405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101135781810151838201526020016100fb565b50505050905090810190601f1680156101405780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6101bf6004803603602081101561016557600080fd5b81019060208101813564010000000081111561018057600080fd5b82018360208201111561019257600080fd5b803590602001918460018302840111640100000000831117156101b457600080fd5b509092509050610372565b005b3480156101cd57600080fd5b506101d661044d565b60408051918252519081900360200190f35b6001546000906101fa5750600061023e565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061022557fe5b6000918252602090912001546001600160a01b03161490505b919050565b60006060610250836101e8565b61028d576040805162461bcd60e51b81526020600482015260096024820152681b9bdd08199bdd5b9960ba1b604482015290519081900360640190fd5b610295610454565b6001600160a01b0384811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f810185900485028301850190955284825291938584019391929183018282801561034d5780601f106103225761010080835404028352916020019161034d565b820191906000526020600020905b81548152906001019060200180831161033057829003601f168201915b5050509183525050600291909101546020918201528151910151909350915050915091565b61037b336101e8565b156103c2576040805162461bcd60e51b8152602060048201526012602482015271185b1c9958591e481c9959da5cdd195c995960721b604482015290519081900360640190fd5b33600081815260208190526040902080546001600160a01b03191690911781556103f060018201848461047e565b5080546001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6810180546001600160a01b0319166001600160a01b03909316929092179091556002909101555050565b6001545b90565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106104bf5782800160ff198235161785556104ec565b828001600101855582156104ec579182015b828111156104ec5782358255916020019190600101906104d1565b506104f89291506104fc565b5090565b61045191905b808211156104f8576000815560010161050256fea265627a7a723158206cd47da7a90d7257103f034a8d5f854664235974679fdbc2f3b5d960e087663564736f6c63430005100032"

// DeployIopOracle deploys a new Ethereum contract, binding an instance of IopOracle to it.
func DeployIopOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IopOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(IopOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IopOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IopOracle{IopOracleCaller: IopOracleCaller{contract: contract}, IopOracleTransactor: IopOracleTransactor{contract: contract}, IopOracleFilterer: IopOracleFilterer{contract: contract}}, nil
}

// IopOracle is an auto generated Go binding around an Ethereum contract.
type IopOracle struct {
	IopOracleCaller     // Read-only binding to the contract
	IopOracleTransactor // Write-only binding to the contract
	IopOracleFilterer   // Log filterer for contract events
}

// IopOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IopOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IopOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IopOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IopOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IopOracleSession struct {
	Contract     *IopOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IopOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IopOracleCallerSession struct {
	Contract *IopOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IopOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IopOracleTransactorSession struct {
	Contract     *IopOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IopOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IopOracleRaw struct {
	Contract *IopOracle // Generic contract binding to access the raw methods on
}

// IopOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IopOracleCallerRaw struct {
	Contract *IopOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IopOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IopOracleTransactorRaw struct {
	Contract *IopOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIopOracle creates a new instance of IopOracle, bound to a specific deployed contract.
func NewIopOracle(address common.Address, backend bind.ContractBackend) (*IopOracle, error) {
	contract, err := bindIopOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IopOracle{IopOracleCaller: IopOracleCaller{contract: contract}, IopOracleTransactor: IopOracleTransactor{contract: contract}, IopOracleFilterer: IopOracleFilterer{contract: contract}}, nil
}

// NewIopOracleCaller creates a new read-only instance of IopOracle, bound to a specific deployed contract.
func NewIopOracleCaller(address common.Address, caller bind.ContractCaller) (*IopOracleCaller, error) {
	contract, err := bindIopOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IopOracleCaller{contract: contract}, nil
}

// NewIopOracleTransactor creates a new write-only instance of IopOracle, bound to a specific deployed contract.
func NewIopOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IopOracleTransactor, error) {
	contract, err := bindIopOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IopOracleTransactor{contract: contract}, nil
}

// NewIopOracleFilterer creates a new log filterer instance of IopOracle, bound to a specific deployed contract.
func NewIopOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IopOracleFilterer, error) {
	contract, err := bindIopOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IopOracleFilterer{contract: contract}, nil
}

// bindIopOracle binds a generic wrapper to an already deployed contract.
func bindIopOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IopOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IopOracle *IopOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IopOracle.Contract.IopOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IopOracle *IopOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IopOracle.Contract.IopOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IopOracle *IopOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IopOracle.Contract.IopOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IopOracle *IopOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IopOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IopOracle *IopOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IopOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IopOracle *IopOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IopOracle.Contract.contract.Transact(opts, method, params...)
}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracle *IopOracleCaller) CountIopNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IopOracle.contract.Call(opts, &out, "countIopNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracle *IopOracleSession) CountIopNodes() (*big.Int, error) {
	return _IopOracle.Contract.CountIopNodes(&_IopOracle.CallOpts)
}

// CountIopNodes is a free data retrieval call binding the contract method 0xf1551da7.
//
// Solidity: function countIopNodes() view returns(uint256)
func (_IopOracle *IopOracleCallerSession) CountIopNodes() (*big.Int, error) {
	return _IopOracle.Contract.CountIopNodes(&_IopOracle.CallOpts)
}

// FindIopNode is a free data retrieval call binding the contract method 0x40830b23.
//
// Solidity: function findIopNode(address _addr) view returns(address, string)
func (_IopOracle *IopOracleCaller) FindIopNode(opts *bind.CallOpts, _addr common.Address) (common.Address, string, error) {
	var out []interface{}
	err := _IopOracle.contract.Call(opts, &out, "findIopNode", _addr)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// FindIopNode is a free data retrieval call binding the contract method 0x40830b23.
//
// Solidity: function findIopNode(address _addr) view returns(address, string)
func (_IopOracle *IopOracleSession) FindIopNode(_addr common.Address) (common.Address, string, error) {
	return _IopOracle.Contract.FindIopNode(&_IopOracle.CallOpts, _addr)
}

// FindIopNode is a free data retrieval call binding the contract method 0x40830b23.
//
// Solidity: function findIopNode(address _addr) view returns(address, string)
func (_IopOracle *IopOracleCallerSession) FindIopNode(_addr common.Address) (common.Address, string, error) {
	return _IopOracle.Contract.FindIopNode(&_IopOracle.CallOpts, _addr)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracle *IopOracleCaller) IopNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _IopOracle.contract.Call(opts, &out, "iopNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracle *IopOracleSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _IopOracle.Contract.IopNodeIsRegistered(&_IopOracle.CallOpts, _addr)
}

// IopNodeIsRegistered is a free data retrieval call binding the contract method 0x338780e0.
//
// Solidity: function iopNodeIsRegistered(address _addr) view returns(bool)
func (_IopOracle *IopOracleCallerSession) IopNodeIsRegistered(_addr common.Address) (bool, error) {
	return _IopOracle.Contract.IopNodeIsRegistered(&_IopOracle.CallOpts, _addr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracle *IopOracleTransactor) RegisterIopNode(opts *bind.TransactOpts, _ipAddr string) (*types.Transaction, error) {
	return _IopOracle.contract.Transact(opts, "registerIopNode", _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracle *IopOracleSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _IopOracle.Contract.RegisterIopNode(&_IopOracle.TransactOpts, _ipAddr)
}

// RegisterIopNode is a paid mutator transaction binding the contract method 0x6f1daed3.
//
// Solidity: function registerIopNode(string _ipAddr) payable returns()
func (_IopOracle *IopOracleTransactorSession) RegisterIopNode(_ipAddr string) (*types.Transaction, error) {
	return _IopOracle.Contract.RegisterIopNode(&_IopOracle.TransactOpts, _ipAddr)
}

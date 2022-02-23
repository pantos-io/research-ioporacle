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

// RegistryContractOracleNode is an auto generated low-level Go binding around an user-defined struct.
type RegistryContractOracleNode struct {
	Addr   common.Address
	IpAddr string
	PubKey []byte
	Stake  *big.Int
	Index  *big.Int
}

// RegistryContractMetaData contains all meta data concerning the RegistryContract contract.
var RegistryContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distKeyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RegisterOracleNode\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findOracleNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistryContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistryContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistryContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"getAggregatorByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistryContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAggregator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"isAggregatorByBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_pubKey\",\"type\":\"bytes\"}],\"name\":\"registerOracleNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryContractMetaData.ABI instead.
var RegistryContractABI = RegistryContractMetaData.ABI

// RegistryContract is an auto generated Go binding around an Ethereum contract.
type RegistryContract struct {
	RegistryContractCaller     // Read-only binding to the contract
	RegistryContractTransactor // Write-only binding to the contract
	RegistryContractFilterer   // Log filterer for contract events
}

// RegistryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistryContractSession struct {
	Contract     *RegistryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryContractCallerSession struct {
	Contract *RegistryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RegistryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryContractTransactorSession struct {
	Contract     *RegistryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RegistryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryContractRaw struct {
	Contract *RegistryContract // Generic contract binding to access the raw methods on
}

// RegistryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryContractCallerRaw struct {
	Contract *RegistryContractCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryContractTransactorRaw struct {
	Contract *RegistryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryContract creates a new instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContract(address common.Address, backend bind.ContractBackend) (*RegistryContract, error) {
	contract, err := bindRegistryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryContract{RegistryContractCaller: RegistryContractCaller{contract: contract}, RegistryContractTransactor: RegistryContractTransactor{contract: contract}, RegistryContractFilterer: RegistryContractFilterer{contract: contract}}, nil
}

// NewRegistryContractCaller creates a new read-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractCaller(address common.Address, caller bind.ContractCaller) (*RegistryContractCaller, error) {
	contract, err := bindRegistryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractCaller{contract: contract}, nil
}

// NewRegistryContractTransactor creates a new write-only instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryContractTransactor, error) {
	contract, err := bindRegistryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryContractTransactor{contract: contract}, nil
}

// NewRegistryContractFilterer creates a new log filterer instance of RegistryContract, bound to a specific deployed contract.
func NewRegistryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryContractFilterer, error) {
	contract, err := bindRegistryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryContractFilterer{contract: contract}, nil
}

// bindRegistryContract binds a generic wrapper to an already deployed contract.
func bindRegistryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.RegistryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegistryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryContract *RegistryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryContract *RegistryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryContract *RegistryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryContract.Contract.contract.Transact(opts, method, params...)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_RegistryContract *RegistryContractCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_RegistryContract *RegistryContractSession) MINSTAKE() (*big.Int, error) {
	return _RegistryContract.Contract.MINSTAKE(&_RegistryContract.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_RegistryContract *RegistryContractCallerSession) MINSTAKE() (*big.Int, error) {
	return _RegistryContract.Contract.MINSTAKE(&_RegistryContract.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_RegistryContract *RegistryContractCaller) CountOracleNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "countOracleNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_RegistryContract *RegistryContractSession) CountOracleNodes() (*big.Int, error) {
	return _RegistryContract.Contract.CountOracleNodes(&_RegistryContract.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_RegistryContract *RegistryContractCallerSession) CountOracleNodes() (*big.Int, error) {
	return _RegistryContract.Contract.CountOracleNodes(&_RegistryContract.CallOpts)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCaller) FindOracleNodeByAddress(opts *bind.CallOpts, _addr common.Address) (RegistryContractOracleNode, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "findOracleNodeByAddress", _addr)

	if err != nil {
		return *new(RegistryContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryContractOracleNode)).(*RegistryContractOracleNode)

	return out0, err

}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractSession) FindOracleNodeByAddress(_addr common.Address) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.FindOracleNodeByAddress(&_RegistryContract.CallOpts, _addr)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCallerSession) FindOracleNodeByAddress(_addr common.Address) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.FindOracleNodeByAddress(&_RegistryContract.CallOpts, _addr)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCaller) FindOracleNodeByIndex(opts *bind.CallOpts, _index *big.Int) (RegistryContractOracleNode, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "findOracleNodeByIndex", _index)

	if err != nil {
		return *new(RegistryContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryContractOracleNode)).(*RegistryContractOracleNode)

	return out0, err

}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractSession) FindOracleNodeByIndex(_index *big.Int) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.FindOracleNodeByIndex(&_RegistryContract.CallOpts, _index)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCallerSession) FindOracleNodeByIndex(_index *big.Int) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.FindOracleNodeByIndex(&_RegistryContract.CallOpts, _index)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCaller) GetAggregator(opts *bind.CallOpts) (RegistryContractOracleNode, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getAggregator")

	if err != nil {
		return *new(RegistryContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryContractOracleNode)).(*RegistryContractOracleNode)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractSession) GetAggregator() (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.GetAggregator(&_RegistryContract.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCallerSession) GetAggregator() (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.GetAggregator(&_RegistryContract.CallOpts)
}

// GetAggregatorByBlock is a free data retrieval call binding the contract method 0x59798dfd.
//
// Solidity: function getAggregatorByBlock(uint256 _block) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCaller) GetAggregatorByBlock(opts *bind.CallOpts, _block *big.Int) (RegistryContractOracleNode, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "getAggregatorByBlock", _block)

	if err != nil {
		return *new(RegistryContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryContractOracleNode)).(*RegistryContractOracleNode)

	return out0, err

}

// GetAggregatorByBlock is a free data retrieval call binding the contract method 0x59798dfd.
//
// Solidity: function getAggregatorByBlock(uint256 _block) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractSession) GetAggregatorByBlock(_block *big.Int) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.GetAggregatorByBlock(&_RegistryContract.CallOpts, _block)
}

// GetAggregatorByBlock is a free data retrieval call binding the contract method 0x59798dfd.
//
// Solidity: function getAggregatorByBlock(uint256 _block) view returns((address,string,bytes,uint256,uint256))
func (_RegistryContract *RegistryContractCallerSession) GetAggregatorByBlock(_block *big.Int) (RegistryContractOracleNode, error) {
	return _RegistryContract.Contract.GetAggregatorByBlock(&_RegistryContract.CallOpts, _block)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCaller) IsAggregator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "isAggregator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractSession) IsAggregator(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.IsAggregator(&_RegistryContract.CallOpts, _addr)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCallerSession) IsAggregator(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.IsAggregator(&_RegistryContract.CallOpts, _addr)
}

// IsAggregatorByBlock is a free data retrieval call binding the contract method 0x68982668.
//
// Solidity: function isAggregatorByBlock(address _addr, uint256 _block) view returns(bool)
func (_RegistryContract *RegistryContractCaller) IsAggregatorByBlock(opts *bind.CallOpts, _addr common.Address, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "isAggregatorByBlock", _addr, _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAggregatorByBlock is a free data retrieval call binding the contract method 0x68982668.
//
// Solidity: function isAggregatorByBlock(address _addr, uint256 _block) view returns(bool)
func (_RegistryContract *RegistryContractSession) IsAggregatorByBlock(_addr common.Address, _block *big.Int) (bool, error) {
	return _RegistryContract.Contract.IsAggregatorByBlock(&_RegistryContract.CallOpts, _addr, _block)
}

// IsAggregatorByBlock is a free data retrieval call binding the contract method 0x68982668.
//
// Solidity: function isAggregatorByBlock(address _addr, uint256 _block) view returns(bool)
func (_RegistryContract *RegistryContractCallerSession) IsAggregatorByBlock(_addr common.Address, _block *big.Int) (bool, error) {
	return _RegistryContract.Contract.IsAggregatorByBlock(&_RegistryContract.CallOpts, _addr, _block)
}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCaller) OracleNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryContract.contract.Call(opts, &out, "oracleNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.OracleNodeIsRegistered(&_RegistryContract.CallOpts, _addr)
}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_RegistryContract *RegistryContractCallerSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _RegistryContract.Contract.OracleNodeIsRegistered(&_RegistryContract.CallOpts, _addr)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0xdf74f12c.
//
// Solidity: function registerOracleNode(string _ipAddr, bytes _pubKey) payable returns()
func (_RegistryContract *RegistryContractTransactor) RegisterOracleNode(opts *bind.TransactOpts, _ipAddr string, _pubKey []byte) (*types.Transaction, error) {
	return _RegistryContract.contract.Transact(opts, "registerOracleNode", _ipAddr, _pubKey)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0xdf74f12c.
//
// Solidity: function registerOracleNode(string _ipAddr, bytes _pubKey) payable returns()
func (_RegistryContract *RegistryContractSession) RegisterOracleNode(_ipAddr string, _pubKey []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegisterOracleNode(&_RegistryContract.TransactOpts, _ipAddr, _pubKey)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0xdf74f12c.
//
// Solidity: function registerOracleNode(string _ipAddr, bytes _pubKey) payable returns()
func (_RegistryContract *RegistryContractTransactorSession) RegisterOracleNode(_ipAddr string, _pubKey []byte) (*types.Transaction, error) {
	return _RegistryContract.Contract.RegisterOracleNode(&_RegistryContract.TransactOpts, _ipAddr, _pubKey)
}

// RegistryContractRegisterOracleNodeIterator is returned from FilterRegisterOracleNode and is used to iterate over the raw logs and unpacked data for RegisterOracleNode events raised by the RegistryContract contract.
type RegistryContractRegisterOracleNodeIterator struct {
	Event *RegistryContractRegisterOracleNode // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegisterOracleNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegisterOracleNode)
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
		it.Event = new(RegistryContractRegisterOracleNode)
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
func (it *RegistryContractRegisterOracleNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegisterOracleNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegisterOracleNode represents a RegisterOracleNode event raised by the RegistryContract contract.
type RegistryContractRegisterOracleNode struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterOracleNode is a free log retrieval operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_RegistryContract *RegistryContractFilterer) FilterRegisterOracleNode(opts *bind.FilterOpts, sender []common.Address) (*RegistryContractRegisterOracleNodeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryContract.contract.FilterLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegisterOracleNodeIterator{contract: _RegistryContract.contract, event: "RegisterOracleNode", logs: logs, sub: sub}, nil
}

// WatchRegisterOracleNode is a free log subscription operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_RegistryContract *RegistryContractFilterer) WatchRegisterOracleNode(opts *bind.WatchOpts, sink chan<- *RegistryContractRegisterOracleNode, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RegistryContract.contract.WatchLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegisterOracleNode)
				if err := _RegistryContract.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
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

// ParseRegisterOracleNode is a log parse operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_RegistryContract *RegistryContractFilterer) ParseRegisterOracleNode(log types.Log) (*RegistryContractRegisterOracleNode, error) {
	event := new(RegistryContractRegisterOracleNode)
	if err := _RegistryContract.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

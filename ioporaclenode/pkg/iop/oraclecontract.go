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

// OracleContractValidationResult is an auto generated low-level Go binding around an user-defined struct.
type OracleContractValidationResult struct {
	Hash   [32]byte
	Result bool
}

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distKeyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SubmitValidationResultLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"ValidateBlockLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"ValidateTransactionLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"blockValidationResultExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findBlockValidationResult\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.ValidationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"findTxValidationResult\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.ValidationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isBlockConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"}],\"name\":\"isValidationFeeReceiver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitBlockValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"txValidationResultExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInWei\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInWei\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyState\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInWei\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyTransaction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

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

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) TOTALFEE() (*big.Int, error) {
	return _OracleContract.Contract.TOTALFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) VALIDATORFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "VALIDATOR_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// VALIDATORFEE is a free data retrieval call binding the contract method 0x7da83e2b.
//
// Solidity: function VALIDATOR_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) VALIDATORFEE() (*big.Int, error) {
	return _OracleContract.Contract.VALIDATORFEE(&_OracleContract.CallOpts)
}

// BlockValidationResultExists is a free data retrieval call binding the contract method 0x72e3de18.
//
// Solidity: function blockValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) BlockValidationResultExists(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "blockValidationResultExists", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockValidationResultExists is a free data retrieval call binding the contract method 0x72e3de18.
//
// Solidity: function blockValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) BlockValidationResultExists(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.BlockValidationResultExists(&_OracleContract.CallOpts, _hash)
}

// BlockValidationResultExists is a free data retrieval call binding the contract method 0x72e3de18.
//
// Solidity: function blockValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) BlockValidationResultExists(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.BlockValidationResultExists(&_OracleContract.CallOpts, _hash)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractCaller) FindBlockValidationResult(opts *bind.CallOpts, _hash [32]byte) (OracleContractValidationResult, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findBlockValidationResult", _hash)

	if err != nil {
		return *new(OracleContractValidationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractValidationResult)).(*OracleContractValidationResult)

	return out0, err

}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractSession) FindBlockValidationResult(_hash [32]byte) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindBlockValidationResult is a free data retrieval call binding the contract method 0x3d27ef97.
//
// Solidity: function findBlockValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractCallerSession) FindBlockValidationResult(_hash [32]byte) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindBlockValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTxValidationResult is a free data retrieval call binding the contract method 0x7b47a02a.
//
// Solidity: function findTxValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractCaller) FindTxValidationResult(opts *bind.CallOpts, _hash [32]byte) (OracleContractValidationResult, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findTxValidationResult", _hash)

	if err != nil {
		return *new(OracleContractValidationResult), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractValidationResult)).(*OracleContractValidationResult)

	return out0, err

}

// FindTxValidationResult is a free data retrieval call binding the contract method 0x7b47a02a.
//
// Solidity: function findTxValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractSession) FindTxValidationResult(_hash [32]byte) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindTxValidationResult(&_OracleContract.CallOpts, _hash)
}

// FindTxValidationResult is a free data retrieval call binding the contract method 0x7b47a02a.
//
// Solidity: function findTxValidationResult(bytes32 _hash) view returns((bytes32,bool))
func (_OracleContract *OracleContractCallerSession) FindTxValidationResult(_hash [32]byte) (OracleContractValidationResult, error) {
	return _OracleContract.Contract.FindTxValidationResult(&_OracleContract.CallOpts, _hash)
}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 stake, uint256 seed) pure returns(bool)
func (_OracleContract *OracleContractCaller) IsValidationFeeReceiver(opts *bind.CallOpts, stake *big.Int, seed *big.Int) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "isValidationFeeReceiver", stake, seed)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 stake, uint256 seed) pure returns(bool)
func (_OracleContract *OracleContractSession) IsValidationFeeReceiver(stake *big.Int, seed *big.Int) (bool, error) {
	return _OracleContract.Contract.IsValidationFeeReceiver(&_OracleContract.CallOpts, stake, seed)
}

// IsValidationFeeReceiver is a free data retrieval call binding the contract method 0xf5ba1d3d.
//
// Solidity: function isValidationFeeReceiver(uint256 stake, uint256 seed) pure returns(bool)
func (_OracleContract *OracleContractCallerSession) IsValidationFeeReceiver(stake *big.Int, seed *big.Int) (bool, error) {
	return _OracleContract.Contract.IsValidationFeeReceiver(&_OracleContract.CallOpts, stake, seed)
}

// TxValidationResultExists is a free data retrieval call binding the contract method 0xdabfc4e6.
//
// Solidity: function txValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCaller) TxValidationResultExists(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "txValidationResultExists", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxValidationResultExists is a free data retrieval call binding the contract method 0xdabfc4e6.
//
// Solidity: function txValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractSession) TxValidationResultExists(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.TxValidationResultExists(&_OracleContract.CallOpts, _hash)
}

// TxValidationResultExists is a free data retrieval call binding the contract method 0xdabfc4e6.
//
// Solidity: function txValidationResultExists(bytes32 _hash) view returns(bool)
func (_OracleContract *OracleContractCallerSession) TxValidationResultExists(_hash [32]byte) (bool, error) {
	return _OracleContract.Contract.TxValidationResultExists(&_OracleContract.CallOpts, _hash)
}

// IsBlockConfirmed is a paid mutator transaction binding the contract method 0x9dc29d89.
//
// Solidity: function isBlockConfirmed(uint256 , bytes32 blockHash, uint256 ) payable returns(bool)
func (_OracleContract *OracleContractTransactor) IsBlockConfirmed(opts *bind.TransactOpts, arg0 *big.Int, blockHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "isBlockConfirmed", arg0, blockHash, arg2)
}

// IsBlockConfirmed is a paid mutator transaction binding the contract method 0x9dc29d89.
//
// Solidity: function isBlockConfirmed(uint256 , bytes32 blockHash, uint256 ) payable returns(bool)
func (_OracleContract *OracleContractSession) IsBlockConfirmed(arg0 *big.Int, blockHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.IsBlockConfirmed(&_OracleContract.TransactOpts, arg0, blockHash, arg2)
}

// IsBlockConfirmed is a paid mutator transaction binding the contract method 0x9dc29d89.
//
// Solidity: function isBlockConfirmed(uint256 , bytes32 blockHash, uint256 ) payable returns(bool)
func (_OracleContract *OracleContractTransactorSession) IsBlockConfirmed(arg0 *big.Int, blockHash [32]byte, arg2 *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.IsBlockConfirmed(&_OracleContract.TransactOpts, arg0, blockHash, arg2)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationResult", _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0x3dd14279.
//
// Solidity: function submitBlockValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x25f08549.
//
// Solidity: function submitTransactionValidationResult(bytes32 _hash, bool _result, uint256[2] _signature) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_hash [32]byte, _result bool, _signature [2]*big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _hash, _result, _signature)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateBlock(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateBlock", _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xdd5e22df.
//
// Solidity: function validateBlock(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateBlock(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _hash) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_hash [32]byte) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _hash)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedReceipt, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactor) VerifyReceipt(opts *bind.TransactOpts, feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedReceipt []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyReceipt", feeInWei, rlpHeader, noOfConfirmations, rlpEncodedReceipt, path, rlpEncodedNodes)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedReceipt, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractSession) VerifyReceipt(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedReceipt []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyReceipt(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedReceipt, path, rlpEncodedNodes)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedReceipt, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactorSession) VerifyReceipt(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedReceipt []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyReceipt(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedReceipt, path, rlpEncodedNodes)
}

// VerifyState is a paid mutator transaction binding the contract method 0xaddd9b38.
//
// Solidity: function verifyState(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedState, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactor) VerifyState(opts *bind.TransactOpts, feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedState []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyState", feeInWei, rlpHeader, noOfConfirmations, rlpEncodedState, path, rlpEncodedNodes)
}

// VerifyState is a paid mutator transaction binding the contract method 0xaddd9b38.
//
// Solidity: function verifyState(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedState, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractSession) VerifyState(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedState []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyState(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedState, path, rlpEncodedNodes)
}

// VerifyState is a paid mutator transaction binding the contract method 0xaddd9b38.
//
// Solidity: function verifyState(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedState, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactorSession) VerifyState(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedState []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyState(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedState, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactor) VerifyTransaction(opts *bind.TransactOpts, feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyTransaction", feeInWei, rlpHeader, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractSession) VerifyTransaction(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes rlpHeader, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_OracleContract *OracleContractTransactorSession) VerifyTransaction(feeInWei *big.Int, rlpHeader []byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, feeInWei, rlpHeader, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
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
	Hash   [32]byte
	Result bool
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitValidationResultLog is a free log retrieval operation binding the contract event 0xf7c4db3475526b6d9da982668512fe2b688df43ef777e449da121e71da154608.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, bytes32 _hash, bool result, uint256 fee)
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

// WatchSubmitValidationResultLog is a free log subscription operation binding the contract event 0xf7c4db3475526b6d9da982668512fe2b688df43ef777e449da121e71da154608.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, bytes32 _hash, bool result, uint256 fee)
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

// ParseSubmitValidationResultLog is a log parse operation binding the contract event 0xf7c4db3475526b6d9da982668512fe2b688df43ef777e449da121e71da154608.
//
// Solidity: event SubmitValidationResultLog(address indexed sender, bytes32 _hash, bool result, uint256 fee)
func (_OracleContract *OracleContractFilterer) ParseSubmitValidationResultLog(log types.Log) (*OracleContractSubmitValidationResultLog, error) {
	event := new(OracleContractSubmitValidationResultLog)
	if err := _OracleContract.contract.UnpackLog(event, "SubmitValidationResultLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidateBlockLogIterator is returned from FilterValidateBlockLog and is used to iterate over the raw logs and unpacked data for ValidateBlockLog events raised by the OracleContract contract.
type OracleContractValidateBlockLogIterator struct {
	Event *OracleContractValidateBlockLog // Event containing the contract specifics and raw log

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
func (it *OracleContractValidateBlockLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidateBlockLog)
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
		it.Event = new(OracleContractValidateBlockLog)
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
func (it *OracleContractValidateBlockLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidateBlockLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidateBlockLog represents a ValidateBlockLog event raised by the OracleContract contract.
type OracleContractValidateBlockLog struct {
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidateBlockLog is a free log retrieval operation binding the contract event 0x7369c156accc2dffc92cb38087ac9ed801757bb90ab4945b0f126b17128a89dd.
//
// Solidity: event ValidateBlockLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) FilterValidateBlockLog(opts *bind.FilterOpts, sender []common.Address) (*OracleContractValidateBlockLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidateBlockLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidateBlockLogIterator{contract: _OracleContract.contract, event: "ValidateBlockLog", logs: logs, sub: sub}, nil
}

// WatchValidateBlockLog is a free log subscription operation binding the contract event 0x7369c156accc2dffc92cb38087ac9ed801757bb90ab4945b0f126b17128a89dd.
//
// Solidity: event ValidateBlockLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) WatchValidateBlockLog(opts *bind.WatchOpts, sink chan<- *OracleContractValidateBlockLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidateBlockLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidateBlockLog)
				if err := _OracleContract.contract.UnpackLog(event, "ValidateBlockLog", log); err != nil {
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

// ParseValidateBlockLog is a log parse operation binding the contract event 0x7369c156accc2dffc92cb38087ac9ed801757bb90ab4945b0f126b17128a89dd.
//
// Solidity: event ValidateBlockLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) ParseValidateBlockLog(log types.Log) (*OracleContractValidateBlockLog, error) {
	event := new(OracleContractValidateBlockLog)
	if err := _OracleContract.contract.UnpackLog(event, "ValidateBlockLog", log); err != nil {
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
	Sender common.Address
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidateTransactionLog is a free log retrieval operation binding the contract event 0x0d6f8b3dd2cfc880fced8136729f786c893eb0a5776ccc9c0d6d734f77fdbc6e.
//
// Solidity: event ValidateTransactionLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) FilterValidateTransactionLog(opts *bind.FilterOpts, sender []common.Address) (*OracleContractValidateTransactionLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidateTransactionLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidateTransactionLogIterator{contract: _OracleContract.contract, event: "ValidateTransactionLog", logs: logs, sub: sub}, nil
}

// WatchValidateTransactionLog is a free log subscription operation binding the contract event 0x0d6f8b3dd2cfc880fced8136729f786c893eb0a5776ccc9c0d6d734f77fdbc6e.
//
// Solidity: event ValidateTransactionLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) WatchValidateTransactionLog(opts *bind.WatchOpts, sink chan<- *OracleContractValidateTransactionLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidateTransactionLog", senderRule)
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

// ParseValidateTransactionLog is a log parse operation binding the contract event 0x0d6f8b3dd2cfc880fced8136729f786c893eb0a5776ccc9c0d6d734f77fdbc6e.
//
// Solidity: event ValidateTransactionLog(address indexed sender, bytes32 _hash)
func (_OracleContract *OracleContractFilterer) ParseValidateTransactionLog(log types.Log) (*OracleContractValidateTransactionLog, error) {
	event := new(OracleContractValidateTransactionLog)
	if err := _OracleContract.contract.UnpackLog(event, "ValidateTransactionLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

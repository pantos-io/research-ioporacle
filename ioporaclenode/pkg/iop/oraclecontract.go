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

// OracleContractIopNode is an auto generated low-level Go binding around an user-defined struct.
type OracleContractIopNode struct {
	Addr   common.Address
	IpAddr string
	Index  *big.Int
}

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"Test\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"countIopNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findIopNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.IopNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findIopNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.IopNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"iopNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerIopNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"result\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"verifyTransactionResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleContractFuncSigs maps the 4-byte function signature to its string representation.
var OracleContractFuncSigs = map[string]string{
	"f1551da7": "countIopNodes()",
	"76ac6e0b": "findIopNodeByAddress(address)",
	"7206c984": "findIopNodeByIndex(uint256)",
	"338780e0": "iopNodeIsRegistered(address)",
	"db512e85": "isLeader(address)",
	"6f1daed3": "registerIopNode(string)",
	"65372147": "result()",
	"ff84d3e9": "verifyTransaction(string,uint256)",
	"fd9f77ec": "verifyTransactionResult(bool)",
}

// OracleContractBin is the compiled bytecode used for deploying new contracts.
var OracleContractBin = "0x608060405234801561001057600080fd5b50610a25806100206000396000f3fe6080604052600436106100865760003560e01c806376ac6e0b1161005957806376ac6e0b14610118578063db512e8514610138578063f1551da714610158578063fd9f77ec1461017a578063ff84d3e91461019a57610086565b8063338780e01461008b57806365372147146100c15780636f1daed3146100d65780637206c984146100eb575b600080fd5b34801561009757600080fd5b506100ab6100a63660046106a1565b6101ba565b6040516100b891906108b2565b60405180910390f35b3480156100cd57600080fd5b506100ab610215565b6100e96100e43660046106dd565b61021e565b005b3480156100f757600080fd5b5061010b610106366004610774565b6102d9565b6040516100b89190610919565b34801561012457600080fd5b5061010b6101333660046106a1565b6103fe565b34801561014457600080fd5b506100ab6101533660046106a1565b6104c0565b34801561016457600080fd5b5061016d610512565b6040516100b89190610931565b34801561018657600080fd5b506100e96101953660046106bf565b610519565b3480156101a657600080fd5b506100e96101b536600461071f565b61052c565b6001546000906101cc57506000610210565b6001600160a01b0382166000818152602081905260409020600201546001805490919081106101f757fe5b6000918252602090912001546001600160a01b03161490505b919050565b60025460ff1681565b610227336101ba565b1561024d5760405162461bcd60e51b8152600401610244906108f9565b60405180910390fd5b33600081815260208190526040902080546001600160a01b031916909117815561027b60018201848461056f565b506001805460028301819055915482820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690910180546001600160a01b0319166001600160a01b039092169190911790555050565b6102e16105ed565b60015482106103025760405162461bcd60e51b815260040161024490610909565b61030a6105ed565b6000806001858154811061031a57fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156103e25780601f106103b7576101008083540402835291602001916103e2565b820191906000526020600020905b8154815290600101906020018083116103c557829003601f168201915b5050509183525050600291909101546020909101529392505050565b6104066105ed565b61040f826101ba565b61042b5760405162461bcd60e51b815260040161024490610909565b6104336105ed565b6001600160a01b0383811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156103e25780601f106103b7576101008083540402835291602001916103e2565b60006104ca6105ed565b6104d3836103fe565b60015490915060009060060243816104e757fe5b0690506006826040015102811015801561050a5750600682604001516001010281105b949350505050565b6001545b90565b6002805460ff1916911515919091179055565b7fd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d28600184848460405161056294939291906108c0565b60405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105b05782800160ff198235161785556105dd565b828001600101855582156105dd579182015b828111156105dd5782358255916020019190600101906105c2565b506105e9929150610617565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b61051691905b808211156105e9576000815560010161061d565b803561063c816109b9565b92915050565b803561063c816109d0565b60008083601f84011261065f57600080fd5b50813567ffffffffffffffff81111561067757600080fd5b60208301915083600182028301111561068f57600080fd5b9250929050565b803561063c816109d9565b6000602082840312156106b357600080fd5b600061050a8484610631565b6000602082840312156106d157600080fd5b600061050a8484610642565b600080602083850312156106f057600080fd5b823567ffffffffffffffff81111561070757600080fd5b6107138582860161064d565b92509250509250929050565b60008060006040848603121561073457600080fd5b833567ffffffffffffffff81111561074b57600080fd5b6107578682870161064d565b9350935050602061076a86828701610696565b9150509250925092565b60006020828403121561078657600080fd5b600061050a8484610696565b61079b8161094c565b82525050565b61079b81610957565b61079b81610968565b60006107bf8385610943565b93506107cc838584610973565b6107d5836109af565b9093019392505050565b60006107ea8261093f565b6107f48185610943565b935061080481856020860161097f565b6107d5816109af565b600061081a601283610943565b71185b1c9958591e481c9959da5cdd195c995960721b815260200192915050565b6000610848600983610943565b681b9bdd08199bdd5b9960ba1b815260200192915050565b805160009060608401906108748582610792565b506020830151848203602086015261088c82826107df565b91505060408301516108a160408601826108a9565b509392505050565b61079b81610516565b6020810161063c82846107a1565b606081016108ce82876107aa565b81810360208301526108e18185876107b3565b90506108f060408301846108a9565b95945050505050565b6020808252810161063c8161080d565b6020808252810161063c8161083b565b6020808252810161092a8184610860565b9392505050565b6020810161063c82846108a9565b5190565b90815260200190565b600061063c8261095c565b151590565b6001600160a01b031690565b600061063c82610516565b82818337506000910152565b60005b8381101561099a578181015183820152602001610982565b838111156109a9576000848401525b50505050565b601f01601f191690565b6109c28161094c565b81146109cd57600080fd5b50565b6109c281610957565b6109c28161051656fea365627a7a723158204bf299544188764794841a157bbe408e9bcb09dafef1988d9df35f62586745376c6578706572696d656e74616cf564736f6c63430005100040"

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
// Solidity: function findIopNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractCaller) FindIopNodeByAddress(opts *bind.CallOpts, _addr common.Address) (OracleContractIopNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findIopNodeByAddress", _addr)

	if err != nil {
		return *new(OracleContractIopNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractIopNode)).(*OracleContractIopNode)

	return out0, err

}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractSession) FindIopNodeByAddress(_addr common.Address) (OracleContractIopNode, error) {
	return _OracleContract.Contract.FindIopNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindIopNodeByAddress is a free data retrieval call binding the contract method 0x76ac6e0b.
//
// Solidity: function findIopNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractCallerSession) FindIopNodeByAddress(_addr common.Address) (OracleContractIopNode, error) {
	return _OracleContract.Contract.FindIopNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractCaller) FindIopNodeByIndex(opts *bind.CallOpts, _index *big.Int) (OracleContractIopNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findIopNodeByIndex", _index)

	if err != nil {
		return *new(OracleContractIopNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractIopNode)).(*OracleContractIopNode)

	return out0, err

}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractSession) FindIopNodeByIndex(_index *big.Int) (OracleContractIopNode, error) {
	return _OracleContract.Contract.FindIopNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindIopNodeByIndex is a free data retrieval call binding the contract method 0x7206c984.
//
// Solidity: function findIopNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractCallerSession) FindIopNodeByIndex(_index *big.Int) (OracleContractIopNode, error) {
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

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address addr) view returns(bool)
func (_OracleContract *OracleContractCaller) IsLeader(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "isLeader", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address addr) view returns(bool)
func (_OracleContract *OracleContractSession) IsLeader(addr common.Address) (bool, error) {
	return _OracleContract.Contract.IsLeader(&_OracleContract.CallOpts, addr)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) IsLeader(addr common.Address) (bool, error) {
	return _OracleContract.Contract.IsLeader(&_OracleContract.CallOpts, addr)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(bool)
func (_OracleContract *OracleContractCaller) Result(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "result")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(bool)
func (_OracleContract *OracleContractSession) Result() (bool, error) {
	return _OracleContract.Contract.Result(&_OracleContract.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() view returns(bool)
func (_OracleContract *OracleContractCallerSession) Result() (bool, error) {
	return _OracleContract.Contract.Result(&_OracleContract.CallOpts)
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

// OracleContractTestIterator is returned from FilterTest and is used to iterate over the raw logs and unpacked data for Test events raised by the OracleContract contract.
type OracleContractTestIterator struct {
	Event *OracleContractTest // Event containing the contract specifics and raw log

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
func (it *OracleContractTestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractTest)
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
		it.Event = new(OracleContractTest)
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
func (it *OracleContractTestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractTestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractTest represents a Test event raised by the OracleContract contract.
type OracleContractTest struct {
	A   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTest is a free log retrieval operation binding the contract event 0x63a242a632efe33c0e210e04e4173612a17efa4f16aa4890bc7e46caece80de0.
//
// Solidity: event Test(uint256 a)
func (_OracleContract *OracleContractFilterer) FilterTest(opts *bind.FilterOpts) (*OracleContractTestIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return &OracleContractTestIterator{contract: _OracleContract.contract, event: "Test", logs: logs, sub: sub}, nil
}

// WatchTest is a free log subscription operation binding the contract event 0x63a242a632efe33c0e210e04e4173612a17efa4f16aa4890bc7e46caece80de0.
//
// Solidity: event Test(uint256 a)
func (_OracleContract *OracleContractFilterer) WatchTest(opts *bind.WatchOpts, sink chan<- *OracleContractTest) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractTest)
				if err := _OracleContract.contract.UnpackLog(event, "Test", log); err != nil {
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

// ParseTest is a log parse operation binding the contract event 0x63a242a632efe33c0e210e04e4173612a17efa4f16aa4890bc7e46caece80de0.
//
// Solidity: event Test(uint256 a)
func (_OracleContract *OracleContractFilterer) ParseTest(log types.Log) (*OracleContractTest, error) {
	event := new(OracleContractTest)
	if err := _OracleContract.contract.UnpackLog(event, "Test", log); err != nil {
		return nil, err
	}
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

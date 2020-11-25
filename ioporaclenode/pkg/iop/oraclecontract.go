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

// OracleContractOracleNode is an auto generated low-level Go binding around an user-defined struct.
type OracleContractOracleNode struct {
	Addr   common.Address
	IpAddr string
	Index  *big.Int
}

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"}],\"name\":\"SubmitVerificationLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"encodeResult\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findOracleNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerOracleNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_witnesses\",\"type\":\"address[]\"}],\"name\":\"submitVerification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleContractFuncSigs maps the 4-byte function signature to its string representation.
var OracleContractFuncSigs = map[string]string{
	"836f187a": "countOracleNodes()",
	"b33d96bd": "encodeResult(uint256,bool)",
	"655a102f": "findOracleNodeByAddress(address)",
	"272132e9": "findOracleNodeByIndex(uint256)",
	"db512e85": "isLeader(address)",
	"140f3daa": "oracleNodeIsRegistered(address)",
	"67cf3e48": "registerOracleNode(string)",
	"18176c05": "submitVerification(uint256,bool,address[])",
	"ff84d3e9": "verifyTransaction(string,uint256)",
}

// OracleContractBin is the compiled bytecode used for deploying new contracts.
var OracleContractBin = "0x608060405234801561001057600080fd5b50610dda806100206000396000f3fe6080604052600436106100865760003560e01c806367cf3e481161005957806367cf3e4814610130578063836f187a14610143578063b33d96bd14610165578063db512e8514610192578063ff84d3e9146101b257610086565b8063140f3daa1461008b57806318176c05146100c1578063272132e9146100e3578063655a102f14610110575b600080fd5b34801561009757600080fd5b506100ab6100a636600461089b565b6101d2565b6040516100b89190610bcf565b60405180910390f35b3480156100cd57600080fd5b506100e16100dc3660046109a8565b61022d565b005b3480156100ef57600080fd5b506101036100fe366004610950565b6102a8565b6040516100b89190610c15565b34801561011c57600080fd5b5061010361012b36600461089b565b6103d6565b6100e161013e3660046108b9565b610498565b34801561014f57600080fd5b5061015861054a565b6040516100b89190610c26565b34801561017157600080fd5b5061018561018036600461096e565b610551565b6040516100b89190610bdd565b34801561019e57600080fd5b506100ab6101ad36600461089b565b61057e565b3480156101be57600080fd5b506100e16101cd3660046108fb565b6105d0565b6001546000906101e457506000610228565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061020f57fe5b6000918252602090912001546001600160a01b03161490505b919050565b600083815260046020908152604090912084815560018101805460ff19168515151790558251909161026691600284019185019061066b565b507f52f0a41c7e47098e4bb0415efbc816874814570bc6da2ef6ea32282bd9790a1b84848460405161029a93929190610c34565b60405180910390a150505050565b6102b06106d0565b60015482106102da5760405162461bcd60e51b81526004016102d190610c05565b60405180910390fd5b6102e26106d0565b600080600185815481106102f257fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156103ba5780601f1061038f576101008083540402835291602001916103ba565b820191906000526020600020905b81548152906001019060200180831161039d57829003601f168201915b5050509183525050600291909101546020909101529392505050565b6103de6106d0565b6103e7826101d2565b6104035760405162461bcd60e51b81526004016102d190610c05565b61040b6106d0565b6001600160a01b0383811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156103ba5780601f1061038f576101008083540402835291602001916103ba565b6104a1336101d2565b156104be5760405162461bcd60e51b81526004016102d190610bf5565b33600081815260208190526040902080546001600160a01b03191690911781556104ec6001820184846106fa565b506001805460028301819055915482820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690910180546001600160a01b0319166001600160a01b039092169190911790555050565b6001545b90565b60608282604051602001610566929190610ba9565b60405160208183030381529060405290505b92915050565b60006105886106d0565b610591836103d6565b60015490915060009060060243816105a557fe5b069050600682604001510281101580156105c85750600682604001516001010281105b949350505050565b60035460008181526002602052604090209081556105f26001820185856106fa565b50600281018290556003805481830181905582546001820183556000929092527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015580546040517fd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d289161029a91879087908790610c6a565b8280548282559060005260206000209081019282156106c0579160200282015b828111156106c057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061068b565b506106cc929150610774565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061073b5782800160ff19823516178555610768565b82800160010185558215610768579182015b8281111561076857823582559160200191906001019061074d565b506106cc929150610798565b61054e91905b808211156106cc5780546001600160a01b031916815560010161077a565b61054e91905b808211156106cc576000815560010161079e565b803561057881610d6e565b600082601f8301126107ce57600080fd5b81356107e16107dc82610cc1565b610c9a565b9150818183526020840193506020810190508385602084028201111561080657600080fd5b60005b83811015610832578161081c88826107b2565b8452506020928301929190910190600101610809565b5050505092915050565b803561057881610d85565b60008083601f84011261085957600080fd5b50813567ffffffffffffffff81111561087157600080fd5b60208301915083600182028301111561088957600080fd5b9250929050565b803561057881610d8e565b6000602082840312156108ad57600080fd5b60006105c884846107b2565b600080602083850312156108cc57600080fd5b823567ffffffffffffffff8111156108e357600080fd5b6108ef85828601610847565b92509250509250929050565b60008060006040848603121561091057600080fd5b833567ffffffffffffffff81111561092757600080fd5b61093386828701610847565b9350935050602061094686828701610890565b9150509250925092565b60006020828403121561096257600080fd5b60006105c88484610890565b6000806040838503121561098157600080fd5b600061098d8585610890565b925050602061099e8582860161083c565b9150509250929050565b6000806000606084860312156109bd57600080fd5b60006109c98686610890565b93505060206109da8682870161083c565b925050604084013567ffffffffffffffff8111156109f757600080fd5b610946868287016107bd565b6000610a0f8383610a17565b505060200190565b610a2081610cf5565b82525050565b6000610a3182610ce8565b610a3b8185610cec565b9350610a4683610ce2565b8060005b83811015610a74578151610a5e8882610a03565b9750610a6983610ce2565b925050600101610a4a565b509495945050505050565b610a2081610d00565b610a20610a9482610d00565b610d4d565b6000610aa482610ce8565b610aae8185610cec565b9350610abe818560208601610d1d565b610ac781610d5e565b9093019392505050565b6000610add8385610cec565b9350610aea838584610d11565b610ac783610d5e565b6000610b00601283610cec565b71185b1c9958591e481c9959da5cdd195c995960721b815260200192915050565b6000610b2e600983610cec565b681b9bdd08199bdd5b9960ba1b815260200192915050565b80516000906060840190610b5a8582610a17565b5060208301518482036020860152610b728282610a99565b9150506040830151610b876040860182610b8f565b509392505050565b610a208161054e565b610a20610ba48261054e565b61054e565b6000610bb58285610b98565b602082019150610bc58284610a88565b5060010192915050565b602081016105788284610a7f565b60208082528101610bee8184610a99565b9392505050565b6020808252810161057881610af3565b6020808252810161057881610b21565b60208082528101610bee8184610b46565b602081016105788284610b8f565b60608101610c428286610b8f565b610c4f6020830185610a7f565b8181036040830152610c618184610a26565b95945050505050565b60608101610c788287610b8f565b8181036020830152610c8b818587610ad1565b9050610c616040830184610b8f565b60405181810167ffffffffffffffff81118282101715610cb957600080fd5b604052919050565b600067ffffffffffffffff821115610cd857600080fd5b5060209081020190565b60200190565b5190565b90815260200190565b600061057882610d05565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610d38578181015183820152602001610d20565b83811115610d47576000848401525b50505050565b600061057882600061057882610d68565b601f01601f191690565b60f81b90565b610d7781610cf5565b8114610d8257600080fd5b50565b610d7781610d00565b610d778161054e56fea365627a7a72315820dc8fb3e5f11e3420e94ef7884bc430d1ca53e1c59da6cd9f26ace998da5fbc8b6c6578706572696d656e74616cf564736f6c63430005100040"

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

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractCaller) CountOracleNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "countOracleNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractSession) CountOracleNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountOracleNodes(&_OracleContract.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CountOracleNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountOracleNodes(&_OracleContract.CallOpts)
}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) view returns(bytes)
func (_OracleContract *OracleContractCaller) EncodeResult(opts *bind.CallOpts, _id *big.Int, _result bool) ([]byte, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "encodeResult", _id, _result)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) view returns(bytes)
func (_OracleContract *OracleContractSession) EncodeResult(_id *big.Int, _result bool) ([]byte, error) {
	return _OracleContract.Contract.EncodeResult(&_OracleContract.CallOpts, _id, _result)
}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) view returns(bytes)
func (_OracleContract *OracleContractCallerSession) EncodeResult(_id *big.Int, _result bool) ([]byte, error) {
	return _OracleContract.Contract.EncodeResult(&_OracleContract.CallOpts, _id, _result)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractCaller) FindOracleNodeByAddress(opts *bind.CallOpts, _addr common.Address) (OracleContractOracleNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findOracleNodeByAddress", _addr)

	if err != nil {
		return *new(OracleContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractOracleNode)).(*OracleContractOracleNode)

	return out0, err

}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractSession) FindOracleNodeByAddress(_addr common.Address) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256))
func (_OracleContract *OracleContractCallerSession) FindOracleNodeByAddress(_addr common.Address) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractCaller) FindOracleNodeByIndex(opts *bind.CallOpts, _index *big.Int) (OracleContractOracleNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findOracleNodeByIndex", _index)

	if err != nil {
		return *new(OracleContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractOracleNode)).(*OracleContractOracleNode)

	return out0, err

}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractSession) FindOracleNodeByIndex(_index *big.Int) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256))
func (_OracleContract *OracleContractCallerSession) FindOracleNodeByIndex(_index *big.Int) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByIndex(&_OracleContract.CallOpts, _index)
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

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) OracleNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "oracleNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x67cf3e48.
//
// Solidity: function registerOracleNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractTransactor) RegisterOracleNode(opts *bind.TransactOpts, _ipAddr string) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "registerOracleNode", _ipAddr)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x67cf3e48.
//
// Solidity: function registerOracleNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractSession) RegisterOracleNode(_ipAddr string) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterOracleNode(&_OracleContract.TransactOpts, _ipAddr)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x67cf3e48.
//
// Solidity: function registerOracleNode(string _ipAddr) payable returns()
func (_OracleContract *OracleContractTransactorSession) RegisterOracleNode(_ipAddr string) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterOracleNode(&_OracleContract.TransactOpts, _ipAddr)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0x18176c05.
//
// Solidity: function submitVerification(uint256 _id, bool _result, address[] _witnesses) returns()
func (_OracleContract *OracleContractTransactor) SubmitVerification(opts *bind.TransactOpts, _id *big.Int, _result bool, _witnesses []common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitVerification", _id, _result, _witnesses)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0x18176c05.
//
// Solidity: function submitVerification(uint256 _id, bool _result, address[] _witnesses) returns()
func (_OracleContract *OracleContractSession) SubmitVerification(_id *big.Int, _result bool, _witnesses []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitVerification(&_OracleContract.TransactOpts, _id, _result, _witnesses)
}

// SubmitVerification is a paid mutator transaction binding the contract method 0x18176c05.
//
// Solidity: function submitVerification(uint256 _id, bool _result, address[] _witnesses) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitVerification(_id *big.Int, _result bool, _witnesses []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitVerification(&_OracleContract.TransactOpts, _id, _result, _witnesses)
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
	Id        *big.Int
	Result    bool
	Witnesses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitVerificationLog is a free log retrieval operation binding the contract event 0x52f0a41c7e47098e4bb0415efbc816874814570bc6da2ef6ea32282bd9790a1b.
//
// Solidity: event SubmitVerificationLog(uint256 id, bool result, address[] witnesses)
func (_OracleContract *OracleContractFilterer) FilterSubmitVerificationLog(opts *bind.FilterOpts) (*OracleContractSubmitVerificationLogIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "SubmitVerificationLog")
	if err != nil {
		return nil, err
	}
	return &OracleContractSubmitVerificationLogIterator{contract: _OracleContract.contract, event: "SubmitVerificationLog", logs: logs, sub: sub}, nil
}

// WatchSubmitVerificationLog is a free log subscription operation binding the contract event 0x52f0a41c7e47098e4bb0415efbc816874814570bc6da2ef6ea32282bd9790a1b.
//
// Solidity: event SubmitVerificationLog(uint256 id, bool result, address[] witnesses)
func (_OracleContract *OracleContractFilterer) WatchSubmitVerificationLog(opts *bind.WatchOpts, sink chan<- *OracleContractSubmitVerificationLog) (event.Subscription, error) {

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "SubmitVerificationLog")
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

// ParseSubmitVerificationLog is a log parse operation binding the contract event 0x52f0a41c7e47098e4bb0415efbc816874814570bc6da2ef6ea32282bd9790a1b.
//
// Solidity: event SubmitVerificationLog(uint256 id, bool result, address[] witnesses)
func (_OracleContract *OracleContractFilterer) ParseSubmitVerificationLog(log types.Log) (*OracleContractSubmitVerificationLog, error) {
	event := new(OracleContractSubmitVerificationLog)
	if err := _OracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
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

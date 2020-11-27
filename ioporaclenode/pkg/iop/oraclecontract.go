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

// OracleContractVerificationResult is an auto generated low-level Go binding around an user-defined struct.
type OracleContractVerificationResult struct {
	Id               *big.Int
	Request          *big.Int
	Result           bool
	Witnesses        []common.Address
	Finalized        *big.Int
	RewardTransfered bool
}

// OracleContractABI is the input ABI used to generate the binding from.
const OracleContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalized\",\"type\":\"uint256\"}],\"name\":\"SubmitVerificationLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"encodeResult\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findOracleNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"findVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"finalized\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rewardTransfered\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.VerificationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerOracleNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_witnesses\",\"type\":\"address[]\"}],\"name\":\"submitVerification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"transferReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"verificationExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OracleContractFuncSigs maps the 4-byte function signature to its string representation.
var OracleContractFuncSigs = map[string]string{
	"836f187a": "countOracleNodes()",
	"b33d96bd": "encodeResult(uint256,bool)",
	"655a102f": "findOracleNodeByAddress(address)",
	"272132e9": "findOracleNodeByIndex(uint256)",
	"0bfb3232": "findVerification(uint256)",
	"db512e85": "isLeader(address)",
	"140f3daa": "oracleNodeIsRegistered(address)",
	"67cf3e48": "registerOracleNode(string)",
	"18176c05": "submitVerification(uint256,bool,address[])",
	"b53e07c4": "transferReward(uint256)",
	"c6704288": "verificationExists(uint256)",
	"ff84d3e9": "verifyTransaction(string,uint256)",
}

// OracleContractBin is the compiled bytecode used for deploying new contracts.
var OracleContractBin = "0x608060405234801561001057600080fd5b506113d2806100206000396000f3fe6080604052600436106100a75760003560e01c8063836f187a11610064578063836f187a14610191578063b33d96bd146101b3578063b53e07c4146101e0578063c670428814610200578063db512e8514610220578063ff84d3e914610240576100a7565b80630bfb3232146100ac578063140f3daa146100e257806318176c051461010f578063272132e914610131578063655a102f1461015e57806367cf3e481461017e575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004610d8a565b610253565b6040516100d991906111e7565b60405180910390f35b3480156100ee57600080fd5b506101026100fd366004610cd5565b61034b565b6040516100d99190611160565b34801561011b57600080fd5b5061012f61012a366004610de2565b6103a4565b005b34801561013d57600080fd5b5061015161014c366004610d8a565b610440565b6040516100d991906111d6565b34801561016a57600080fd5b50610151610179366004610cd5565b610565565b61012f61018c366004610cf3565b610627565b34801561019d57600080fd5b506101a66106d9565b6040516100d991906111f8565b3480156101bf57600080fd5b506101d36101ce366004610da8565b6106e0565b6040516100d9919061116e565b3480156101ec57600080fd5b5061012f6101fb366004610d8a565b61070d565b34801561020c57600080fd5b5061010261021b366004610d8a565b610935565b34801561022c57600080fd5b5061010261023b366004610cd5565b610949565b61012f61024e366004610d35565b61099b565b61025b610a6b565b61026482610935565b6102895760405162461bcd60e51b8152600401610280906111c6565b60405180910390fd5b610291610a6b565b600083815260056020908152604091829020825160c08101845281548152600182015481840152600282015460ff16151581850152600382018054855181860281018601909652808652919492936060860193929083018282801561031f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610301575b50505091835250506004820154602082015260059091015460ff1615156040909101529150505b919050565b60015460009061035d57506000610346565b6001600160a01b03821660008181526020819052604090206002015460018054909190811061038857fe5b6000918252602090912001546001600160a01b03161492915050565b60068054600190810191829055600082815260056020908152604090912083815591820186905560028201805460ff191686151517905583516103ef91600384019190860190610aa5565b5043600601600482018190556040517f1bb1b6460e047ed647182ddcdeda6e34c3020b43ffe0f0a5719915ef260e42da9161043191859189918991899161123f565b60405180910390a15050505050565b610448610b0a565b60015482106104695760405162461bcd60e51b8152600401610280906111c6565b610471610b0a565b6000806001858154811061048157fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156105495780601f1061051e57610100808354040283529160200191610549565b820191906000526020600020905b81548152906001019060200180831161052c57829003601f168201915b5050509183525050600291909101546020909101529392505050565b61056d610b0a565b6105768261034b565b6105925760405162461bcd60e51b8152600401610280906111c6565b61059a610b0a565b6001600160a01b0383811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156105495780601f1061051e57610100808354040283529160200191610549565b6106303361034b565b1561064d5760405162461bcd60e51b815260040161028090611196565b33600081815260208190526040902080546001600160a01b031916909117815561067b600182018484610b34565b506001805460028301819055915482820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690910180546001600160a01b0319166001600160a01b039092169190911790555050565b6001545b90565b606082826040516020016106f592919061113a565b60405160208183030381529060405290505b92915050565b610715610a6b565b61071e82610253565b905043816080015111156107445760405162461bcd60e51b815260040161028090611186565b8060a00151156107665760405162461bcd60e51b8152600401610280906111a6565b600160a0820181905281516000908152600560209081526040918290208451815581850151938101939093559083015160028301805460ff1916911515919091179055606083015180518493926107c4926003850192910190610aa5565b506080820151600482015560a0909101516005909101805460ff191691151591909117905560005b6001548111610930576107fd610b0a565b6000806001848154811061080d57fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156108d55780601f106108aa576101008083540402835291602001916108d5565b820191906000526020600020905b8154815290600101906020018083116108b857829003601f168201915b50505091835250506002919091015460209091015280516040519192506001600160a01b03169060009066038d7ea4c680009082818181858883f19350505050158015610926573d6000803e3d6000fd5b50506001016107ec565b505050565b600090815260056020526040902054151590565b6000610953610b0a565b61095c83610565565b600154909150600090600602438161097057fe5b069050600682604001510281101580156109935750600682604001516001010281105b949350505050565b60015466038d7ea4c68000023410156109c65760405162461bcd60e51b8152600401610280906111b6565b600480546001908101918290556000828152600260205260409020828155906109f29082018686610b34565b5060028101839055600380548183018190556001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018290556040517fd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d2890610431908490889088908890611206565b6040518060c00160405280600081526020016000815260200160001515815260200160608152602001600081526020016000151581525090565b828054828255906000526020600020908101928215610afa579160200282015b82811115610afa57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610ac5565b50610b06929150610bae565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b755782800160ff19823516178555610ba2565b82800160010185558215610ba2579182015b82811115610ba2578235825591602001919060010190610b87565b50610b06929150610bd2565b6106dd91905b80821115610b065780546001600160a01b0319168155600101610bb4565b6106dd91905b80821115610b065760008155600101610bd8565b803561070781611366565b600082601f830112610c0857600080fd5b8135610c1b610c16826112b9565b611292565b91508181835260208401935060208101905083856020840282011115610c4057600080fd5b60005b83811015610c6c5781610c568882610bec565b8452506020928301929190910190600101610c43565b5050505092915050565b80356107078161137d565b60008083601f840112610c9357600080fd5b50813567ffffffffffffffff811115610cab57600080fd5b602083019150836001820283011115610cc357600080fd5b9250929050565b803561070781611386565b600060208284031215610ce757600080fd5b60006109938484610bec565b60008060208385031215610d0657600080fd5b823567ffffffffffffffff811115610d1d57600080fd5b610d2985828601610c81565b92509250509250929050565b600080600060408486031215610d4a57600080fd5b833567ffffffffffffffff811115610d6157600080fd5b610d6d86828701610c81565b93509350506020610d8086828701610cca565b9150509250925092565b600060208284031215610d9c57600080fd5b60006109938484610cca565b60008060408385031215610dbb57600080fd5b6000610dc78585610cca565b9250506020610dd885828601610c76565b9150509250929050565b600080600060608486031215610df757600080fd5b6000610e038686610cca565b9350506020610e1486828701610c76565b925050604084013567ffffffffffffffff811115610e3157600080fd5b610d8086828701610bf7565b6000610e498383610e51565b505060200190565b610e5a816112ed565b82525050565b6000610e6b826112e0565b610e7581856112e4565b9350610e80836112da565b8060005b83811015610eae578151610e988882610e3d565b9750610ea3836112da565b925050600101610e84565b509495945050505050565b6000610ec4826112e0565b610ece81856112e4565b9350610ed9836112da565b8060005b83811015610eae578151610ef18882610e3d565b9750610efc836112da565b925050600101610edd565b610e5a816112f8565b610e5a610f1c826112f8565b611345565b6000610f2c826112e0565b610f3681856112e4565b9350610f46818560208601611315565b610f4f81611356565b9093019392505050565b6000610f6583856112e4565b9350610f72838584611309565b610f4f83611356565b6000610f88600d836112e4565b6c1b9bdd08199a5b985b1a5e9959609a1b815260200192915050565b6000610fb16012836112e4565b71185b1c9958591e481c9959da5cdd195c995960721b815260200192915050565b6000610fdf6019836112e4565b7f72657761726420616c7265616479207472616e73666572656400000000000000815260200192915050565b60006110186011836112e4565b706d73672076616c756520746f6f206c6f7760781b815260200192915050565b60006110456009836112e4565b681b9bdd08199bdd5b9960ba1b815260200192915050565b805160009060608401906110718582610e51565b50602083015184820360208601526110898282610f21565b915050604083015161109e6040860182611120565b509392505050565b805160009060c08401906110ba8582611120565b5060208301516110cd6020860182611120565b5060408301516110e06040860182610f07565b50606083015184820360608601526110f88282610eb9565b915050608083015161110d6080860182611120565b5060a083015161109e60a0860182610f07565b610e5a816106dd565b610e5a611135826106dd565b6106dd565b60006111468285611129565b6020820191506111568284610f10565b5060010192915050565b602081016107078284610f07565b6020808252810161117f8184610f21565b9392505050565b6020808252810161070781610f7b565b6020808252810161070781610fa4565b6020808252810161070781610fd2565b602080825281016107078161100b565b6020808252810161070781611038565b6020808252810161117f818461105d565b6020808252810161117f81846110a6565b602081016107078284611120565b606081016112148287611120565b8181036020830152611227818587610f59565b90506112366040830184611120565b95945050505050565b60a0810161124d8288611120565b61125a6020830187611120565b6112676040830186610f07565b81810360608301526112798185610e60565b90506112886080830184611120565b9695505050505050565b60405181810167ffffffffffffffff811182821017156112b157600080fd5b604052919050565b600067ffffffffffffffff8211156112d057600080fd5b5060209081020190565b60200190565b5190565b90815260200190565b6000610707826112fd565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015611330578181015183820152602001611318565b8381111561133f576000848401525b50505050565b600061070782600061070782611360565b601f01601f191690565b60f81b90565b61136f816112ed565b811461137a57600080fd5b50565b61136f816112f8565b61136f816106dd56fea365627a7a7231582072bc42ea57c361325a05c97d74614c6cf8b06e775db4c9fbdd97d9ae8251449b6c6578706572696d656e74616cf564736f6c63430005100040"

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
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
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
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
func (_OracleContract *OracleContractSession) EncodeResult(_id *big.Int, _result bool) ([]byte, error) {
	return _OracleContract.Contract.EncodeResult(&_OracleContract.CallOpts, _id, _result)
}

// EncodeResult is a free data retrieval call binding the contract method 0xb33d96bd.
//
// Solidity: function encodeResult(uint256 _id, bool _result) pure returns(bytes)
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

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool,address[],uint256,bool))
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
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool,address[],uint256,bool))
func (_OracleContract *OracleContractSession) FindVerification(_id *big.Int) (OracleContractVerificationResult, error) {
	return _OracleContract.Contract.FindVerification(&_OracleContract.CallOpts, _id)
}

// FindVerification is a free data retrieval call binding the contract method 0x0bfb3232.
//
// Solidity: function findVerification(uint256 _id) view returns((uint256,uint256,bool,address[],uint256,bool))
func (_OracleContract *OracleContractCallerSession) FindVerification(_id *big.Int) (OracleContractVerificationResult, error) {
	return _OracleContract.Contract.FindVerification(&_OracleContract.CallOpts, _id)
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

// TransferReward is a paid mutator transaction binding the contract method 0xb53e07c4.
//
// Solidity: function transferReward(uint256 _id) returns()
func (_OracleContract *OracleContractTransactor) TransferReward(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "transferReward", _id)
}

// TransferReward is a paid mutator transaction binding the contract method 0xb53e07c4.
//
// Solidity: function transferReward(uint256 _id) returns()
func (_OracleContract *OracleContractSession) TransferReward(_id *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.TransferReward(&_OracleContract.TransactOpts, _id)
}

// TransferReward is a paid mutator transaction binding the contract method 0xb53e07c4.
//
// Solidity: function transferReward(uint256 _id) returns()
func (_OracleContract *OracleContractTransactorSession) TransferReward(_id *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.TransferReward(&_OracleContract.TransactOpts, _id)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractTransactor) VerifyTransaction(opts *bind.TransactOpts, _hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "verifyTransaction", _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
func (_OracleContract *OracleContractSession) VerifyTransaction(_hash string, _confirmations *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.VerifyTransaction(&_OracleContract.TransactOpts, _hash, _confirmations)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0xff84d3e9.
//
// Solidity: function verifyTransaction(string _hash, uint256 _confirmations) payable returns()
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
	Request   *big.Int
	Result    bool
	Witnesses []common.Address
	Finalized *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitVerificationLog is a free log retrieval operation binding the contract event 0x1bb1b6460e047ed647182ddcdeda6e34c3020b43ffe0f0a5719915ef260e42da.
//
// Solidity: event SubmitVerificationLog(uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
func (_OracleContract *OracleContractFilterer) FilterSubmitVerificationLog(opts *bind.FilterOpts) (*OracleContractSubmitVerificationLogIterator, error) {

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "SubmitVerificationLog")
	if err != nil {
		return nil, err
	}
	return &OracleContractSubmitVerificationLogIterator{contract: _OracleContract.contract, event: "SubmitVerificationLog", logs: logs, sub: sub}, nil
}

// WatchSubmitVerificationLog is a free log subscription operation binding the contract event 0x1bb1b6460e047ed647182ddcdeda6e34c3020b43ffe0f0a5719915ef260e42da.
//
// Solidity: event SubmitVerificationLog(uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
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

// ParseSubmitVerificationLog is a log parse operation binding the contract event 0x1bb1b6460e047ed647182ddcdeda6e34c3020b43ffe0f0a5719915ef260e42da.
//
// Solidity: event SubmitVerificationLog(uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
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

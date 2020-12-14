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
const OracleContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RegisterOracleNodeLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalized\",\"type\":\"uint256\"}],\"name\":\"SubmitVerificationLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"}],\"name\":\"TransferRewardLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"VerifyTransactionLog\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"encodeResult\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findOracleNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"findVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"request\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"finalized\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"rewardTransfered\",\"type\":\"bool\"}],\"internalType\":\"structOracleContract.VerificationResult\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLeader\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"}],\"name\":\"registerOracleNode\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"_witnesses\",\"type\":\"address[]\"}],\"name\":\"submitVerification\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"transferReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"verificationExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_confirmations\",\"type\":\"uint256\"}],\"name\":\"verifyTransaction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OracleContractFuncSigs maps the 4-byte function signature to its string representation.
var OracleContractFuncSigs = map[string]string{
	"836f187a": "countOracleNodes()",
	"b33d96bd": "encodeResult(uint256,bool)",
	"655a102f": "findOracleNodeByAddress(address)",
	"272132e9": "findOracleNodeByIndex(uint256)",
	"0bfb3232": "findVerification(uint256)",
	"4c051f14": "getLeader()",
	"db512e85": "isLeader(address)",
	"140f3daa": "oracleNodeIsRegistered(address)",
	"67cf3e48": "registerOracleNode(string)",
	"18176c05": "submitVerification(uint256,bool,address[])",
	"b53e07c4": "transferReward(uint256)",
	"c6704288": "verificationExists(uint256)",
	"ff84d3e9": "verifyTransaction(string,uint256)",
}

// OracleContractBin is the compiled bytecode used for deploying new contracts.
var OracleContractBin = "0x608060405234801561001057600080fd5b5061156d806100206000396000f3fe6080604052600436106100c25760003560e01c806367cf3e481161007f578063b53e07c411610059578063b53e07c414610210578063c670428814610230578063db512e8514610250578063ff84d3e914610270576100c2565b806367cf3e48146101ae578063836f187a146101c1578063b33d96bd146101e3576100c2565b80630bfb3232146100c7578063140f3daa146100fd57806318176c051461012a578063272132e91461014c5780634c051f1414610179578063655a102f1461018e575b600080fd5b3480156100d357600080fd5b506100e76100e2366004610eb6565b610283565b6040516100f49190611382565b60405180910390f35b34801561010957600080fd5b5061011d610118366004610e01565b61037b565b6040516100f491906112db565b34801561013657600080fd5b5061014a610145366004610f0e565b6103d4565b005b34801561015857600080fd5b5061016c610167366004610eb6565b610498565b6040516100f49190611371565b34801561018557600080fd5b5061016c6105bd565b34801561019a57600080fd5b5061016c6101a9366004610e01565b610707565b61014a6101bc366004610e1f565b6107c9565b3480156101cd57600080fd5b506101d66108a3565b6040516100f49190611393565b3480156101ef57600080fd5b506102036101fe366004610ed4565b6108a9565b6040516100f491906112e9565b34801561021c57600080fd5b5061014a61022b366004610eb6565b6108d6565b34801561023c57600080fd5b5061011d61024b366004610eb6565b610a52565b34801561025c57600080fd5b5061011d61026b366004610e01565b610a66565b61014a61027e366004610e61565b610ab8565b61028b610b97565b61029482610a52565b6102b95760405162461bcd60e51b81526004016102b090611361565b60405180910390fd5b6102c1610b97565b600083815260056020908152604091829020825160c08101845281548152600182015481840152600282015460ff16151581850152600382018054855181860281018601909652808652919492936060860193929083018282801561034f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610331575b50505091835250506004820154602082015260059091015460ff1615156040909101529150505b919050565b60015460009061038d57506000610376565b6001600160a01b0382166000818152602081905260409020600201546001805490919081106103b857fe5b6000918252602090912001546001600160a01b03161492915050565b6103dd33610a66565b6103f95760405162461bcd60e51b81526004016102b090611331565b60068054600190810191829055600082815260056020908152604090912083815591820186905560028201805460ff1916861515179055835161044491600384019190860190610bd1565b50436006016004820181905560405133917f86551a930b3ee10d60fad9f4cfa9c1ad0ce871dc8318cdd7233523ad7b4b426e916104899186918a918a918a91906113da565b60405180910390a25050505050565b6104a0610c36565b60015482106104c15760405162461bcd60e51b81526004016102b090611361565b6104c9610c36565b600080600185815481106104d957fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156105a15780601f10610576576101008083540402835291602001916105a1565b820191906000526020600020905b81548152906001019060200180831161058457829003601f168201915b5050509183525050600291909101546020909101529392505050565b6105c5610c36565b60005b6001548110156106eb576105da610c36565b600080600184815481106105ea57fe5b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160608101845281549094168452600181810180548551601f6002600019958416156101000295909501909216939093049081018590048502830185019095528482529193858401939192918301828280156106b25780601f10610687576101008083540402835291602001916106b2565b820191906000526020600020905b81548152906001019060200180831161069557829003601f168201915b5050505050815260200160028201548152505090506106d48160000151610a66565b156106e25791506107049050565b506001016105c8565b5060405162461bcd60e51b81526004016102b090611351565b90565b61070f610c36565b6107188261037b565b6107345760405162461bcd60e51b81526004016102b090611361565b61073c610c36565b6001600160a01b0383811660009081526020818152604091829020825160608101845281549094168452600180820180548551600261010094831615949094026000190190911692909204601f81018590048502830185019095528482529193858401939192918301828280156105a15780601f10610576576101008083540402835291602001916105a1565b6107d23361037b565b156107ef5760405162461bcd60e51b81526004016102b090611311565b33600081815260208190526040902080546001600160a01b031916909117815561081d600182018484610c60565b5060018054600283018190558254818301835560009283527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf690910180546001600160a01b0319166001600160a01b0390921691909117905560405133917f1e8e81a8a2dad3aa858f8e7f00c25da73b2a1692fc9f7d3dfd7f96412100c30991a2505050565b60015490565b606082826040516020016108be9291906112b5565b60405160208183030381529060405290505b92915050565b6108de610b97565b6108e782610283565b9050438160800151111561090d5760405162461bcd60e51b81526004016102b090611301565b8060a001511561092f5760405162461bcd60e51b81526004016102b090611321565b600160a0820181905281516000908152600560209081526040918290208451815581850151938101939093559083015160028301805460ff19169115159190911790556060830151805184939261098d926003850192910190610bd1565b506080820151600482015560a0909101516005909101805460ff191691151591909117905560005b816060015151811015610a2257816060015181815181106109d257fe5b60200260200101516001600160a01b03166108fc66038d7ea4c680009081150290604051600060405180830381858888f19350505050158015610a19573d6000803e3d6000fd5b506001016109b5565b5060405182907f9ea512e9a8b840dbe7dee54e29eefbfbafbc78dfb1af17bb3ab25b828fca573590600090a25050565b600090815260056020526040902054151590565b6000610a70610c36565b610a7983610707565b6001549091506000906006024381610a8d57fe5b06905060068260400151028110158015610ab05750600682604001516001010281105b949350505050565b60015466038d7ea4c6800002341015610ae35760405162461bcd60e51b81526004016102b090611341565b60048054600190810191829055600082815260026020526040902082815590610b0f9082018686610c60565b5060028101839055600380548183018190556001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018290556040517fd2b9ae20b01dc42907609e33ebc81783e68018abdd61e1c616f3ce800d470d2890610b889084908890889088906113a1565b60405180910390a15050505050565b6040518060c00160405280600081526020016000815260200160001515815260200160608152602001600081526020016000151581525090565b828054828255906000526020600020908101928215610c26579160200282015b82811115610c2657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610bf1565b50610c32929150610cda565b5090565b604051806060016040528060006001600160a01b0316815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ca15782800160ff19823516178555610cce565b82800160010185558215610cce579182015b82811115610cce578235825591602001919060010190610cb3565b50610c32929150610cfe565b61070491905b80821115610c325780546001600160a01b0319168155600101610ce0565b61070491905b80821115610c325760008155600101610d04565b80356108d081611501565b600082601f830112610d3457600080fd5b8135610d47610d4282611454565b61142d565b91508181835260208401935060208101905083856020840282011115610d6c57600080fd5b60005b83811015610d985781610d828882610d18565b8452506020928301929190910190600101610d6f565b5050505092915050565b80356108d081611518565b60008083601f840112610dbf57600080fd5b50813567ffffffffffffffff811115610dd757600080fd5b602083019150836001820283011115610def57600080fd5b9250929050565b80356108d081611521565b600060208284031215610e1357600080fd5b6000610ab08484610d18565b60008060208385031215610e3257600080fd5b823567ffffffffffffffff811115610e4957600080fd5b610e5585828601610dad565b92509250509250929050565b600080600060408486031215610e7657600080fd5b833567ffffffffffffffff811115610e8d57600080fd5b610e9986828701610dad565b93509350506020610eac86828701610df6565b9150509250925092565b600060208284031215610ec857600080fd5b6000610ab08484610df6565b60008060408385031215610ee757600080fd5b6000610ef38585610df6565b9250506020610f0485828601610da2565b9150509250929050565b600080600060608486031215610f2357600080fd5b6000610f2f8686610df6565b9350506020610f4086828701610da2565b925050604084013567ffffffffffffffff811115610f5d57600080fd5b610eac86828701610d23565b6000610f758383610f7d565b505060200190565b610f8681611488565b82525050565b6000610f978261147b565b610fa1818561147f565b9350610fac83611475565b8060005b83811015610fda578151610fc48882610f69565b9750610fcf83611475565b925050600101610fb0565b509495945050505050565b6000610ff08261147b565b610ffa818561147f565b935061100583611475565b8060005b83811015610fda57815161101d8882610f69565b975061102883611475565b925050600101611009565b610f8681611493565b610f8661104882611493565b6114e0565b60006110588261147b565b611062818561147f565b93506110728185602086016114b0565b61107b816114f1565b9093019392505050565b6000611091838561147f565b935061109e8385846114a4565b61107b836114f1565b60006110b4600d8361147f565b6c1b9bdd08199a5b985b1a5e9959609a1b815260200192915050565b60006110dd60128361147f565b71185b1c9958591e481c9959da5cdd195c995960721b815260200192915050565b600061110b60198361147f565b7f72657761726420616c7265616479207472616e73666572656400000000000000815260200192915050565b6000611144600e8361147f565b6d3737ba103a3432903632b0b232b960911b815260200192915050565b600061116e60118361147f565b706d73672076616c756520746f6f206c6f7760781b815260200192915050565b600061119b60098361147f565b683737903632b0b232b960b91b815260200192915050565b60006111c060098361147f565b681b9bdd08199bdd5b9960ba1b815260200192915050565b805160009060608401906111ec8582610f7d565b5060208301518482036020860152611204828261104d565b9150506040830151611219604086018261129b565b509392505050565b805160009060c0840190611235858261129b565b506020830151611248602086018261129b565b50604083015161125b6040860182611033565b50606083015184820360608601526112738282610fe5565b9150506080830151611288608086018261129b565b5060a083015161121960a0860182611033565b610f8681610704565b610f866112b082610704565b610704565b60006112c182856112a4565b6020820191506112d1828461103c565b5060010192915050565b602081016108d08284611033565b602080825281016112fa818461104d565b9392505050565b602080825281016108d0816110a7565b602080825281016108d0816110d0565b602080825281016108d0816110fe565b602080825281016108d081611137565b602080825281016108d081611161565b602080825281016108d08161118e565b602080825281016108d0816111b3565b602080825281016112fa81846111d8565b602080825281016112fa8184611221565b602081016108d0828461129b565b606081016113af828761129b565b81810360208301526113c2818587611085565b90506113d1604083018461129b565b95945050505050565b60a081016113e8828861129b565b6113f5602083018761129b565b6114026040830186611033565b81810360608301526114148185610f8c565b9050611423608083018461129b565b9695505050505050565b60405181810167ffffffffffffffff8111828210171561144c57600080fd5b604052919050565b600067ffffffffffffffff82111561146b57600080fd5b5060209081020190565b60200190565b5190565b90815260200190565b60006108d082611498565b151590565b6001600160a01b031690565b82818337506000910152565b60005b838110156114cb5781810151838201526020016114b3565b838111156114da576000848401525b50505050565b60006108d08260006108d0826114fb565b601f01601f191690565b60f81b90565b61150a81611488565b811461151557600080fd5b50565b61150a81611493565b61150a8161070456fea365627a7a72315820a9dc834e93c0ec14a2510567ff3a148ad08d333c494ad1daf29dd7cd8a9830356c6578706572696d656e74616cf564736f6c63430005100040"

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

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns((address,string,uint256))
func (_OracleContract *OracleContractCaller) GetLeader(opts *bind.CallOpts) (OracleContractOracleNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getLeader")

	if err != nil {
		return *new(OracleContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractOracleNode)).(*OracleContractOracleNode)

	return out0, err

}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns((address,string,uint256))
func (_OracleContract *OracleContractSession) GetLeader() (OracleContractOracleNode, error) {
	return _OracleContract.Contract.GetLeader(&_OracleContract.CallOpts)
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns((address,string,uint256))
func (_OracleContract *OracleContractCallerSession) GetLeader() (OracleContractOracleNode, error) {
	return _OracleContract.Contract.GetLeader(&_OracleContract.CallOpts)
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

// OracleContractRegisterOracleNodeLogIterator is returned from FilterRegisterOracleNodeLog and is used to iterate over the raw logs and unpacked data for RegisterOracleNodeLog events raised by the OracleContract contract.
type OracleContractRegisterOracleNodeLogIterator struct {
	Event *OracleContractRegisterOracleNodeLog // Event containing the contract specifics and raw log

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
func (it *OracleContractRegisterOracleNodeLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractRegisterOracleNodeLog)
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
		it.Event = new(OracleContractRegisterOracleNodeLog)
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
func (it *OracleContractRegisterOracleNodeLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractRegisterOracleNodeLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractRegisterOracleNodeLog represents a RegisterOracleNodeLog event raised by the OracleContract contract.
type OracleContractRegisterOracleNodeLog struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterOracleNodeLog is a free log retrieval operation binding the contract event 0x1e8e81a8a2dad3aa858f8e7f00c25da73b2a1692fc9f7d3dfd7f96412100c309.
//
// Solidity: event RegisterOracleNodeLog(address indexed sender)
func (_OracleContract *OracleContractFilterer) FilterRegisterOracleNodeLog(opts *bind.FilterOpts, sender []common.Address) (*OracleContractRegisterOracleNodeLogIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "RegisterOracleNodeLog", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractRegisterOracleNodeLogIterator{contract: _OracleContract.contract, event: "RegisterOracleNodeLog", logs: logs, sub: sub}, nil
}

// WatchRegisterOracleNodeLog is a free log subscription operation binding the contract event 0x1e8e81a8a2dad3aa858f8e7f00c25da73b2a1692fc9f7d3dfd7f96412100c309.
//
// Solidity: event RegisterOracleNodeLog(address indexed sender)
func (_OracleContract *OracleContractFilterer) WatchRegisterOracleNodeLog(opts *bind.WatchOpts, sink chan<- *OracleContractRegisterOracleNodeLog, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "RegisterOracleNodeLog", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractRegisterOracleNodeLog)
				if err := _OracleContract.contract.UnpackLog(event, "RegisterOracleNodeLog", log); err != nil {
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

// ParseRegisterOracleNodeLog is a log parse operation binding the contract event 0x1e8e81a8a2dad3aa858f8e7f00c25da73b2a1692fc9f7d3dfd7f96412100c309.
//
// Solidity: event RegisterOracleNodeLog(address indexed sender)
func (_OracleContract *OracleContractFilterer) ParseRegisterOracleNodeLog(log types.Log) (*OracleContractRegisterOracleNodeLog, error) {
	event := new(OracleContractRegisterOracleNodeLog)
	if err := _OracleContract.contract.UnpackLog(event, "RegisterOracleNodeLog", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Sender    common.Address
	Id        *big.Int
	Request   *big.Int
	Result    bool
	Witnesses []common.Address
	Finalized *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitVerificationLog is a free log retrieval operation binding the contract event 0x86551a930b3ee10d60fad9f4cfa9c1ad0ce871dc8318cdd7233523ad7b4b426e.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
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

// WatchSubmitVerificationLog is a free log subscription operation binding the contract event 0x86551a930b3ee10d60fad9f4cfa9c1ad0ce871dc8318cdd7233523ad7b4b426e.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
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

// ParseSubmitVerificationLog is a log parse operation binding the contract event 0x86551a930b3ee10d60fad9f4cfa9c1ad0ce871dc8318cdd7233523ad7b4b426e.
//
// Solidity: event SubmitVerificationLog(address indexed sender, uint256 id, uint256 request, bool result, address[] witnesses, uint256 finalized)
func (_OracleContract *OracleContractFilterer) ParseSubmitVerificationLog(log types.Log) (*OracleContractSubmitVerificationLog, error) {
	event := new(OracleContractSubmitVerificationLog)
	if err := _OracleContract.contract.UnpackLog(event, "SubmitVerificationLog", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OracleContractTransferRewardLogIterator is returned from FilterTransferRewardLog and is used to iterate over the raw logs and unpacked data for TransferRewardLog events raised by the OracleContract contract.
type OracleContractTransferRewardLogIterator struct {
	Event *OracleContractTransferRewardLog // Event containing the contract specifics and raw log

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
func (it *OracleContractTransferRewardLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractTransferRewardLog)
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
		it.Event = new(OracleContractTransferRewardLog)
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
func (it *OracleContractTransferRewardLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractTransferRewardLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractTransferRewardLog represents a TransferRewardLog event raised by the OracleContract contract.
type OracleContractTransferRewardLog struct {
	Request *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferRewardLog is a free log retrieval operation binding the contract event 0x9ea512e9a8b840dbe7dee54e29eefbfbafbc78dfb1af17bb3ab25b828fca5735.
//
// Solidity: event TransferRewardLog(uint256 indexed request)
func (_OracleContract *OracleContractFilterer) FilterTransferRewardLog(opts *bind.FilterOpts, request []*big.Int) (*OracleContractTransferRewardLogIterator, error) {

	var requestRule []interface{}
	for _, requestItem := range request {
		requestRule = append(requestRule, requestItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "TransferRewardLog", requestRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransferRewardLogIterator{contract: _OracleContract.contract, event: "TransferRewardLog", logs: logs, sub: sub}, nil
}

// WatchTransferRewardLog is a free log subscription operation binding the contract event 0x9ea512e9a8b840dbe7dee54e29eefbfbafbc78dfb1af17bb3ab25b828fca5735.
//
// Solidity: event TransferRewardLog(uint256 indexed request)
func (_OracleContract *OracleContractFilterer) WatchTransferRewardLog(opts *bind.WatchOpts, sink chan<- *OracleContractTransferRewardLog, request []*big.Int) (event.Subscription, error) {

	var requestRule []interface{}
	for _, requestItem := range request {
		requestRule = append(requestRule, requestItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "TransferRewardLog", requestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractTransferRewardLog)
				if err := _OracleContract.contract.UnpackLog(event, "TransferRewardLog", log); err != nil {
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

// ParseTransferRewardLog is a log parse operation binding the contract event 0x9ea512e9a8b840dbe7dee54e29eefbfbafbc78dfb1af17bb3ab25b828fca5735.
//
// Solidity: event TransferRewardLog(uint256 indexed request)
func (_OracleContract *OracleContractFilterer) ParseTransferRewardLog(log types.Log) (*OracleContractTransferRewardLog, error) {
	event := new(OracleContractTransferRewardLog)
	if err := _OracleContract.contract.UnpackLog(event, "TransferRewardLog", log); err != nil {
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

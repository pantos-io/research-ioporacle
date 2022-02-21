package iop

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func encodeValidateResult(hash common.Hash, result bool) ([]byte, error) {
	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: boolTy,
		},
	}
	return arguments.Pack(hash, result)
}

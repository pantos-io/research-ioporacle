package iop

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func encodeValidateResult(hash common.Hash, result bool, typ ValidateRequest_Type) ([]byte, error) {
	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)
	uint8Ty, _ := abi.NewType("uint8", "", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: boolTy,
		},
		{
			Type: uint8Ty,
		},
	}
	return arguments.Pack(hash, result, uint8(typ))
}

package iop

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

func encodeValidateTransactionResult(id *big.Int, result bool) ([]byte, error) {
	uint256Ty, _ := abi.NewType("uint256", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: boolTy,
		},
	}
	return arguments.Pack(id, result)
}

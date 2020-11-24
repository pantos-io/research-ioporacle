package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TransactionVerifier interface {
	VerifyTransaction(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error)
}

type transactionVerifierImpl struct {
	ethClient      *ethclient.Client
	oracleContract *OracleContract
}

func NewTransactionVerifier(ethClient *ethclient.Client, oracleContract *OracleContract) *transactionVerifierImpl {
	return &transactionVerifierImpl{
		ethClient:      ethClient,
		oracleContract: oracleContract,
	}
}

func (t transactionVerifierImpl) VerifyTransaction(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error) {
	receipt, err := t.ethClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		if err.Error() == "not found" {
			return false, nil
		}
		return false, fmt.Errorf("transaction receipt: %w", err)
	}
	blockNumber, err := t.ethClient.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("blocknumber: %w", err)
	}
	confirmed := blockNumber - receipt.BlockNumber.Uint64()
	return confirmed >= confirmations, nil
}

func (t transactionVerifierImpl) VerifyTransactionRemote(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error) {
	count, err := t.oracleContract.CountIopNodes(nil)
	if err != nil {
		return false, fmt.Errorf("count iop nodes: %w", err)
	}
	for i := int64(0); i < count.Int64(); i++ {
		_, _, err := t.oracleContract.FindIopNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return false, fmt.Errorf("count iop nodes: %w", err)
		}
	}
	return false, err
}

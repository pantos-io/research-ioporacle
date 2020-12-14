package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionVerifier interface {
	VerifyTransaction(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error)
}

type transactionVerifierImpl struct {
	ethClient *ethclient.Client
}

func NewTransactionVerifier(ethClient *ethclient.Client) *transactionVerifierImpl {
	return &transactionVerifierImpl{
		ethClient: ethClient,
	}
}

func (t *transactionVerifierImpl) VerifyTransaction(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error) {
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

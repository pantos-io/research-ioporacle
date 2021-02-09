package iop

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"math/big"
)

type ValidateTransactionResult struct {
	id        *big.Int
	valid     bool
	signature []byte
}

type Validator struct {
	suite     pairing.Suite
	dkg       *DistKeyGenerator
	ethClient *ethclient.Client
}

func NewValidator(suite pairing.Suite, ethClient *ethclient.Client) *Validator {
	return &Validator{
		suite:     suite,
		ethClient: ethClient,
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (*ValidateTransactionResult, error) {
	receipt, err := v.ethClient.TransactionReceipt(ctx, txHash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil && found {
		return nil, fmt.Errorf("transaction receipt: %w", err)
	}

	blockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	valid := false
	if found {
		confirmed := blockNumber - receipt.BlockNumber.Uint64()
		valid = confirmed >= confirmations
	}

	message, err := encodeValidateTransactionResult(id, valid)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}

	distKey, err := v.dkg.DistKeyShare()
	if err != nil {
		return nil, fmt.Errorf("dist key share: %w", err)
	}

	sig, err := tbls.Sign(v.suite, distKey.Share, message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %w", err)
	}

	return &ValidateTransactionResult{
		id,
		valid,
		sig,
	}, nil
}

func (v *Validator) SetDistKeyGenerator(dkg *DistKeyGenerator) {
	v.dkg = dkg
}

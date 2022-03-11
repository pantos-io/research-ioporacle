package iop

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/tbls"
)

const CONFIRMATIONS uint64 = 5

type ValidateResult struct {
	hash        common.Hash
	valid       bool
	blockNumber *big.Int
	signature   []byte
}

type Validator struct {
	sync.RWMutex
	suite          pairing.Suite
	oracleContract *OracleContract
	dkg            *DistKeyGenerator
	ethClient      *ethclient.Client
}

func NewValidator(suite pairing.Suite, oracleContract *OracleContract, ethClient *ethclient.Client) *Validator {
	return &Validator{
		suite:          suite,
		oracleContract: oracleContract,
		ethClient:      ethClient,
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
	receipt, err := v.ethClient.TransactionReceipt(ctx, hash)
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
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid)
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

	return &ValidateResult{
		hash,
		valid,
		receipt.BlockNumber,
		sig,
	}, nil
}

func (v *Validator) ValidateBlock(ctx context.Context, hash common.Hash) (*ValidateResult, error) {
	block, err := v.ethClient.BlockByHash(ctx, hash)
	found := !errors.Is(err, ethereum.NotFound)
	if err != nil && found {
		return nil, fmt.Errorf("block: %w", err)
	}

	latestBlockNumber, err := v.ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("blocknumber: %w", err)
	}

	var blockNumber *big.Int
	valid := false
	if found {
		blockNumber = block.Number()
		confirmed := latestBlockNumber - block.NumberU64()
		valid = confirmed >= CONFIRMATIONS
	}

	message, err := encodeValidateResult(hash, valid)
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

	return &ValidateResult{
		hash,
		valid,
		blockNumber,
		sig,
	}, nil
}

func (v *Validator) SetDistKeyGenerator(dkg *DistKeyGenerator) {
	v.dkg = dkg
}

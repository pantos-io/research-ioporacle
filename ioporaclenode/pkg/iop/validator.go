package iop

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/tbls"
)

const CONFIRMATIONS uint64 = 6

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
	txRequests     map[common.Hash]*OracleContractValidateTransactionLog
	blockRequests  map[common.Hash]*OracleContractValidateBlockLog
	dkg            *DistKeyGenerator
	ethClient      *ethclient.Client
}

func NewValidator(suite pairing.Suite, oracleContract *OracleContract, ethClient *ethclient.Client) *Validator {
	return &Validator{
		suite:          suite,
		oracleContract: oracleContract,
		txRequests:     make(map[common.Hash]*OracleContractValidateTransactionLog),
		blockRequests:  make(map[common.Hash]*OracleContractValidateBlockLog),
		ethClient:      ethClient,
	}
}

func (v *Validator) WatchAndHandleValidateTransactionLog(ctx context.Context) error {
	sink := make(chan *OracleContractValidateTransactionLog)
	defer close(sink)

	sub, err := v.oracleContract.WatchValidateTransactionLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			v.Lock()
			v.txRequests[event.Hash] = event
			v.Unlock()
			log.Infof("Stored request for transaction %s", common.Hash(event.Hash))
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (v *Validator) WatchAndHandleValidateBlockLog(ctx context.Context) error {
	sink := make(chan *OracleContractValidateBlockLog)
	defer close(sink)

	sub, err := v.oracleContract.WatchValidateBlockLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			v.Lock()
			v.blockRequests[event.Hash] = event
			v.Unlock()
			log.Infof("Stored request for block %s", common.Hash(event.Hash))
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, hash common.Hash) (*ValidateResult, error) {

	var request *OracleContractValidateTransactionLog

loop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			v.RLock()
			if r, ok := v.txRequests[hash]; ok {
				request = r
				v.RUnlock()
				break loop
			}
			v.RUnlock()
			time.Sleep(500 * time.Millisecond)
		}
	}

	receipt, err := v.ethClient.TransactionReceipt(ctx, request.Hash)
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

	var request *OracleContractValidateBlockLog

loop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			v.RLock()
			if r, ok := v.blockRequests[hash]; ok {
				request = r
				v.RUnlock()
				break loop
			}
			v.RUnlock()
			time.Sleep(500 * time.Millisecond)
		}
	}

	block, err := v.ethClient.BlockByHash(ctx, request.Hash)
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

package iop

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"math/big"
	"sync"
	"time"
)

type ValidateTransactionResult struct {
	id          *big.Int
	valid       bool
	blockNumber *big.Int
	signature   []byte
}

type Validator struct {
	sync.RWMutex
	suite          pairing.Suite
	oracleContract *OracleContract
	requests       map[uint64]*OracleContractValidateTransactionLog
	dkg            *DistKeyGenerator
	ethClient      *ethclient.Client
}

func NewValidator(suite pairing.Suite, oracleContract *OracleContract, ethClient *ethclient.Client) *Validator {
	return &Validator{
		suite:          suite,
		oracleContract: oracleContract,
		requests:       make(map[uint64]*OracleContractValidateTransactionLog),
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
			v.requests[event.Id.Uint64()] = event
			v.Unlock()
			log.Infof("Stored request %d", event.Id.Uint64())
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (v *Validator) ValidateTransaction(ctx context.Context, id *big.Int) (*ValidateTransactionResult, error) {

	var request *OracleContractValidateTransactionLog

loop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			v.RLock()
			if r, ok := v.requests[id.Uint64()]; ok {
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
		valid = confirmed >= request.Confirmations.Uint64()
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
		receipt.BlockNumber,
		sig,
	}, nil
}

func (v *Validator) SetDistKeyGenerator(dkg *DistKeyGenerator) {
	v.dkg = dkg
}

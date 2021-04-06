package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"math/big"
	"sync"
	"time"
)

type Aggregator struct {
	suite             pairing.Suite
	ethClient         *ethclient.Client
	dkg               *DistKeyGenerator
	connectionManager *ConnectionManager
	oracleContract    *OracleContract
	registryContract  *RegistryContractWrapper
	account           common.Address
	ecdsaPrivateKey   *ecdsa.PrivateKey
	t                 int
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	oracleContract *OracleContract,
	registryContract *RegistryContractWrapper,
	account common.Address,
	ecdsaPrivateKey *ecdsa.PrivateKey,
) *Aggregator {
	return &Aggregator{
		suite:             suite,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		account:           account,
		ecdsaPrivateKey:   ecdsaPrivateKey,
	}
}

func (a *Aggregator) WatchAndHandleValidateTransactionLog(ctx context.Context) error {
	sink := make(chan *OracleContractValidateTransactionLog)
	defer close(sink)

	sub, err := a.oracleContract.WatchValidateTransactionLog(
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
			log.Infof("Received verify transaction event %s", event.Id)
			isAggregator, err := a.registryContract.IsAggregator(nil, a.account)
			if err != nil {
				log.Errorf("is aggregator: %v", err)
				continue
			}
			if !isAggregator {
				continue
			}
			if err := a.HandleValidateTransactionLog(ctx, event); err != nil {
				log.Errorf("handle validate transaction log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (a *Aggregator) HandleValidateTransactionLog(ctx context.Context, event *OracleContractValidateTransactionLog) error {
	result, s, err := a.AggregateValidateTransactionResults(
		ctx,
		event.Id,
		common.BytesToHash(event.Hash[:]),
		event.Confirmations.Uint64(),
	)
	if err != nil {
		return fmt.Errorf("aggregate validate transaction results: %w", err)
	}

	sig, err := SignatureToBig(s)
	if err != nil {
		return fmt.Errorf("signature to big int: %w", err)
	}

	auth := bind.NewKeyedTransactor(a.ecdsaPrivateKey)
	if _, err = a.oracleContract.SubmitValidationResult(auth, event.Id, result, sig); err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}
	return nil
}

func (a *Aggregator) AggregateValidateTransactionResults(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (bool, []byte, error) {

	positiveResults := make([][]byte, 0)
	negativeResults := make([][]byte, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	nodes, err := a.registryContract.FindOracleNodes()
	if err != nil {
		return false, nil, fmt.Errorf("find nodes: %w", err)
	}

	for _, node := range nodes {

		conn, err := a.connectionManager.FindByAddress(node.Addr)
		if err != nil {
			log.Errorf("find connection by address: %v", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			client := NewOracleNodeClient(conn)
			ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
			result, err := client.ValidateTransaction(ctxTimeout, &ValidateTransactionRequest{
				Id: id.Int64(),
			})
			cancel()
			if err != nil {
				log.Errorf("validate transaction: %v", err)
				return
			}

			mutex.Lock()
			if result.Valid {
				positiveResults = append(positiveResults, result.Signature)
			} else {
				negativeResults = append(negativeResults, result.Signature)
			}
			mutex.Unlock()
		}()
	}

	wg.Wait()

	distKey, err := a.dkg.DistKeyShare()
	if err != nil {
		return false, nil, fmt.Errorf("dist key share: %w", err)
	}

	pubPoly := share.NewPubPoly(a.suite.G2(), a.suite.G2().Point().Base(), distKey.Commits)

	result := false
	sigShares := negativeResults

	if len(positiveResults) >= len(negativeResults) {
		result = true
		sigShares = positiveResults
	}

	message, err := encodeValidateTransactionResult(id, result)
	if err != nil {
		return false, nil, fmt.Errorf("encode validate transaction result: %w", err)
	}

	signature, err := tbls.Recover(a.suite, pubPoly, message, sigShares, a.t, len(nodes))
	if err != nil {
		return false, nil, fmt.Errorf("recover signature: %w", err)
	}

	return result, signature, nil
}

func (a *Aggregator) SetDistKeyGenerator(dkg *DistKeyGenerator) {
	a.dkg = dkg
}

func (a *Aggregator) SetThreshold(threshold int) {
	a.t = threshold
}

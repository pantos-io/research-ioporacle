package iop

import (
	"context"
	"fmt"
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
	node              *OracleNode
	connectionManager *ConnectionManager
	registryContract  *RegistryContractWrapper
	t                 int
}

func NewAggregator(
	suite pairing.Suite,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	registryContract *RegistryContractWrapper,
) *Aggregator {
	return &Aggregator{
		suite:             suite,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		registryContract:  registryContract,
	}
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
				Id:            id.Int64(),
				Tx:            txHash.String(),
				Confirmations: confirmations,
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

	distKey, err := a.node.DistKeyShare()
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

func (a *Aggregator) SetNode(node *OracleNode) {
	a.node = node
}

func (a *Aggregator) SetThreshold(threshold int) {
	a.t = threshold
}

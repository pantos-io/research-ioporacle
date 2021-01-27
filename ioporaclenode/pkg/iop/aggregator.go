package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3/share"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"google.golang.org/grpc"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"math/big"
	"sync"
	"time"
)

type Aggregator interface {
	Aggregate(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (bool, [][]byte, error)
}

type aggregatorImpl struct {
	ethClient        *ethclient.Client
	registryContract *RegistryContractWrapper
	distKey          *dkg.DistKeyShare
	nodes            []RegistryContractOracleNode
	t                int
}

func NewAggregator(
	ethClient *ethclient.Client,
	registryContract *RegistryContractWrapper,
	distKey *dkg.DistKeyShare,
	nodes []RegistryContractOracleNode,
	t int,
) *aggregatorImpl {
	return &aggregatorImpl{
		ethClient:        ethClient,
		registryContract: registryContract,
		distKey:          distKey,
		nodes:            nodes,
		t:                t,
	}
}

func (a *aggregatorImpl) Aggregate(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (bool, []byte, error) {

	positiveResults := make([][]byte, 0)
	negativeResults := make([][]byte, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range a.nodes {
		wg.Add(1)
		go func(node RegistryContractOracleNode) {
			defer wg.Done()
			conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
			if err != nil {
				log.Errorf("dial %s: %v", node.IpAddr, err)
				return
			}
			client := NewOracleNodeClient(conn)
			ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
			result, err := client.ValidateTransaction(ctxTimeout, &ValidateTransactionRequest{
				Id:            id.Int64(),
				Tx:            txHash.String(),
				Confirmations: confirmations,
			})
			cancel()
			_ = conn.Close()
			if err != nil {
				log.Errorf("verify transaction %s: %v", node.IpAddr, err)
				return
			}
			mutex.Lock()

			if result.Valid {
				positiveResults = append(positiveResults, result.Signature)
			} else {
				negativeResults = append(negativeResults, result.Signature)
			}
			mutex.Unlock()
		}(v)
	}

	wg.Wait()

	suite := bn256.NewSuite()
	suiteG2 := bn256.NewSuiteG2()
	pubPoly := share.NewPubPoly(suiteG2, suiteG2.Point().Base(), a.distKey.Commits)

	result := false
	sigShares := negativeResults

	if len(positiveResults) >= len(negativeResults) {
		result = true
		sigShares = positiveResults
	}

	message, err := encodeVerificationResult(id, result)
	if err != nil {
		return false, nil, fmt.Errorf("encode verification result: %w", err)
	}

	signature, err := tbls.Recover(suite, pubPoly, message, sigShares, a.t, len(a.nodes))
	if err != nil {
		return false, nil, fmt.Errorf("recover signature: %v", err)
	}

	return result, signature, nil
}

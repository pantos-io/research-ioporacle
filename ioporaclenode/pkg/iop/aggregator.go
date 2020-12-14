package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"math"
	"math/big"
	"sync"
	"time"
)

type Aggregator interface {
	Aggregate(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (bool, [][]byte, error)
}

type aggregatorImpl struct {
	ethClient        *ethclient.Client
	registryContract *RegistryContract
	account          common.Address
}

func NewAggregator(ethClient *ethclient.Client, registryContract *RegistryContract, account common.Address) *aggregatorImpl {
	return &aggregatorImpl{
		ethClient:        ethClient,
		registryContract: registryContract,
		account:          account,
	}
}

func (a *aggregatorImpl) Aggregate(ctx context.Context, id *big.Int, txHash common.Hash, confirmations uint64) (bool, [][]byte, error) {
	nodes, err := a.FindIopNodes()
	if err != nil {
		return false, nil, fmt.Errorf("find nodes: %w", err)
	}

	positiveResults := make([][]byte, 0)
	negativeResults := make([][]byte, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range nodes {
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
			result, err := client.VerifyTransaction(ctxTimeout, &VerifyTransactionRequest{
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

			if result.Result {
				positiveResults = append(positiveResults, result.Signature)
			} else {
				negativeResults = append(negativeResults, result.Signature)
			}
			mutex.Unlock()
		}(v)
	}

	wg.Wait()
	majority := int(math.Ceil(float64(len(nodes)+2) / 2.0))
	if len(positiveResults) >= majority {
		return true, positiveResults, nil
	} else if len(negativeResults) >= majority {
		return false, negativeResults, nil
	}

	return false, nil, fmt.Errorf("no majority")
}

func (a *aggregatorImpl) FindIopNodes() (map[common.Address]RegistryContractOracleNode, error) {
	count, err := a.registryContract.CountOracleNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count iop nodes: %w", err)
	}
	nodeEntries := make(map[common.Address]RegistryContractOracleNode)
	for i := int64(0); i < count.Int64(); i++ {
		node, err := a.registryContract.FindOracleNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("count iop nodes: %w", err)
		}
		nodeEntries[node.Addr] = node
	}
	delete(nodeEntries, a.account)
	return nodeEntries, nil
}

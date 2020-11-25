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

type TransactionVerifier interface {
	VerifyTransaction(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error)
	VerifyTransactionRemote(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, []common.Address, error)
}

type transactionVerifierImpl struct {
	ethClient      *ethclient.Client
	oracleContract *OracleContract
	account        common.Address
}

func NewTransactionVerifier(ethClient *ethclient.Client, oracleContract *OracleContract, account common.Address) *transactionVerifierImpl {
	return &transactionVerifierImpl{
		ethClient:      ethClient,
		oracleContract: oracleContract,
		account:        account,
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

func (t *transactionVerifierImpl) VerifyTransactionRemote(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, []common.Address, error) {
	nodes, err := t.FindIopNodes()
	if err != nil {
		return false, nil, fmt.Errorf("find nodes: %w", err)
	}

	positiveResults := make([]common.Address, 0)
	negativeResults := make([]common.Address, 0)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range nodes {
		wg.Add(1)
		go func(node OracleContractOracleNode) {
			defer wg.Done()
			conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())
			if err != nil {
				log.Errorf("dial %s: %v", node.IpAddr, err)
				return
			}
			client := NewOracleNodeClient(conn)
			ctxTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
			result, err := client.VerifyTransaction(ctxTimeout, &VerifyTransactionRequest{
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
				positiveResults = append(positiveResults, node.Addr)
			} else {
				negativeResults = append(negativeResults, node.Addr)
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

func (t *transactionVerifierImpl) FindIopNodes() (map[common.Address]OracleContractOracleNode, error) {
	count, err := t.oracleContract.CountOracleNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count iop nodes: %w", err)
	}
	nodeEntries := make(map[common.Address]OracleContractOracleNode)
	for i := int64(0); i < count.Int64(); i++ {
		node, err := t.oracleContract.FindOracleNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("count iop nodes: %w", err)
		}
		nodeEntries[node.Addr] = node
	}
	delete(nodeEntries, t.account)
	return nodeEntries, nil
}

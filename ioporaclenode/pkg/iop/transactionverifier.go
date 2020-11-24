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
	VerifyTransactionRemote(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error)
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
	log.Infof("%d", confirmed)
	return confirmed >= confirmations, nil
}

func (t *transactionVerifierImpl) VerifyTransactionRemote(ctx context.Context, txHash common.Hash, confirmations uint64) (bool, error) {
	nodes, err := t.FindIopNodes()
	if err != nil {
		return false, fmt.Errorf("find nodes: %w", err)
	}

	results := make(map[string]bool)
	frequency := make(map[bool]int)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, v := range nodes {
		wg.Add(1)
		go func(node oracleNodeEntry) {
			defer wg.Done()
			conn, err := grpc.Dial(node.ipAddr, grpc.WithInsecure())
			if err != nil {
				log.Errorf("dial %s: %v", node.ipAddr, err)
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
				log.Errorf("verify transaction %s: %v", node.ipAddr, err)
				return
			}
			results[node.ipAddr] = result.Result

			mutex.Lock()
			frequency[result.Result] += 1
			mutex.Unlock()
		}(v)
	}

	wg.Wait()
	majority := int(math.Ceil(float64(len(nodes)) / 2.0))
	if frequency[true] > majority {
		return true, nil
	} else if frequency[false] > majority {
		return false, nil
	}

	return false, fmt.Errorf("no majority")
}

func (t *transactionVerifierImpl) FindIopNodes() (map[common.Address]oracleNodeEntry, error) {
	count, err := t.oracleContract.CountIopNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count iop nodes: %w", err)
	}
	nodeEntries := make(map[common.Address]oracleNodeEntry)
	for i := int64(0); i < count.Int64(); i++ {
		addr, ipAddr, err := t.oracleContract.FindIopNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("count iop nodes: %w", err)
		}
		nodeEntries[addr] = oracleNodeEntry{
			address: addr,
			ipAddr:  ipAddr,
		}
	}
	delete(nodeEntries, t.account)
	return nodeEntries, nil
}

package iop

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type RegistryContractWrapper struct {
	*RegistryContract
}

func (r *RegistryContractWrapper) FindIopNodes() ([]RegistryContractOracleNode, error) {
	count, err := r.CountOracleNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodeEntries := make([]RegistryContractOracleNode, count.Int64())
	for i := int64(0); i < count.Int64(); i++ {
		node, err := r.FindOracleNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
		}
		nodeEntries[i] = node
	}
	return nodeEntries, nil
}

func (r *RegistryContractWrapper) FindIopNodesMap() (map[common.Address]RegistryContractOracleNode, error) {
	count, err := r.CountOracleNodes(nil)
	if err != nil {
		return nil, fmt.Errorf("count oracle nodes: %w", err)
	}
	nodeEntries := make(map[common.Address]RegistryContractOracleNode)
	for i := int64(0); i < count.Int64(); i++ {
		node, err := r.FindOracleNodeByIndex(nil, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
		}
		nodeEntries[node.Addr] = node
	}
	return nodeEntries, nil
}
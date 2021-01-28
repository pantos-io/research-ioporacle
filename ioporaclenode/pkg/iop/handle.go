package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"math/big"
	"time"
)

func (n *OracleNode) handleRegisterOracleNodeLog(event *RegistryContractRegisterOracleNodeLog) error {
	node, err := n.registryContract.FindOracleNodeByAddress(nil, event.Sender)
	if err != nil {
		return fmt.Errorf("find oracle node by address %s: %w", event.Sender, err)
	}
	//n.nodes = append(n.nodes, node)
	if _, err := n.connectionManager.NewConnection(node); err != nil {
		return fmt.Errorf("new connection: %w", err)
	}
	log.Infof("New connection to %s", event.Sender.String())
	return nil
}

func (n *OracleNode) handleVerifyTransactionLog(ctx context.Context, event *OracleContractVerifyTransactionLog) error {
	result, sig, err := n.aggregator.AggregateValidateTransactionResults(ctx, event.Id, common.BytesToHash(event.Hash[:]), event.Confirmations.Uint64())
	if err != nil {
		return fmt.Errorf("verify transaction remote: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	_, err = n.oracleContract.SubmitVerification(auth, event.Id, result, [2]*big.Int{
		new(big.Int).SetBytes(sig[:32]),
		new(big.Int).SetBytes(sig[32:64]),
	})
	if err != nil {
		return fmt.Errorf("verify transaction result: %v", err)
	}
	return nil
}

func (n *OracleNode) handleDistributedKeyGenerationLog(event *DistKeyContractDistKeyGenerationLog) error {

	time.Sleep(5 * time.Second)

	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}

	pubKeys := make([]kyber.Point, len(nodes))
	for i, node := range nodes {
		pubKey := n.suiteG2.Point()
		if err := pubKey.UnmarshalBinary(node.PubKey); err != nil {
			return fmt.Errorf("unmarshal public key: %w", err)
		}
		pubKeys[i] = pubKey
	}

	threshold := int(event.Threshold.Int64())
	distKeyGenerator, err := dkg.NewDistKeyGenerator(
		n.suiteG2,
		n.blsPrivateKey,
		pubKeys,
		threshold,
	)
	if err != nil {
		return fmt.Errorf("new dkg: %w", err)
	}
	n.dkg = distKeyGenerator
	n.aggregator.SetThreshold(threshold)
	time.Sleep(5 * time.Second)

	deals, err := n.dkg.Deals()
	if err != nil {
		return fmt.Errorf("deals: %w", err)
	}

	n.broadCastDeals(nodes, deals)

	timeout := time.After(30 * time.Second)

loop:
	for {
		select {
		case <-timeout:
			n.dkg.SetTimeout()
			break loop
		default:
			if n.dkg.Certified() {
				break loop
			}
			time.Sleep(50 * time.Millisecond)
		}
	}

	log.Infof("Qualified shares: %v\n", n.dkg.QualifiedShares())
	log.Infof("QUAL: %v\n", n.dkg.QUAL())

	distKey, err := n.dkg.DistKeyShare()
	if err != nil {
		return fmt.Errorf("distributed key share: %w", err)
	}

	log.Infof("Public Key: %v", distKey.Public())
	pubBinary, err := distKey.Public().MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal public key: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	_, err = n.distKeyContract.SetPublicKey(auth, [4]*big.Int{
		new(big.Int).SetBytes(pubBinary[:32]),
		new(big.Int).SetBytes(pubBinary[32:64]),
		new(big.Int).SetBytes(pubBinary[64:96]),
		new(big.Int).SetBytes(pubBinary[96:128]),
	})
	if err != nil {
		return fmt.Errorf("set public key: %w", err)
	}

	return nil
}

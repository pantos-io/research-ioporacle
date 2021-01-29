package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"time"
)

func (n *OracleNode) handleRegisterOracleNodeLog(event *RegistryContractRegisterOracleNodeLog) error {
	node, err := n.registryContract.FindOracleNodeByAddress(nil, event.Sender)
	if err != nil {
		return fmt.Errorf("find oracle node by address %s: %w", event.Sender, err)
	}
	if _, err := n.connectionManager.NewConnection(node); err != nil {
		return fmt.Errorf("new connection: %w", err)
	}
	log.Infof("New connection to %s", event.Sender.String())
	return nil
}

func (n *OracleNode) handleVerifyTransactionLog(ctx context.Context, event *OracleContractValidateTransactionLog) error {
	result, s, err := n.aggregator.AggregateValidateTransactionResults(
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

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	if _, err = n.oracleContract.SubmitValidationResult(auth, event.Id, result, sig); err != nil {
		return fmt.Errorf("submit verification: %w", err)
	}
	return nil
}

func (n *OracleNode) handleDistributedKeyGenerationLog(event *DistKeyContractDistKeyGenerationLog) error {

	nodes, err := n.registryContract.FindOracleNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}

	pubKeys := make([]kyber.Point, len(nodes))
	for i, node := range nodes {
		pubKey := n.suite.Point()
		if err := pubKey.UnmarshalBinary(node.PubKey); err != nil {
			return fmt.Errorf("unmarshal public key: %w", err)
		}
		pubKeys[i] = pubKey
	}

	threshold := int(event.Threshold.Int64())
	n.aggregator.SetThreshold(threshold)
	n.dkg, err = dkg.NewDistKeyGenerator(
		n.suite,
		n.blsPrivateKey,
		pubKeys,
		threshold,
	)
	if err != nil {
		return fmt.Errorf("new dkg: %w", err)
	}

	//Wait until every participant is prepared. TODO: Wait for head n
	time.Sleep(5 * time.Second)

	deals, err := n.dkg.Deals()
	if err != nil {
		return fmt.Errorf("deals: %w", err)
	}

	n.broadCastDeals(nodes, deals)

loop:
	for {
		select {
		case <-time.After(DkgTimeout):
			n.dkg.SetTimeout()
			break loop
		default:
			if n.dkg.Certified() {
				break loop
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

	log.Infof("Qualified shares: %v", n.dkg.QualifiedShares())
	log.Infof("QUAL: %v", n.dkg.QUAL())

	distKey, err := n.dkg.DistKeyShare()
	if err != nil {
		return fmt.Errorf("distributed key share: %w", err)
	}

	log.Infof("Public Key: %v", distKey.Public())
	pubKey, err := PubKeyToBig(distKey.Public())
	if err != nil {
		return fmt.Errorf("public key to big int: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	if _, err = n.distKeyContract.SetPublicKey(auth, pubKey); err != nil {
		return fmt.Errorf("set public key: %w", err)
	}

	return nil
}

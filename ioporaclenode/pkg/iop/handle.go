package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
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

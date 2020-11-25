package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type oracleNode struct {
	UnimplementedOracleNodeServer
	server         *grpc.Server
	txVerifier     TransactionVerifier
	oracleContract *OracleContract
	privateKey     *ecdsa.PrivateKey
	account        common.Address
}

func NewOracleNode(txVerifier TransactionVerifier, oracleContract *OracleContract, privateKey *ecdsa.PrivateKey, account common.Address) *oracleNode {
	grpcServer := grpc.NewServer()
	node := &oracleNode{
		server:         grpcServer,
		txVerifier:     txVerifier,
		oracleContract: oracleContract,
		privateKey:     privateKey,
		account:        account,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node
}

func (n *oracleNode) Serve(lis net.Listener) error {
	go func() {
		err := n.watchVerifyTransactionLog(context.Background())
		if err != nil {
			log.Errorf("watch verify transaction log: %v", err)
		}
	}()
	err := n.register(lis.Addr().String())
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}
	return n.server.Serve(lis)
}

func (n *oracleNode) watchVerifyTransactionLog(ctx context.Context) error {
	sink := make(chan *OracleContractVerifyTransactionLog)
	defer close(sink)

	sub, err := n.oracleContract.WatchVerifyTransactionLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			isLeader, err := n.oracleContract.IsLeader(nil, n.account)
			if err != nil {
				log.Errorf("is leader: %v", err)
				continue
			}
			if isLeader {
				err = n.handleVerifyTransactionLog(ctx, event)
				if err != nil {
					log.Errorf("handle verify transaction log: %v", err)
				}
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *oracleNode) handleVerifyTransactionLog(ctx context.Context, event *OracleContractVerifyTransactionLog) error {
	result, err := n.txVerifier.VerifyTransactionRemote(ctx, common.HexToHash(event.Hash), event.Confirmations.Uint64())
	if err != nil {
		return fmt.Errorf("verify transaction remote: %w", err)
	}
	auth := bind.NewKeyedTransactor(n.privateKey)
	_, err = n.oracleContract.SubmitVerification(auth, result)
	if err != nil {
		return fmt.Errorf("verify transaction result: %v", err)
	}
	return nil
}

func (n *oracleNode) VerifyTransaction(ctx context.Context, request *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	log.Infof("%v", request)
	result, err := n.txVerifier.VerifyTransaction(ctx, common.HexToHash(request.Tx), request.Confirmations)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verify transaction: %v", err)
	}
	return &VerifyTransactionResponse{
		Id:     request.Id,
		Result: result,
	}, nil
}

func (n *oracleNode) register(ipAddr string) error {
	isRegistered, err := n.oracleContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.privateKey)
	if !isRegistered {
		_, err = n.oracleContract.RegisterOracleNode(auth, ipAddr)
		if err != nil {
			return fmt.Errorf("register iop node: %w", err)
		}
	}
	return nil
}

func (n *oracleNode) Stop() {
	n.server.Stop()
}

func (n *oracleNode) GracefulStop() {
	n.server.GracefulStop()
}

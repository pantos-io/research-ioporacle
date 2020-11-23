package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type oracleNode struct {
	UnimplementedOracleNodeServer
	server     *grpc.Server
	txVerifier TransactionVerifier
	iopOracle  *IopOracleContract
	privateKey *ecdsa.PrivateKey
	account    string
}

func NewOracleNode(txVerifier TransactionVerifier, iopOracle *IopOracleContract, privateKey *ecdsa.PrivateKey) (*oracleNode, error) {
	grpcServer := grpc.NewServer()
	account, err := AddressFromPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %w", err)
	}
	node := &oracleNode{
		server:     grpcServer,
		txVerifier: txVerifier,
		iopOracle:  iopOracle,
		privateKey: privateKey,
		account:    account,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node, nil
}

func (n *oracleNode) Serve(lis net.Listener) error {
	err := n.register(lis.Addr().String())
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}
	return n.server.Serve(lis)
}

func (n *oracleNode) Stop() {
	n.server.Stop()
}

func (n *oracleNode) GracefulStop() {
	n.server.GracefulStop()
}

func (n *oracleNode) register(ipAddr string) error {
	isRegistered, err := n.iopOracle.IopNodeIsRegistered(nil, common.HexToAddress(n.account))
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.privateKey)
	if !isRegistered {
		_, err = n.iopOracle.RegisterIopNode(auth, ipAddr)
		if err != nil {
			return fmt.Errorf("register iop node: %w", err)
		}
	}
	return nil
}

func (n *oracleNode) VerifyTransaction(ctx context.Context, request *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	result, err := n.txVerifier.VerifyTransaction(ctx, common.HexToHash(request.Tx), request.Confirmations)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verify transaction: %v", err)
	}
	return &VerifyTransactionResponse{
		Id:     request.Id,
		Result: result,
	}, nil
}

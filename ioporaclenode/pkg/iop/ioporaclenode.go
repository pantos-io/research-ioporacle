package iop

import (
	"context"
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
}

func NewOracleNode(txVerifier TransactionVerifier) *oracleNode {
	grpcServer := grpc.NewServer()
	node := &oracleNode{
		server:     grpcServer,
		txVerifier: txVerifier,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node
}

func (n *oracleNode) Serve(lis net.Listener) error {
	return n.server.Serve(lis)
}

func (n *oracleNode) Stop() {
	n.server.Stop()
}

func (n *oracleNode) GracefulStop() {
	n.server.GracefulStop()
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

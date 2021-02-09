package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
)

func (n *OracleNode) SendDeal(_ context.Context, request *SendDealRequest) (*SendDealResponse, error) {
	_, err := n.dkg.ProcessDeal(pbToDeal(request.Deal))
	if err != nil {
		return nil, fmt.Errorf("handle deal: %w", err)
	}
	return &SendDealResponse{}, nil
}

func (n *OracleNode) ValidateTransaction(ctx context.Context, request *ValidateTransactionRequest) (*ValidateTransactionResponse, error) {
	result, err := n.validator.ValidateTransaction(
		ctx,
		big.NewInt(request.Id),
		common.HexToHash(request.Tx),
		request.Confirmations,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate transaction: %v", err)
	}
	return validateTransactionResultToResponse(result), nil
}

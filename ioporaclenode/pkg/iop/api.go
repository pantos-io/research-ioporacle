package iop

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
)

func (n *OracleNode) ProcessDeal(_ context.Context, request *ProcessDealRequest) (*ProcessDealResponse, error) {
	log.Infof("Process deal from node %d", request.Deal.Index)

	response, err := n.dkg.ProcessDeal(pbToDeal(request.Deal))
	if err != nil {
		return nil, fmt.Errorf("process deal: %w", err)
	}

	n.broadCastResponse(response)

	return &ProcessDealResponse{
		Response: responseToPb(response),
	}, nil
}

func (n *OracleNode) ProcessResponse(ctx context.Context, request *ProcessResponseRequest) (*ProcessResponseResponse, error) {
	log.Infof("Process response with dealer %d and verifier %d", request.Response.Index, request.Response.Response.Index)

	for {
		select {
		case <-ctx.Done():
			return &ProcessResponseResponse{}, ctx.Err()
		default:
			_, err := n.dkg.ProcessResponse(pbToResponse(request.Response))
			if errors.Is(err, vss.ErrNoDealBeforeResponse) {
				continue
			} else if err != nil {
				return &ProcessResponseResponse{}, fmt.Errorf("process response: %w", err)
			}
			return &ProcessResponseResponse{
				Justification: nil,
			}, nil
		}
	}
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

package iop

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
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

func dealToPb(deal *dkg.Deal) *Deal {
	return &Deal{
		Index:     deal.Index,
		Deal:      encryptedDealToPb(deal.Deal),
		Signature: deal.Signature,
	}
}

func pbToDeal(deal *Deal) *dkg.Deal {
	return &dkg.Deal{
		Index:     deal.Index,
		Deal:      pbToEncryptedDeal(deal.Deal),
		Signature: deal.Signature,
	}
}

func encryptedDealToPb(encryptedDeal *vss.EncryptedDeal) *EncryptedDeal {
	return &EncryptedDeal{
		DhKey:     encryptedDeal.DHKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func pbToEncryptedDeal(encryptedDeal *EncryptedDeal) *vss.EncryptedDeal {
	return &vss.EncryptedDeal{
		DHKey:     encryptedDeal.DhKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func validateTransactionResultToResponse(result *ValidateTransactionResult) *ValidateTransactionResponse {
	return &ValidateTransactionResponse{
		Id:        result.id.Int64(),
		Valid:     result.valid,
		Signature: result.signature,
	}
}

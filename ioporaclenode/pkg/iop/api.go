package iop

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (n *OracleNode) SendDeal(_ context.Context, request *SendDealRequest) (*SendDealResponse, error) {
	_, err := n.dkg.ProcessDeal(PbToDeal(request.Deal))
	if err != nil {
		return nil, fmt.Errorf("handle deal: %w", err)
	}
	return &SendDealResponse{}, nil
}

func (n *OracleNode) ValidateTransaction(ctx context.Context, request *ValidateRequest) (*ValidateResponse, error) {
	result, err := n.validator.ValidateTransaction(
		ctx,
		common.BytesToHash(request.Hash),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate transaction: %v", err)
	}
	return ValidateResultToResponse(result), nil
}

func (n *OracleNode) ValidateBlock(ctx context.Context, request *ValidateRequest) (*ValidateResponse, error) {
	result, err := n.validator.ValidateBlock(
		ctx,
		common.BytesToHash(request.Hash),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "validate block: %v", err)
	}
	return ValidateResultToResponse(result), nil
}

func DealToPb(deal *dkg.Deal) *Deal {
	return &Deal{
		Index:     deal.Index,
		Deal:      EncryptedDealToPb(deal.Deal),
		Signature: deal.Signature,
	}
}

func PbToDeal(deal *Deal) *dkg.Deal {
	return &dkg.Deal{
		Index:     deal.Index,
		Deal:      PbToEncryptedDeal(deal.Deal),
		Signature: deal.Signature,
	}
}

func EncryptedDealToPb(encryptedDeal *vss.EncryptedDeal) *EncryptedDeal {
	return &EncryptedDeal{
		DhKey:     encryptedDeal.DHKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func PbToEncryptedDeal(encryptedDeal *EncryptedDeal) *vss.EncryptedDeal {
	return &vss.EncryptedDeal{
		DHKey:     encryptedDeal.DhKey,
		Signature: encryptedDeal.Signature,
		Nonce:     encryptedDeal.Nonce,
		Cipher:    encryptedDeal.Cipher,
	}
}

func ValidateResultToResponse(result *ValidateResult) *ValidateResponse {
	resp := &ValidateResponse{
		Hash:        result.hash[:],
		Valid:       result.valid,
		Signature:   result.signature,
	}

	if result.blockNumber != nil {
		resp.BlockNumber = result.blockNumber.Int64()
	}

	return resp
}

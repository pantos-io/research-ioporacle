package iop

import (
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
)

func validateTransactionResultToResponse(result *ValidateTransactionResult) *ValidateTransactionResponse {
	return &ValidateTransactionResponse{
		Id:        result.id.Int64(),
		Valid:     result.valid,
		Signature: result.signature,
	}
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

func responseToPb(response *dkg.Response) *Response {
	return &Response{
		Index:    response.Index,
		Response: vssResponseToPb(response.Response),
	}
}

func pbToResponse(response *Response) *dkg.Response {
	return &dkg.Response{
		Index:    response.Index,
		Response: pbToVSSResponse(response.Response),
	}
}

func pbToVSSResponse(response *VSSResponse) *vss.Response {
	return &vss.Response{
		SessionID: response.SessionID,
		Index:     response.Index,
		Status:    response.Status,
		Signature: response.Signature,
	}
}

func vssResponseToPb(response *vss.Response) *VSSResponse {
	return &VSSResponse{
		SessionID: response.SessionID,
		Index:     response.Index,
		Status:    response.Status,
		Signature: response.Signature,
	}
}

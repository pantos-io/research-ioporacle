package iop

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go.dedis.ch/kyber/v3"
	"math/big"
)

func AddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("could not cast to public key ecdsa")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

func HexToScalar(suite kyber.Group, hexScalar string) (kyber.Scalar, error) {
	b, err := hex.DecodeString(hexScalar)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in scalar", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for scalar")
	}
	s := suite.Scalar()
	if err := s.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("unmarshal scalar binary: %w", err)
	}
	return s, nil
}

func PubKeyToBig(point kyber.Point) ([4]*big.Int, error) {
	b, err := point.MarshalBinary()
	if err != nil {
		return [4]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(b) != 128 {
		return [4]*big.Int{}, fmt.Errorf("invalid public key length")
	}

	return [4]*big.Int{
		new(big.Int).SetBytes(b[:32]),
		new(big.Int).SetBytes(b[32:64]),
		new(big.Int).SetBytes(b[64:96]),
		new(big.Int).SetBytes(b[96:128]),
	}, nil
}

func SignatureToBig(sig []byte) ([2]*big.Int, error) {
	if len(sig) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid signature length")
	}
	return [2]*big.Int{
		new(big.Int).SetBytes(sig[:32]),
		new(big.Int).SetBytes(sig[32:64]),
	}, nil
}

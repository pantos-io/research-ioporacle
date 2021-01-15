package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/share"
	"go.dedis.ch/kyber/v3/sign/tbls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
	"net"
)

type oracleNode struct {
	UnimplementedOracleNodeServer
	server           *grpc.Server
	txVerifier       TransactionVerifier
	aggregator       Aggregator
	ethClient        *ethclient.Client
	oracleContract   *OracleContract
	registryContract *RegistryContract
	ecdsaPrivateKey  *ecdsa.PrivateKey
	blsPrivateKey    kyber.Scalar
	blsPublicKey     kyber.Point
	account          common.Address
	suite            pairing.Suite
	priShare         *share.PriShare
}

func NewOracleNode(
	ethClient *ethclient.Client,
	txVerifier TransactionVerifier,
	aggregator Aggregator,
	oracleContract *OracleContract,
	registryContract *RegistryContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	blsPrivateKey kyber.Scalar,
	account common.Address,
	suite pairing.Suite,
) *oracleNode {
	grpcServer := grpc.NewServer()
	node := &oracleNode{
		server:           grpcServer,
		ethClient:        ethClient,
		txVerifier:       txVerifier,
		aggregator:       aggregator,
		oracleContract:   oracleContract,
		registryContract: registryContract,
		ecdsaPrivateKey:  ecdsaPrivateKey,
		blsPrivateKey:    blsPrivateKey,
		blsPublicKey:     suite.G2().Point().Mul(blsPrivateKey, nil),
		account:          account,
		suite:            suite,
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
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			isLeader, err := n.registryContract.IsLeader(nil, n.account)
			if err != nil {
				log.Errorf("is leader: %v", err)
				continue
			}
			if !isLeader {
				continue
			}
			err = n.handleVerifyTransactionLog(ctx, event)
			if err != nil {
				log.Errorf("handle verify transaction log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *oracleNode) handleVerifyTransactionLog(ctx context.Context, event *OracleContractVerifyTransactionLog) error {
	result, _, err := n.aggregator.Aggregate(ctx, event.Id, common.BytesToHash(event.Hash[:]), event.Confirmations.Uint64())
	if err != nil {
		return fmt.Errorf("verify transaction remote: %w", err)
	}
	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	_, err = n.oracleContract.SubmitVerification(auth, event.Id, result, [2]*big.Int{big.NewInt(0), big.NewInt(0)})
	if err != nil {
		return fmt.Errorf("verify transaction result: %v", err)
	}
	return nil
}

func (n *oracleNode) VerifyTransaction(ctx context.Context, request *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	result, err := n.txVerifier.VerifyTransaction(ctx, common.HexToHash(request.Tx), request.Confirmations)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verify transaction: %v", err)
	}

	uint256Ty, _ := abi.NewType("uint256", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: boolTy,
		},
	}

	message, _ := arguments.Pack(
		big.NewInt(request.Id),
		result,
	)

	sig, err := tbls.Sign(n.suite, n.priShare, message)
	if err != nil {
		return nil, fmt.Errorf("tbls sign: %v", err)
	}

	return &VerifyTransactionResponse{
		Id:        request.Id,
		Result:    result,
		Signature: sig,
	}, nil
}

func (n *oracleNode) register(ipAddr string) error {
	isRegistered, err := n.registryContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	b, err := n.blsPublicKey.MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal bls public key: %v", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	if !isRegistered {
		_, err = n.registryContract.RegisterOracleNode(auth, ipAddr, b)
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

package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
	"net"
)

type oracleNode struct {
	UnimplementedOracleNodeServer
	server         *grpc.Server
	txVerifier     TransactionVerifier
	ethClient      *ethclient.Client
	oracleContract *OracleContract
	privateKey     *ecdsa.PrivateKey
	account        common.Address
}

func NewOracleNode(ethClient *ethclient.Client, txVerifier TransactionVerifier, oracleContract *OracleContract, privateKey *ecdsa.PrivateKey, account common.Address) *oracleNode {
	grpcServer := grpc.NewServer()
	node := &oracleNode{
		server:         grpcServer,
		ethClient:      ethClient,
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
	go func() {
		err := n.watchSubmitVerificationLog(context.Background())
		if err != nil {
			log.Errorf("watch submit verification log: %v", err)
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
	result, witnesses, err := n.txVerifier.VerifyTransactionRemote(ctx, common.HexToHash(event.Hash), event.Confirmations.Uint64())
	if err != nil {
		return fmt.Errorf("verify transaction remote: %w", err)
	}
	auth := bind.NewKeyedTransactor(n.privateKey)
	_, err = n.oracleContract.SubmitVerification(auth, event.Id, result, witnesses)
	if err != nil {
		return fmt.Errorf("verify transaction result: %v", err)
	}
	return nil
}

func (n *oracleNode) watchSubmitVerificationLog(ctx context.Context) error {
	sink := make(chan *OracleContractSubmitVerificationLog)
	defer close(sink)

	sub, err := n.oracleContract.WatchSubmitVerificationLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		[]common.Address{n.account},
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			err = n.handleSubmitVerificationLog(ctx, event)
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

func (n *oracleNode) handleSubmitVerificationLog(ctx context.Context, event *OracleContractSubmitVerificationLog) error {
	sink := make(chan *types.Header)
	defer close(sink)

	sub, err := n.ethClient.SubscribeNewHead(context.Background(), sink)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case head := <-sink:
			if head.Number.Cmp(event.Finalized) != 0 {
				continue
			}
			auth := bind.NewKeyedTransactor(n.privateKey)
			_, err := n.oracleContract.TransferReward(auth, event.Request)
			if err != nil {
				return fmt.Errorf("transfer reward %w, err")
			}
			return nil
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *oracleNode) VerifyTransaction(ctx context.Context, request *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	result, err := n.txVerifier.VerifyTransaction(ctx, common.HexToHash(request.Tx), request.Confirmations)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verify transaction: %v", err)
	}
	encodedResult, err := n.oracleContract.EncodeResult(nil, big.NewInt(request.Id), result)
	if err != nil {
		return nil, fmt.Errorf("encode result: %w", err)
	}
	hash := crypto.Keccak256Hash(encodedResult)
	signature, err := crypto.Sign(hash.Bytes(), n.privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign result: %w", err)
	}
	return &VerifyTransactionResponse{
		Id:        request.Id,
		Result:    result,
		Signature: signature,
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

package iop

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	vss "go.dedis.ch/kyber/v3/share/vss/pedersen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"math/big"
	"net"
	"time"
)

type OracleNode struct {
	UnimplementedOracleNodeServer
	server           *grpc.Server
	validator        *Validator
	aggregator       Aggregator
	ethClient        *ethclient.Client
	oracleContract   *OracleContract
	registryContract *RegistryContractWrapper
	raffleContract   *RaffleContract
	distKeyContract  *DistKeyContract
	dkg              *dkg.DistKeyGenerator
	ecdsaPrivateKey  *ecdsa.PrivateKey
	blsPrivateKey    kyber.Scalar
	blsPublicKey     kyber.Point
	account          common.Address
	suite            pairing.Suite
	distKey          *dkg.DistKeyShare
	threshold        int
	clients          map[int]OracleNodeClient
	nodes            []RegistryContractOracleNode
}

func NewOracleNode(
	ethClient *ethclient.Client,
	validator *Validator,
	oracleContract *OracleContract,
	registryContract *RegistryContractWrapper,
	raffleContract *RaffleContract,
	distKeyContract *DistKeyContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	blsPrivateKey kyber.Scalar,
	account common.Address,
	suite pairing.Suite,
) *OracleNode {
	grpcServer := grpc.NewServer()
	node := &OracleNode{
		server:           grpcServer,
		ethClient:        ethClient,
		validator:        validator,
		oracleContract:   oracleContract,
		registryContract: registryContract,
		raffleContract:   raffleContract,
		distKeyContract:  distKeyContract,
		ecdsaPrivateKey:  ecdsaPrivateKey,
		blsPrivateKey:    blsPrivateKey,
		blsPublicKey:     suite.G2().Point().Mul(blsPrivateKey, nil),
		account:          account,
		suite:            suite,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node
}

func (n *OracleNode) Serve(lis net.Listener) error {
	go func() {
		err := n.watchDistKeyGenerationLog(context.Background())
		if err != nil {
			log.Errorf("watch distributed key generation log: %v", err)
		}
	}()
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

func (n *OracleNode) watchDistKeyGenerationLog(ctx context.Context) error {
	sink := make(chan *DistKeyContractDistKeyGenerationLog)
	defer close(sink)

	sub, err := n.distKeyContract.WatchDistKeyGenerationLog(
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
			err = n.handleDistributedKeyGenerationLog(event)
			if err != nil {
				log.Errorf("handle distributed key generation log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *OracleNode) handleDistributedKeyGenerationLog(event *DistKeyContractDistKeyGenerationLog) error {

	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}

	n.nodes = nodes

	pubKeys := make([]kyber.Point, len(nodes))
	for i, node := range nodes {
		pubKey := bn256.NewSuiteG2().Point()
		err = pubKey.UnmarshalBinary(node.PubKey)
		if err != nil {
			return fmt.Errorf("unmarshal public key: %w", err)
		}
		pubKeys[i] = pubKey
	}

	n.clients, err = n.connect(nodes)
	if err != nil {
		return fmt.Errorf("connect to nodes: %w", err)
	}

	n.threshold = int(event.Threshold.Int64())
	distKeyGenerator, err := dkg.NewDistKeyGenerator(bn256.NewSuiteG2(), n.blsPrivateKey, pubKeys, n.threshold)
	if err != nil {
		return fmt.Errorf("dkg: %w", err)
	}
	n.dkg = distKeyGenerator

	time.Sleep(5 * time.Second)

	deals, err := n.dkg.Deals()
	if err != nil {
		return fmt.Errorf("deals: %w", err)
	}

	n.broadCastDeals(deals)

	timeout := time.After(30 * time.Second)

loop:
	for {
		select {
		case <-timeout:
			n.dkg.SetTimeout()
			break loop
		default:
			if n.dkg.Certified() {
				break loop
			}
			time.Sleep(50 * time.Millisecond)
		}
	}

	log.Infof("qualified shares: %v\n", n.dkg.QualifiedShares())
	log.Infof("QUAL: %v\n", n.dkg.QUAL())

	n.distKey, err = n.dkg.DistKeyShare()
	if err != nil {
		return fmt.Errorf("distributed key share: %w", err)
	}

	log.Infof("Public Key: %v", n.distKey.Public())
	pubBinary, err := n.distKey.Public().MarshalBinary()
	if err != nil {
		return fmt.Errorf("marshal public key: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	_, err = n.distKeyContract.SetPublicKey(auth, [4]*big.Int{
		new(big.Int).SetBytes(pubBinary[:32]),
		new(big.Int).SetBytes(pubBinary[32:64]),
		new(big.Int).SetBytes(pubBinary[64:96]),
		new(big.Int).SetBytes(pubBinary[96:128]),
	})
	if err != nil {
		return fmt.Errorf("set public key: %w", err)
	}

	return nil
}

func (n *OracleNode) connect(nodes []RegistryContractOracleNode) (map[int]OracleNodeClient, error) {
	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return nil, fmt.Errorf("find nodes: %w", err)
	}

	otherNodes := make(map[int]OracleNodeClient)
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Addr == n.account {
			continue
		}
		conn, err := grpc.Dial(nodes[i].IpAddr, grpc.WithInsecure())
		if err != nil {
			log.Errorf("dial %s: %v", nodes[i].IpAddr, err)
			continue
		}

		log.Printf("Connected to node %s", nodes[i].Addr.String())
		otherNodes[i] = NewOracleNodeClient(conn)
	}
	return otherNodes, nil
}

func (n *OracleNode) broadCastDeals(deals map[int]*dkg.Deal) {
	log.Infof("Broadcasting deals")
	for i, deal := range deals {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		request := &ProcessDealRequest{
			Deal: dealToPb(deal),
		}
		_, err := n.clients[i].ProcessDeal(ctx, request)
		if err != nil {
			log.Errorf("process deal: %v", err)
		}
		cancel()
	}
}

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

func (n *OracleNode) broadCastResponse(response *dkg.Response) {
	log.Infof("Broadcasting response with dealer %d and verifier %d", response.Index, response.Response.Index)
	for _, otherNode := range n.clients {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		request := &ProcessResponseRequest{
			Response: responseToPb(response),
		}
		otherNode := otherNode
		go func() {
			_, err := otherNode.ProcessResponse(ctx, request)
			if err != nil {
				log.Errorf("process response: %v", err)
			}
			cancel()
		}()
	}
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

func (n *OracleNode) watchVerifyTransactionLog(ctx context.Context) error {
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
			isAggregator, err := n.registryContract.IsAggregator(nil, n.account)
			if err != nil {
				log.Errorf("is aggregator: %v", err)
				continue
			}
			if !isAggregator {
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

func (n *OracleNode) handleVerifyTransactionLog(ctx context.Context, event *OracleContractVerifyTransactionLog) error {
	aggregator := NewAggregator(n.ethClient, n.registryContract, n.account, n.distKey, n.nodes, n.threshold)
	result, sig, err := aggregator.Aggregate(ctx, event.Id, common.BytesToHash(event.Hash[:]), event.Confirmations.Uint64())
	if err != nil {
		return fmt.Errorf("verify transaction remote: %w", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	_, err = n.oracleContract.SubmitVerification(auth, event.Id, result, [2]*big.Int{
		new(big.Int).SetBytes(sig[:32]),
		new(big.Int).SetBytes(sig[32:64]),
	})
	if err != nil {
		return fmt.Errorf("verify transaction result: %v", err)
	}
	return nil
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

func (n *OracleNode) register(ipAddr string) error {
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

func (n *OracleNode) DistKeyShare() (*dkg.DistKeyShare, error) {
	if n.distKey == nil {
		return nil, errors.New("no dist key share")
	}
	return n.distKey, nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()
}

func (n *OracleNode) GracefulStop() {
	n.server.GracefulStop()
}

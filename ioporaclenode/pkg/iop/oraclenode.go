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
	dkg "go.dedis.ch/kyber/v3/share/dkg/pedersen"
	"google.golang.org/grpc"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"net"
	"time"
)

type OracleNode struct {
	UnimplementedOracleNodeServer
	server            *grpc.Server
	connectionManager *ConnectionManager
	validator         *Validator
	aggregator        Aggregator
	ethClient         *ethclient.Client
	oracleContract    *OracleContract
	registryContract  *RegistryContractWrapper
	raffleContract    *RaffleContract
	distKeyContract   *DistKeyContract
	dkg               *dkg.DistKeyGenerator
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	blsPublicKey      kyber.Point
	account           common.Address
	suiteG2           *bn256.Suite
	distKey           *dkg.DistKeyShare
	threshold         int
	nodes             []RegistryContractOracleNode
}

func NewOracleNode(
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	validator *Validator,
	oracleContract *OracleContract,
	registryContract *RegistryContractWrapper,
	raffleContract *RaffleContract,
	distKeyContract *DistKeyContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	blsPrivateKey kyber.Scalar,
	account common.Address,
	suite *bn256.Suite,
) *OracleNode {
	grpcServer := grpc.NewServer()
	node := &OracleNode{
		server:            grpcServer,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		validator:         validator,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		raffleContract:    raffleContract,
		distKeyContract:   distKeyContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		blsPrivateKey:     blsPrivateKey,
		blsPublicKey:      suite.Point().Mul(blsPrivateKey, nil),
		account:           account,
		suiteG2:           suite,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node
}

func (n *OracleNode) Serve(lis net.Listener) error {
	err := n.initConnections()
	if err != nil {
		return fmt.Errorf("init connections: %w", err)
	}
	n.watch()
	err = n.register(lis.Addr().String())
	if err != nil {
		return fmt.Errorf("register: %v", err)
	}
	return n.server.Serve(lis)
}

func (n *OracleNode) initConnections() error {
	log.Info("Initialize connections to other nodes")
	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}
	n.nodes = nodes
	for i := 0; i < len(nodes); i++ {
		if nodes[i].Addr == n.account {
			continue
		}
		_, err := n.connectionManager.NewConnection(nodes[i])
		if err != nil {
			log.Errorf("new connection %s: %v", nodes[i].IpAddr, err)
			continue
		}
		log.Infof("New connection to %s", nodes[i].Addr.String())
	}
	return nil
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

func (n *OracleNode) broadCastDeals(deals map[int]*dkg.Deal) {
	log.Infof("Broadcasting deals")
	for i, deal := range deals {
		conn, err := n.connectionManager.FindByAddress(n.nodes[i].Addr)
		if err != nil {
			log.Errorf("find connection by address: %v", err)
			continue
		}
		client := NewOracleNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		request := &ProcessDealRequest{
			Deal: dealToPb(deal),
		}
		_, err = client.ProcessDeal(ctx, request)
		if err != nil {
			log.Errorf("process deal: %v", err)
		}
		cancel()
	}
}

func (n *OracleNode) broadCastResponse(response *dkg.Response) {
	log.Infof("Broadcasting response with dealer %d and verifier %d", response.Index, response.Response.Index)
	for _, otherNode := range n.nodes {
		if otherNode.Addr == n.account {
			continue
		}
		conn, err := n.connectionManager.FindByAddress(otherNode.Addr)
		if err != nil {
			log.Errorf("find connection by address: %v", err)
			continue
		}
		client := NewOracleNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		request := &ProcessResponseRequest{
			Response: responseToPb(response),
		}
		go func() {
			_, err := client.ProcessResponse(ctx, request)
			if err != nil {
				log.Errorf("process response: %v", err)
			}
			cancel()
		}()
	}
}

func (n *OracleNode) DistKeyShare() (*dkg.DistKeyShare, error) {
	if n.distKey == nil {
		return nil, errors.New("no dist key share")
	}
	return n.distKey, nil
}

func (n *OracleNode) Stop() {
	n.connectionManager.Close()
	n.server.Stop()
}

func (n *OracleNode) GracefulStop() {
	n.connectionManager.Close()
	n.server.GracefulStop()
}

package iop

import (
	"context"
	"crypto/ecdsa"
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
	aggregator        *Aggregator
	ethClient         *ethclient.Client
	oracleContract    *OracleContract
	registryContract  *RegistryContractWrapper
	raffleContract    *RaffleContract
	distKeyContract   *DistKeyContract
	dkg               *dkg.DistKeyGenerator
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	suiteG2           *bn256.Suite
}

func NewOracleNode(
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	validator *Validator,
	aggregator *Aggregator,
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
		aggregator:        aggregator,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		raffleContract:    raffleContract,
		distKeyContract:   distKeyContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		blsPrivateKey:     blsPrivateKey,
		account:           account,
		suiteG2:           suite,
	}
	RegisterOracleNodeServer(grpcServer, node)
	return node
}

func (n *OracleNode) Serve(lis net.Listener) error {
	if err := n.initConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}
	n.watch(context.Background())
	if err := n.register(lis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}
	return n.server.Serve(lis)
}

func (n *OracleNode) initConnections() error {
	log.Info("Initialize connections to other nodes")
	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}
	for _, node := range nodes {
		if node.Addr == n.account {
			continue
		}
		if _, err := n.connectionManager.NewConnection(node); err != nil {
			log.Errorf("new connection %s: %v", node.IpAddr, err)
			continue
		}
		log.Infof("New connection to %s", node.Addr.String())
	}
	return nil
}

func (n *OracleNode) register(ipAddr string) error {
	isRegistered, err := n.registryContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	blsPublicKey := n.suiteG2.Point().Mul(n.blsPrivateKey, nil)
	b, err := blsPublicKey.MarshalBinary()
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

func (n *OracleNode) broadCastDeals(nodes []RegistryContractOracleNode, deals map[int]*dkg.Deal) {
	log.Infof("Broadcasting deals")
	for i, deal := range deals {
		conn, err := n.connectionManager.FindByAddress(nodes[i].Addr)
		if err != nil {
			log.Errorf("find connection by address: %v", err)
			continue
		}
		client := NewOracleNodeClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		request := &ProcessDealRequest{
			Deal: dealToPb(deal),
		}
		if _, err := client.ProcessDeal(ctx, request); err != nil {
			log.Errorf("process deal: %v", err)
		}
		cancel()
	}
}

func (n *OracleNode) broadCastResponse(response *dkg.Response) error {
	log.Infof("Broadcasting response with dealer %d and verifier %d", response.Index, response.Response.Index)
	nodes, err := n.registryContract.FindIopNodes()
	if err != nil {
		return fmt.Errorf("find nodes: %w", err)
	}

	for _, otherNode := range nodes {
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
			if _, err := client.ProcessResponse(ctx, request); err != nil {
				log.Errorf("process response: %v", err)
			}
			cancel()
		}()
	}
	return nil
}

func (n *OracleNode) DistKeyShare() (*dkg.DistKeyShare, error) {
	return n.dkg.DistKeyShare()
}

func (n *OracleNode) Stop() {
	n.server.Stop()
	n.ethClient.Close()
	n.connectionManager.Close()
}

func (n *OracleNode) GracefulStop() {
	n.server.GracefulStop()
	n.ethClient.Close()
	n.connectionManager.Close()
}

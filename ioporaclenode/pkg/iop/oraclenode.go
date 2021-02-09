package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	iota "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/trinary"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"google.golang.org/grpc"
	"net"
)

type OracleNode struct {
	UnimplementedOracleNodeServer
	dkg               *DistKeyGenerator
	ethClient         *ethclient.Client
	server            *grpc.Server
	connectionManager *ConnectionManager
	validator         *Validator
	aggregator        *Aggregator
	iotaClient        *iota.API
	oracleContract    *OracleContract
	registryContract  *RegistryContractWrapper
	distKeyContract   *DistKeyContract
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	seed              trinary.Trytes
	suite             suites.Suite
}

func NewOracleNode(
	dkg *DistKeyGenerator,
	ethClient *ethclient.Client,
	iotaClient *iota.API,
	connectionManager *ConnectionManager,
	validator *Validator,
	aggregator *Aggregator,
	oracleContract *OracleContract,
	registryContract *RegistryContractWrapper,
	distKeyContract *DistKeyContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	blsPrivateKey kyber.Scalar,
	account common.Address,
	seed trinary.Trytes,
	suite suites.Suite,
) *OracleNode {
	grpcServer := grpc.NewServer()
	node := &OracleNode{
		dkg:               dkg,
		server:            grpcServer,
		ethClient:         ethClient,
		iotaClient:        iotaClient,
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
		oracleContract:    oracleContract,
		registryContract:  registryContract,
		distKeyContract:   distKeyContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		blsPrivateKey:     blsPrivateKey,
		account:           account,
		seed:              seed,
		suite:             suite,
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
	nodes, err := n.registryContract.FindOracleNodes()
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

	blsPublicKey := n.suite.Point().Mul(n.blsPrivateKey, nil)
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

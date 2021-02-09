package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	oracleContract    *OracleContract
	registryContract  *RegistryContractWrapper
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	seed              trinary.Trytes
	suite             suites.Suite
}

func NewOracleNode(
	dkg *DistKeyGenerator,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	validator *Validator,
	aggregator *Aggregator,
	registryContract *RegistryContractWrapper,
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
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
		registryContract:  registryContract,
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

	go func() {
		err := n.watchAndHandleRegisterOracleNodeLog(context.Background())
		if err != nil {
			log.Errorf("watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		err := n.aggregator.WatchAndHandleValidateTransactionLog(context.Background())
		if err != nil {
			log.Errorf("watch and handle validate transaction log: %v", err)
		}
	}()

	go func() {
		err := n.dkg.ListenAndProcess(context.Background())
		if err != nil {
			log.Errorf("watch and handle dkg log: %v", err)
		}
	}()

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

func (n *OracleNode) watchAndHandleRegisterOracleNodeLog(ctx context.Context) error {
	sink := make(chan *RegistryContractRegisterOracleNodeLog)
	defer close(sink)

	sub, err := n.registryContract.WatchRegisterOracleNodeLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received register oracle node event %s", event.Sender.String())
			if err = n.handleRegisterOracleNodeLog(event); err != nil {
				log.Errorf("handle register oracle node log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *OracleNode) handleRegisterOracleNodeLog(event *RegistryContractRegisterOracleNodeLog) error {
	node, err := n.registryContract.FindOracleNodeByAddress(nil, event.Sender)
	if err != nil {
		return fmt.Errorf("find oracle node by address %s: %w", event.Sender, err)
	}
	if _, err := n.connectionManager.NewConnection(node); err != nil {
		return fmt.Errorf("new connection: %w", err)
	}
	log.Infof("New connection to %s", event.Sender.String())
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

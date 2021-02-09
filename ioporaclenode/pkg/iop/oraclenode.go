package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	iota "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/trinary"
	zmq "github.com/pebbe/zmq4"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"google.golang.org/grpc"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"net"
)

type OracleNode struct {
	UnimplementedOracleNodeServer
	server            *grpc.Server
	serverLis         net.Listener
	ethClient         *ethclient.Client
	iotaAPI           *iota.API
	zmqSocket         *zmq.Socket
	registryContract  *RegistryContractWrapper
	oracleContract    *OracleContract
	distKeyContract   *DistKeyContract
	suite             suites.Suite
	ecdsaPrivateKey   *ecdsa.PrivateKey
	blsPrivateKey     kyber.Scalar
	account           common.Address
	seed              trinary.Trytes
	dkg               *DistKeyGenerator
	connectionManager *ConnectionManager
	validator         *Validator
	aggregator        *Aggregator
}

func NewOracleNode(c Config) (*OracleNode, error) {
	server := grpc.NewServer()
	serverLis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		return nil, fmt.Errorf("listen on %s: %v", c.BindAddress, err)
	}

	ethClient, err := ethclient.Dial(c.Ethereum.Address)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	iotaAPI, err := iota.ComposeAPI(iota.HTTPClientSettings{URI: c.IOTA.Address})
	if err != nil {
		return nil, fmt.Errorf("iota client: %v", err)
	}

	zmqSocket, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, fmt.Errorf("zmq client: %v", err)
	}
	if err := zmqSocket.Connect(c.IOTA.Zmq); err != nil {
		return nil, fmt.Errorf("connect zmq: %v", err)
	}

	registryContract, err := NewRegistryContract(common.HexToAddress(c.Contracts.RegistryContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("registry contract: %v", err)
	}

	registryContractWrapper := &RegistryContractWrapper{
		RegistryContract: registryContract,
	}

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	distKeyContract, err := NewDistKeyContract(common.HexToAddress(c.Contracts.DistKeyContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("dist key contract: %v", err)
	}

	suite := bn256.NewSuiteG2()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}

	blsPrivateKey, err := HexToScalar(suite, c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to scalar: %v", err)
	}

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	connectionManager := NewConnectionManager()
	validator := NewValidator(suite, ethClient)
	aggregator := NewAggregator(
		suite,
		ethClient,
		connectionManager,
		oracleContract,
		registryContractWrapper,
		account,
		ecdsaPrivateKey,
	)
	dkg := NewDistKeyGenerator(
		suite,
		connectionManager,
		aggregator,
		zmqSocket,
		iotaAPI,
		registryContractWrapper,
		distKeyContract,
		ecdsaPrivateKey,
		blsPrivateKey,
		account,
		c.IOTA.Seed,
		c.IOTA.BroadcastAddress,
	)
	validator.SetDistKeyGenerator(dkg)
	aggregator.SetDistKeyGenerator(dkg)

	node := &OracleNode{
		server:            server,
		serverLis:         serverLis,
		ethClient:         ethClient,
		iotaAPI:           iotaAPI,
		zmqSocket:         zmqSocket,
		registryContract:  registryContractWrapper,
		oracleContract:    oracleContract,
		distKeyContract:   distKeyContract,
		suite:             suite,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		blsPrivateKey:     blsPrivateKey,
		account:           account,
		seed:              c.IOTA.Seed,
		dkg:               dkg,
		connectionManager: connectionManager,
		validator:         validator,
		aggregator:        aggregator,
	}
	RegisterOracleNodeServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {
	if err := n.initConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	go func() {
		if err := n.watchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleValidateTransactionLog(context.Background()); err != nil {
			log.Errorf("watch and handle validate transaction log: %v", err)
		}
	}()

	go func() {
		if err := n.dkg.ListenAndProcess(context.Background()); err != nil {
			log.Errorf("watch and handle dkg log: %v", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
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
	_ = n.zmqSocket.Close()
	n.ethClient.Close()
	n.connectionManager.Close()
}

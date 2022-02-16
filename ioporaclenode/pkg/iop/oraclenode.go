package iop

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iotaledger/iota.go/trinary"
	iota "github.com/iotaledger/iota.go/v2"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnimplementedOracleNodeServer
	server            *grpc.Server
	serverLis         net.Listener
	targetEthClient   *ethclient.Client
	sourceEthClient   *ethclient.Client
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

	targetEthClient, err := ethclient.Dial(c.Ethereum.TargetAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	sourceEthClient, err := ethclient.Dial(c.Ethereum.SourceAddress)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	iotaAPI := iota.NewNodeHTTPAPIClient("https://api.lb-0.h.chrysalis-devnet.iota.cafe")
	if err != nil {
		return nil, fmt.Errorf("iota client: %v", err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker("mqtt.lb-0.h.chrysalis-devnet.iota.cafe:1883")
	mqttClient := mqtt.NewClient(opts)

	registryContract, err := NewRegistryContract(common.HexToAddress(c.Contracts.RegistryContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("registry contract: %v", err)
	}

	registryContractWrapper := &RegistryContractWrapper{
		RegistryContract: registryContract,
	}

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), targetEthClient)
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	distKeyContract, err := NewDistKeyContract(common.HexToAddress(c.Contracts.DistKeyContractAddress), targetEthClient)
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

	connectionManager := NewConnectionManager(registryContractWrapper, account)
	validator := NewValidator(suite, oracleContract, sourceEthClient)
	aggregator := NewAggregator(
		suite,
		targetEthClient,
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
		mqttClient,
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
		targetEthClient:   targetEthClient,
		sourceEthClient:   sourceEthClient,
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
	if err := n.connectionManager.InitConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	go func() {
		if err := n.dkg.ListenAndProcess(context.Background()); err != nil {
			log.Errorf("watch and handle dkg log: %v", err)
		}
	}()

	go func() {
		if err := n.connectionManager.WatchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		if err := n.aggregator.WatchAndHandleValidateTransactionLog(context.Background()); err != nil {
			log.Errorf("watch and handle validate transaction log: %v", err)
		}
	}()

	go func() {
		if err := n.validator.WatchAndHandleValidateTransactionLog(context.Background()); err != nil {
			log.Errorf("watch and handle validate transaction log: %v", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
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

	minStake, err := n.registryContract.MINSTAKE(nil)
	if err != nil {
		return fmt.Errorf("min stake: %v", err)
	}

	auth := bind.NewKeyedTransactor(n.ecdsaPrivateKey)
	auth.Value = minStake

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
	n.targetEthClient.Close()
	n.sourceEthClient.Close()
	n.connectionManager.Close()
}

package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"ioporaclenode/internal/pkg/kyber/pairing/bn256"
	"ioporaclenode/pkg/iop"
	"net"
	"os"
	"os/signal"
)

var (
	addrFlag             = flag.String("address", "127.0.0.1:25565", "server address")
	ethFlag              = flag.String("eth", "ws://127.0.0.1:7545", "eth node address")
	oracleContractFlag   = flag.String("oracleContract", "0x50878A73Ed933d206317a2992A5052F702d0f112", "oracle contract address")
	registryContractFlag = flag.String("registryContract", "0xD1F96e353655eB1D2Bf7B13C83C034405d2CF2AF", "registry contract address")
	raffleContractFlag   = flag.String("raffleContract", "0xb4ee1244Ca06eD5F0ff150e1b0B0644fF42af3E8", "raffle contract address")
	distKeyContractFlag  = flag.String("distKeyContract", "0x7D92382eB89eD906D03fE75d9dCBA643822992ED", "dist key contract address")
	ecdsaPrivateKeyFlag  = flag.String("ecdsaPrivateKey", "0xe63ff25be694842b3d25f3c8981dbe44b36b23a6effdbe04f9ee11e7965c922b", "private key")
	blsPrivateKeyFlag    = flag.String("blsPrivateKey", "0x2e931ebbc908ec1993a789166f5690ee2ea34830df69a0fd0fc6a456b4aa8a46", "value of the private share")
)

func main() {
	flag.Parse()
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	registryContract, err := iop.NewRegistryContract(common.HexToAddress(*registryContractFlag), ethClient)
	if err != nil {
		log.Fatalf("registry contract: %v", err)
	}

	registryContractWrapper := &iop.RegistryContractWrapper{
		RegistryContract: registryContract,
	}

	oracleContract, err := iop.NewOracleContract(common.HexToAddress(*oracleContractFlag), ethClient)
	if err != nil {
		log.Fatalf("oracle contract: %v", err)
	}

	raffleContract, err := iop.NewRaffleContract(common.HexToAddress(*raffleContractFlag), ethClient)
	if err != nil {
		log.Fatalf("raffle contract: %v", err)
	}

	distKeyContract, err := iop.NewDistKeyContract(common.HexToAddress(*distKeyContractFlag), ethClient)
	if err != nil {
		log.Fatalf("dist key contract: %v", err)
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(*ecdsaPrivateKeyFlag)
	if err != nil {
		log.Fatalf("hex to ecdsa: %v", err)
	}

	suite := bn256.NewSuiteG2()
	blsPrivateKey, err := iop.HexToScalar(suite, *blsPrivateKeyFlag)
	if err != nil {
		log.Fatalf("hex to scalar: %v", err)
	}

	hexAddress, err := iop.AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		log.Fatalf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	connectionManager := iop.NewConnectionManager()
	validator := iop.NewValidator(suite, ethClient)
	aggregator := iop.NewAggregator(suite, ethClient, connectionManager, registryContractWrapper)
	node := iop.NewOracleNode(
		ethClient,
		connectionManager,
		validator,
		aggregator,
		oracleContract,
		registryContractWrapper,
		raffleContract,
		distKeyContract,
		ecdsaPrivateKey,
		blsPrivateKey,
		account,
		suite,
	)
	validator.SetNode(node)
	aggregator.SetNode(node)

	go func() {
		if err := node.Serve(lis); err != nil {
			log.Fatalf("serve %s: %v", lis.Addr(), err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	node.GracefulStop()
}

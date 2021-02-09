package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	iota "github.com/iotaledger/iota.go/api"
	zmq "github.com/pebbe/zmq4"
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
	iotaFlag             = flag.String("iota", "https://nodes.devnet.iota.org:443", "iota node address")
	zmqFlag              = flag.String("zmq", "tcp://127.0.0.1:5556", "zmq address")
	oracleContractFlag   = flag.String("oracleContract", "0x5D02e5caA40735c22206B6f2bc4509Ce8461cC30", "oracle contract address")
	registryContractFlag = flag.String("registryContract", "0xB16A8FebD068e5e9BFdB34d1123C50EEb9B59114", "registry contract address")
	distKeyContractFlag  = flag.String("distKeyContract", "0x23bEA5A93ca502c3A5a92f5E81496c5e9ce4A4a8", "dist key contract address")
	ecdsaPrivateKeyFlag  = flag.String("ecdsaPrivateKey", "0xe63ff25be694842b3d25f3c8981dbe44b36b23a6effdbe04f9ee11e7965c922b", "private key")
	blsPrivateKeyFlag    = flag.String("blsPrivateKey", "0x2e931ebbc908ec1993a789166f5690ee2ea34830df69a0fd0fc6a456b4aa8a46", "value of the private share")
	seedFlag             = flag.String("seed", "KQXMQNRFGVQECXKURNQUYZKDLZLNWJEMABXSHMREMSGRLDUHCCAWPFFZIQFCAWK9ZVQXJEYLINJ9WRCNZ", "iota seed")
	distKeyAddress       = flag.String("distKeyAddress", "DIIBHRZJLNIYTGFRCTHRYOHOJMJIGB9JXLFCUZLXDUCLOEPOEPRSHQDYZKEFDJPCOSQ9VMDSAMYMZPIDDAS9SKHSFX", "dist key broadcast channel")
)

func main() {
	flag.Parse()
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	iotaClient, err := iota.ComposeAPI(iota.HTTPClientSettings{URI: *iotaFlag})
	if err != nil {
		log.Fatalf("iota client: %v", err)
	}

	zmqClient, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatalf("zmq client: %v", err)
	}
	if err := zmqClient.Connect(*zmqFlag); err != nil {
		log.Fatalf("connect zmq: %v", err)
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
	aggregator := iop.NewAggregator(
		suite,
		ethClient,
		connectionManager,
		oracleContract,
		registryContractWrapper,
		account,
		ecdsaPrivateKey,
	)
	dkg := iop.NewDistKeyGenerator(
		suite,
		connectionManager,
		aggregator,
		zmqClient,
		iotaClient,
		registryContractWrapper,
		distKeyContract,
		ecdsaPrivateKey,
		blsPrivateKey,
		account,
		*seedFlag,
		*distKeyAddress,
	)
	validator.SetDistKeyGenerator(dkg)
	aggregator.SetDistKeyGenerator(dkg)

	node := iop.NewOracleNode(
		dkg,
		ethClient,
		connectionManager,
		validator,
		aggregator,
		registryContractWrapper,
		ecdsaPrivateKey,
		blsPrivateKey,
		account,
		*seedFlag,
		suite,
	)

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

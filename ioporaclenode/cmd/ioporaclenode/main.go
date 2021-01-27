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
	oracleContractFlag   = flag.String("oracleContract", "0x32742536554092C1f8d47A1E71aBd1daaEF55aBB", "oracle contract address")
	registryContractFlag = flag.String("registryContract", "0x06983D416C40273bc22a9492F3a0D22AFAEdC26F", "registry contract address")
	raffleContractFlag   = flag.String("raffleContract", "0x13aEeF0aa983851296707A37C5529C83FA3f2ee5", "raffle contract address")
	distKeyContractFlag  = flag.String("distKeyContract", "0x998BdE7FF93774DF2087e23Df1d1d37Ff0977d5C", "dist key contract address")
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

	suite := bn256.NewSuite()
	blsPrivateKey, err := iop.HexToScalar(suite.G2(), *blsPrivateKeyFlag)

	hexAddress, err := iop.AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		log.Fatalf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	validator := iop.NewValidator(suite, nil, ethClient)
	node := iop.NewOracleNode(
		ethClient,
		validator,
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

	go func() {
		if err := node.Serve(lis); err != nil {
			log.Fatalf("serve %s: %v", lis.Addr(), err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	node.GracefulStop()
	ethClient.Close()
}

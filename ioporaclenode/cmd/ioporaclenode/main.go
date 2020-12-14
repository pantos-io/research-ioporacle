package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"ioporaclenode/pkg/iop"
	"net"
	"os"
	"os/signal"
)

var (
	addrFlag             = flag.String("address", "127.0.0.1:25565", "server address")
	ethFlag              = flag.String("eth", "ws://127.0.0.1:7545", "eth node address")
	oracleContractFlag   = flag.String("oracleContract", "0x45489337BFcdc944667b24ef8F4F252Dd206c27A", "oracle contract address")
	registryContractFlag = flag.String("registryContract", "0xa409953b6Fc60B75ab31e3617D8eeFE1B3B5e55c", "registry contract address")
	keyFlag              = flag.String("key", "e63ff25be694842b3d25f3c8981dbe44b36b23a6effdbe04f9ee11e7965c922b", "private key")
)

func main() {
	flag.Parse()
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	oracleContract, err := iop.NewOracleContract(common.HexToAddress(*oracleContractFlag), ethClient)
	if err != nil {
		log.Fatalf("oracle contract: %v", err)
	}

	registryContract, err := iop.NewRegistryContract(common.HexToAddress(*registryContractFlag), ethClient)
	if err != nil {
		log.Fatalf("registry contract: %v", err)
	}

	key, err := crypto.HexToECDSA(*keyFlag)
	if err != nil {
		log.Fatalf("hex to ecdsa: %v", err)
	}

	hexAddress, err := iop.AddressFromPrivateKey(key)
	if err != nil {
		log.Fatalf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	aggregator := iop.NewAggregator(ethClient, registryContract, account)
	txVerifier := iop.NewTransactionVerifier(ethClient)
	oracleNode := iop.NewOracleNode(ethClient, txVerifier, aggregator, oracleContract, registryContract, key, account)

	go func() {
		if err := oracleNode.Serve(lis); err != nil {
			log.Fatalf("serve %s: %v", lis.Addr(), err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	oracleNode.GracefulStop()
	ethClient.Close()
}

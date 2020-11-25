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
	addrFlag     = flag.String("address", "127.0.0.1:25565", "server address")
	ethFlag      = flag.String("eth", "ws://127.0.0.1:7545", "eth node address")
	contractFlag = flag.String("contract", "0x5ea54a26420904c558647ec5fF5bA5424e6367f0", "contract address")
	keyFlag      = flag.String("key", "e63ff25be694842b3d25f3c8981dbe44b36b23a6effdbe04f9ee11e7965c922b", "private key")
)

func main() {
	flag.Parse()
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	contract, err := iop.NewOracleContract(common.HexToAddress(*contractFlag), ethClient)
	key, err := crypto.HexToECDSA(*keyFlag)
	if err != nil {
		log.Fatalf("hex to ecdsa: %v", err)
	}

	hexAddress, err := iop.AddressFromPrivateKey(key)
	if err != nil {
		log.Fatalf("address from private key: %w", err)
	}
	account := common.HexToAddress(hexAddress)

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	txVerifier := iop.NewTransactionVerifier(ethClient, contract, account)
	oracleNode := iop.NewOracleNode(txVerifier, contract, key, account)

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

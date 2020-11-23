package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"net"
	"oraclenode/pkg/iop"
	"os"
	"os/signal"
)

var (
	addrFlag     = flag.String("address", "127.0.0.1:25565", "server address")
	ethFlag      = flag.String("eth", "ws://127.0.0.1:7545", "eth node address")
	contractFlag = flag.String("contract", "0x2f441c10dB56A8dbfe4E9602C38A658f9Fb8aAA5", "contract address")
	keyFlag      = flag.String("key", "405760727b443e63701fd8d9d95f2834d114cc90196efde6acda60ecae09d975", "private key")
)

func main() {
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	contract, err := iop.NewIopOracleContract(common.HexToAddress(*contractFlag), ethClient)

	key, err := crypto.HexToECDSA(*keyFlag)
	if err != nil {
		log.Fatalf("hex to ecdsa: %v", err)
	}

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	txVerifier := iop.NewTransactionVerifier(ethClient, contract)
	oracleNode, err := iop.NewOracleNode(txVerifier, contract, key)
	if err != nil {
		log.Fatalf("create oracle node: %v", err)
	}

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

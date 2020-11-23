package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"net"
	"oraclenode/pkg/iop"
	"os"
	"os/signal"
)

var (
	addrFlag = flag.String("address", "0.0.0.0:25565", "server address")
	ethFlag  = flag.String("eth", "ws://127.0.0.1:7545", "eth node address")
)

func main() {
	ethClient, err := ethclient.Dial(*ethFlag)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatalf("listen on %s: %v", *addrFlag, err)
	}

	txVerifier := iop.NewTransactionVerifier(ethClient)
	oracleNode := iop.NewOracleNode(txVerifier)
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

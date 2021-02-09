package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"ioporaclenode/pkg/iop"
	"os"
	"os/signal"
)

func main() {
	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config: %v", err)
	}

	var config iop.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unmarshal config into struct, %v", err)
	}
	log.Printf("%v", config)

	log.Infof("Loaded config file %s", *configFile)

	node, err := iop.NewOracleNode(config)
	if err != nil {
		log.Fatalf("new oracle node: %v", err)
	}

	go func() {
		if err := node.Run(); err != nil {
			log.Fatalf("run node: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	node.Stop()
}

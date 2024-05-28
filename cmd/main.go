package main

import (
	"L0/internal/config"
	natsConnector "L0/internal/nats_connector"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = natsConnector.Run( //to add connection
		cfg.ClusterID,
		cfg.ClientID,
		cfg.ChannelName,
		cfg.URL,
		MessageHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	//err = connection.Close()
	//if err != nil {
	//	log.Print(err)
	//}
	<-sigCh

	// TODO: init connection to nats-streaming nats_connector

	// TODO init storage

	// TODO run server
}

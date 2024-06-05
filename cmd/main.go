package main

import (
	"L0/internal/config"
	"L0/internal/http_server"
	"L0/internal/storage"
	"context"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	cacheCapacity       = 1000
	numberOfSubscribers = 10
	numberOfWorkers     = 10
	messagesBufferSize  = 100
)

func main() {
	log.Println("Starting server...")

	//init config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// init storage (cache and db)
	strg, err := storage.Load(
		cfg.DB.UserName,
		cfg.DB.Password,
		cfg.DB.Address,
		cfg.DB.Database,
		cacheCapacity,
	)
	if err != nil {
		log.Fatal(err)
	}

	// connecting to nats-streaming server
	log.Println("Connecting to nats-streaming server...")
	connection, err := stan.Connect(
		cfg.NatsStreaming.ClusterID,
		cfg.NatsStreaming.ClientID,
		stan.NatsURL(cfg.NatsStreaming.URL),
	)
	if err != nil {
		log.Fatal(err)
	}

	messages := make(chan []byte, messagesBufferSize)

	var wg sync.WaitGroup

	// creating subscribers of the target channel
	for i := 0; i < numberOfSubscribers; i++ {
		wg.Add(1)
		go NewSubscriber(
			connection,
			cfg.NatsStreaming.ChannelName,
			cfg.NatsStreaming.QueueGroup,
			&wg,
			messages,
		)
	}
	// creating workers, that will be handling messages
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go NewWorker(
			messages,
			&wg,
			MessageHandler,
			strg,
		)
	}

	// launching server
	httpServer := http_server.Start(
		cfg.HTTPServer.Address,
		cfg.HTTPServer.Timeout,
		cfg.HTTPServer.TimeoutIdle,
		strg,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	close(messages)
	wg.Wait()
	log.Println("Shutting down http server...")
	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("Disconnecting nats-streaming server...")
	err = connection.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println("Disconnecting database...")
	err = strg.StorageDB.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println("Exiting server...")
}

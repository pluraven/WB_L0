package main

import (
	"L0/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("config: %+v\n", cfg)

	// TODO: init connection to nats-streaming nats-channel

	// TODO init storage

	// TODO run server

}

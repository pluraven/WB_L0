package main

import (
	"L0/internal/storage"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
)

func NewSubscriber(
	connection stan.Conn,
	channelName string,
	queueGroup string,
	wg *sync.WaitGroup,
	messages chan<- []byte) {
	defer wg.Done()
	_, err := connection.QueueSubscribe(
		channelName,
		queueGroup,
		func(msg *stan.Msg) {
			messages <- msg.Data
		},
		stan.DurableName(queueGroup))
	if err != nil {
		log.Fatal("Error while subscribing: ", err)
	}
}

type messageHandler func([]byte, *storage.Storage)

func NewWorker(messages <-chan []byte, wg *sync.WaitGroup, handler messageHandler, strg *storage.Storage) {
	defer wg.Done()
	for msg := range messages {
		handler(msg, strg)
	}
}

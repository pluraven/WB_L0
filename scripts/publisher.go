package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "sender"
	channel   = "test-channel"
)

func main() {
	connection, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	var choice int = 1
	for {
		fmt.Print("0. Exit\n1. Publish new message\nYour choice: ")
		fmt.Scan(&choice)
		if choice == 0 {
			break
		}
		orderUID := faker.UUIDDigit()
		transaction := faker.UUIDDigit()
		chrtID := rand.Intn(100000)
		price := rand.Intn(100000)
		date := time.Now().Format("2006-01-02T15:04:05Z")
		fmt.Println("Message:")
		message := fmt.Sprintf(
			`{
		"order_uid": "%s",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "%s",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": %d,
			"track_number": "WBILMTESTTRACK",
			"price": %d,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "%s",
		"oof_shard": 1
	  }`,
			orderUID,
			transaction,
			chrtID,
			price,
			date,
		)
		fmt.Println(message)
		err = connection.Publish(channel, []byte(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Message sent successfully.")
		fmt.Printf("OrderUID:\n%s\n", orderUID)
	}
}

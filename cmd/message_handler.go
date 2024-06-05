package main

import (
	"L0/internal/storage"
	"encoding/json"
	"log"
	"time"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDT    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type CorrectMessage struct {
	OrderUID string `json:"order_uid"`
	TrackNum string `json:"track_num"`
	Entry    string `json:"entry"`
	Delivery
	Payment
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OOFShard          int       `json:"oof_shard"`
}

func MessageHandler(message []byte, strg *storage.Storage) {
	log.Printf("Recived message\n")
	var correctMessage CorrectMessage
	err := json.Unmarshal(message, &correctMessage)
	if err != nil {
		log.Printf("Invalid message: %v\n", err)
	} else {
		log.Printf("Message is valid\nWriting message to storage...\n")
		orderToSave := &storage.JsonData{ID: correctMessage.OrderUID, Data: string(message), DataString: string(message)}
		err = strg.WriteToDB(orderToSave)
		if err != nil {
			log.Printf("Error writting message to DB: %v\n", err)
		}
		strg.PutInCache(orderToSave.ID, orderToSave.Data)
	}
}

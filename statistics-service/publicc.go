package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	// Пример события inventory.created
	event := map[string]string{
		"product_id": "123",
		"name":       "Test Product",
		"user_id":    "user1",
	}
	data, _ := json.Marshal(event)

	if err := nc.Publish("inventory.created", data); err != nil {
		log.Fatal(err)
	}

	log.Println("Sent inventory.created event")
}

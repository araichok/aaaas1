package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	userID := "test-user"

	// Тест: inventory.created
	invEvent := map[string]string{
		"product_id": "p1",
		"name":       "Test Product",
		"user_id":    userID,
	}
	send(nc, "inventory.created", invEvent)

	// Тест: order.created
	orderEvent := map[string]string{
		"order_id": "o1",
		"user_id":  userID,
		"time":     time.Now().Format(time.RFC3339),
	}
	send(nc, "order.created", orderEvent)

	log.Println("Events sent.")
}

func send(nc *nats.Conn, subject string, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}
	if err := nc.Publish(subject, b); err != nil {
		log.Println("Publish error:", err)
	}
}

package events

import (
	"encoding/json"
	"log"
	"statistics-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

type InventoryEvent struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
}

type OrderEvent struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
	Time    string `json:"time"`
}

func SubscribeInventoryEvents(nc *nats.Conn, useCase *usecase.StatisticsUseCase) {

	nc.Subscribe("inventory.created", func(msg *nats.Msg) {
		var event InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("[ERROR] decoding inventory.created event:", err)
			return
		}
		log.Printf("[NATS] Received inventory.created: %+v\n", event)
		useCase.IncrementInventoryCreated()
	})

	nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		var event InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("[ERROR] decoding inventory.updated event:", err)
			return
		}
		log.Printf("[NATS] Received inventory.updated: %+v\n", event)
		useCase.IncrementInventoryUpdated()
	})

	nc.Subscribe("inventory.deleted", func(msg *nats.Msg) {
		var event InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("[ERROR] decoding inventory.deleted event:", err)
			return
		}
		log.Printf("[NATS] Received inventory.deleted: %+v\n", event)
		useCase.IncrementInventoryDeleted()
	})
}

func SubscribeOrderEvents(nc *nats.Conn, useCase *usecase.StatisticsUseCase) {
	nc.Subscribe("order.created", func(msg *nats.Msg) {
		log.Println("[NATS] Received order.created")

		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding order.created event:", err)
			return
		}

		log.Printf("Parsed order: %+v\n", event)
		useCase.IncrementOrderCreated(event.UserID, event.Time)
	})

	nc.Subscribe("order.updated", func(msg *nats.Msg) {
		log.Println("[NATS] Received order.updated")

		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding order.updated event:", err)
			return
		}

		log.Printf("Parsed order (updated): %+v\n", event)
		useCase.IncrementOrderUpdated(event.UserID)
	})

}

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
	UserID    string `json:"user_id"`
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
			log.Println("Error decoding inventory.created event:", err)
			return
		}
		useCase.IncrementInventoryCreated(event.UserID)
	})

	nc.Subscribe("inventory.updated", func(msg *nats.Msg) {
		var event InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding inventory.updated event:", err)
			return
		}
		useCase.IncrementInventoryUpdated(event.UserID)
	})

	nc.Subscribe("inventory.deleted", func(msg *nats.Msg) {
		var event InventoryEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding inventory.deleted event:", err)
			return
		}
		useCase.IncrementInventoryDeleted(event.UserID)
	})
}

func SubscribeOrderEvents(nc *nats.Conn, useCase *usecase.StatisticsUseCase) {
	nc.Subscribe("order.created", func(msg *nats.Msg) {
		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding order.created event:", err)
			return
		}
		useCase.IncrementOrderCreated(event.UserID, event.Time)
	})

	nc.Subscribe("order.updated", func(msg *nats.Msg) {
		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding order.updated event:", err)
			return
		}
		useCase.IncrementOrderUpdated(event.UserID)
	})

	nc.Subscribe("order.deleted", func(msg *nats.Msg) {
		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Println("Error decoding order.deleted event:", err)
			return
		}
		useCase.IncrementOrderDeleted(event.UserID)
	})
}

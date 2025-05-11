package events

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"inventory-service/internal/domain"
)

type EventPublisher struct {
	nc *nats.Conn
}

func NewEventPublisher(nc *nats.Conn) *EventPublisher {
	return &EventPublisher{nc: nc}
}

func (p *EventPublisher) PublishInventoryEvent(action string, product *domain.Product) {
	event := map[string]interface{}{
		"product_id": product.ID,
		"name":       product.Name,
		"stock":      product.Stock,
		"price":      product.Price,
		"category":   product.Category,
		"action":     action,
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling inventory event:", err)
		return
	}

	subject := "inventory." + action
	if err := p.nc.Publish(subject, data); err != nil {
		log.Println("Error publishing to NATS:", err)
	} else {
		log.Println("Published to NATS:", subject)
	}
}

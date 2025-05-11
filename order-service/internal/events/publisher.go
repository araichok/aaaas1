package events

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type EventPublisher struct {
	nc *nats.Conn
}

func NewEventPublisher(nc *nats.Conn) *EventPublisher {
	return &EventPublisher{nc: nc}
}

func (p *EventPublisher) PublishOrderEvent(action string, orderID string, userID string) {
	event := map[string]interface{}{
		"order_id": orderID,
		"user_id":  userID,
		"time":     time.Now().Format(time.RFC3339),
	}
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshaling order event:", err)
		return
	}

	subject := "order." + action
	if err := p.nc.Publish(subject, data); err != nil {
		log.Println("Error publishing to NATS:", err)
	} else {
		log.Println("[NATS] Published to:", subject)
	}
}

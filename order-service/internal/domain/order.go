package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusCompleted OrderStatus = "completed"
	StatusCancelled OrderStatus = "cancelled"
)

type ProductItem struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
}

type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Items     []ProductItem      `json:"items" bson:"items"`
	Status    OrderStatus        `json:"status" bson:"status"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}

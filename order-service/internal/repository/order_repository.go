package repository

import (
	"context"
	"order-service/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{collection: db.Collection("orders")}
}

func (r *OrderRepository) Create(order *domain.Order) error {
	order.CreatedAt = time.Now().Unix()
	_, err := r.collection.InsertOne(context.TODO(), order)
	return err
}

func (r *OrderRepository) GetByID(id string) (*domain.Order, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	var order domain.Order
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&order)
	return &order, err
}

func (r *OrderRepository) UpdateStatus(id string, status domain.OrderStatus) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"status": status}},
	)
	return err
}

func (r *OrderRepository) GetByUser(userID string) ([]domain.Order, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var orders []domain.Order
	for cursor.Next(context.TODO()) {
		var order domain.Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	return orders, nil
}

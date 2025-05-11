package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
	ordersCol  *mongo.Collection
}

func NewMongoRepository(uri string) *MongoRepository {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return &MongoRepository{
		client:     client,
		collection: client.Database("statistics").Collection("inventory_stats"),
		ordersCol:  client.Database("statistics").Collection("orders_stats"),
	}
}

// --- Inventory: Глобальная статистика (без user_id) ---

func (r *MongoRepository) SaveInventoryEvent() {
	_, _ = r.collection.UpdateOne(context.TODO(),
		bson.M{"type": "global"},
		bson.M{"$inc": bson.M{"created": 1}},
		options.Update().SetUpsert(true),
	)
}

func (r *MongoRepository) SaveInventoryUpdate() {
	_, _ = r.collection.UpdateOne(context.TODO(),
		bson.M{"type": "global"},
		bson.M{"$inc": bson.M{"updated": 1}},
		options.Update().SetUpsert(true),
	)
}

func (r *MongoRepository) SaveInventoryDelete() {
	_, _ = r.collection.UpdateOne(context.TODO(),
		bson.M{"type": "global"},
		bson.M{"$inc": bson.M{"deleted": 1}},
		options.Update().SetUpsert(true),
	)
}

func (r *MongoRepository) GetInventoryCount() int32 {
	var result struct {
		Created int32 `bson:"created"`
	}
	_ = r.collection.FindOne(context.TODO(), bson.M{"type": "global"}).Decode(&result)
	return result.Created
}

// --- Orders: Статистика по user_id ---

func (r *MongoRepository) SaveOrderCreated(userID string, timeStr string) {
	log.Println("[MONGO] SaveOrderCreated for:", userID)

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Println("Error parsing time:", err)
		return
	}
	hour := parsedTime.Hour()

	_, err = r.ordersCol.UpdateOne(context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$inc": bson.M{
			"created":                             1,
			"hourly_orders." + strconv.Itoa(hour): 1,
		}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Println("Mongo update error:", err)
	}
}

func (r *MongoRepository) SaveOrderUpdated(userID string) {
	_, _ = r.ordersCol.UpdateOne(context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$inc": bson.M{"updated": 1}},
		options.Update().SetUpsert(true),
	)
}

func (r *MongoRepository) SaveOrderDeleted(userID string) {
	_, _ = r.ordersCol.UpdateOne(context.TODO(),
		bson.M{"user_id": userID},
		bson.M{"$inc": bson.M{"deleted": 1}},
		options.Update().SetUpsert(true),
	)
}

func (r *MongoRepository) GetOrderStats(userID string) (created, updated, deleted int32, hourlyOrders map[int]int32) {
	var result struct {
		Created int32         `bson:"created"`
		Updated int32         `bson:"updated"`
		Deleted int32         `bson:"deleted"`
		Hourly  map[int]int32 `bson:"hourly_orders"`
	}
	err := r.ordersCol.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&result)
	if err != nil {
		log.Println("Error fetching order stats:", err)
		return 0, 0, 0, make(map[int]int32)
	}

	return result.Created, result.Updated, result.Deleted, result.Hourly
}

package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/internal/domain"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id string) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id string) error
	List() ([]domain.Product, error)
}

type mongoProductRepo struct {
	collection *mongo.Collection
}

func NewMongoProductRepo(db *mongo.Database) ProductRepository {
	return &mongoProductRepo{
		collection: db.Collection("products"),
	}
}

func (r *mongoProductRepo) Create(p *domain.Product) error {
	p.ID = primitive.NewObjectID().Hex()
	_, err := r.collection.InsertOne(context.TODO(), p)
	return err
}

func (r *mongoProductRepo) GetByID(id string) (*domain.Product, error) {
	var product domain.Product
	err := r.collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

func (r *mongoProductRepo) Update(p *domain.Product) error {
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"id": p.ID}, bson.M{"$set": p})
	return err
}

func (r *mongoProductRepo) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (r *mongoProductRepo) List() ([]domain.Product, error) {
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	if err := cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

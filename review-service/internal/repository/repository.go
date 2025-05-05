package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"review-service/internal/domain"
)

type ReviewRepository interface {
	Save(ctx context.Context, review *domain.Review) (*domain.Review, error)
	Update(ctx context.Context, review *domain.Review) (*domain.Review, error)
	FindByID(ctx context.Context, id uint64) (*domain.Review, error)
}

type MongoReviewRepository struct {
	collection *mongo.Collection
}

func NewMongoReviewRepository(client *mongo.Client, dbName, Collection string) *MongoReviewRepository {
	collection := client.Database(dbName).Collection(Collection)
	return &MongoReviewRepository{collection: collection}
}

func (r *MongoReviewRepository) Save(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	_, err := r.collection.InsertOne(ctx, review)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *MongoReviewRepository) Update(ctx context.Context, review *domain.Review) (*domain.Review, error) {
	filter := bson.M{"id": review.ID}
	update := bson.M{
		"$set": bson.M{
			"productid": review.ProductID,
			"userid":    review.UserID,
			"rating":    review.Rating,
			"comment":   review.Comment,
			"updatedat": review.UpdatedAt,
		},
	}
	result := r.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("review not found")
		}
		return nil, result.Err()
	}

	var updatedReview domain.Review
	if err := result.Decode(&updatedReview); err != nil {
		return nil, err
	}
	return &updatedReview, nil
}

func (r *MongoReviewRepository) FindByID(ctx context.Context, id uint64) (*domain.Review, error) {
	filter := bson.M{"id": id}
	var review domain.Review
	err := r.collection.FindOne(ctx, filter).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("review not found")
		}
		return nil, err
	}
	return &review, nil
}

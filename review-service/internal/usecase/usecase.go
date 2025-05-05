package usecase

import (
	"context"
	"time"

	"review-service/internal/domain"
	"review-service/internal/repository"
)

type ReviewUsecase interface {
	CreateReview(ctx context.Context, productID, userID uint64, rating float64, comment string) (*domain.Review, error)
	UpdateReview(ctx context.Context, id, productID, userID uint64, rating float64, comment string) (*domain.Review, error)
	GetReview(ctx context.Context, id uint64) (*domain.Review, error)
}

type reviewUsecase struct {
	repo repository.ReviewRepository
}

func NewReviewUsecase(r repository.ReviewRepository) ReviewUsecase {
	return &reviewUsecase{repo: r}
}

func (u *reviewUsecase) CreateReview(ctx context.Context, productID, userID uint64, rating float64, comment string) (*domain.Review, error) {
	review := &domain.Review{
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.repo.Save(ctx, review)
}

func (u *reviewUsecase) UpdateReview(ctx context.Context, id, productID, userID uint64, rating float64, comment string) (*domain.Review, error) {
	review := &domain.Review{
		ID:        id,
		ProductID: productID,
		UserID:    userID,
		Rating:    rating,
		Comment:   comment,
		UpdatedAt: time.Now(),
	}

	return u.repo.Update(ctx, review)
}

func (u *reviewUsecase) GetReview(ctx context.Context, id uint64) (*domain.Review, error) {
	return u.repo.FindByID(ctx, id)
}

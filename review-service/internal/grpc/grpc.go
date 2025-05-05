package grpc

import (
	"context"

	"review-service/internal/domain"
	"review-service/internal/usecase"
	"review-service/review-service/proto/reviewpb"
)

type ReviewServer struct {
	reviewpb.UnimplementedReviewServiceServer
	reviewUsecase usecase.ReviewUsecase
}

func NewReviewServer(uc usecase.ReviewUsecase) *ReviewServer {
	return &ReviewServer{
		reviewUsecase: uc,
	}
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *reviewpb.CreateReviewRequest) (*reviewpb.CreateReviewResponse, error) {
	review, err := s.reviewUsecase.CreateReview(ctx, req.GetProductId(), req.GetUserId(), req.GetRating(), req.GetComment())
	if err != nil {
		return nil, err
	}

	return &reviewpb.CreateReviewResponse{
		Review: toProtoReview(review),
	}, nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *reviewpb.UpdateReviewRequest) (*reviewpb.UpdateReviewResponse, error) {
	review, err := s.reviewUsecase.UpdateReview(ctx, req.GetId(), req.GetProductId(), req.GetUserId(), req.GetRating(), req.GetComment())
	if err != nil {
		return nil, err
	}

	return &reviewpb.UpdateReviewResponse{
		Review: toProtoReview(review),
	}, nil
}

func (s *ReviewServer) GetReview(ctx context.Context, req *reviewpb.GetReviewRequest) (*reviewpb.GetReviewResponse, error) {
	review, err := s.reviewUsecase.GetReview(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &reviewpb.GetReviewResponse{
		Review: toProtoReview(review),
	}, nil
}

func toProtoReview(r *domain.Review) *reviewpb.Review {
	return &reviewpb.Review{
		Id:        r.ID,
		ProductId: r.ProductID,
		UserId:    r.UserID,
		Rating:    r.Rating,
		Comment:   r.Comment,
	}
}

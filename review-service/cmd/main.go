package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"review-service/database"
	qwerty "review-service/internal/grpc"
	"review-service/internal/repository"
	"review-service/internal/usecase"
	"review-service/review-service/proto/reviewpb"
)

func main() {

	db := database.InitMongo()
	reviewRepo := repository.NewMongoReviewRepository(db.Client(), "review_db", "reviews")

	reviewUsecase := usecase.NewReviewUsecase(reviewRepo)
	reviewServer := qwerty.NewReviewServer(reviewUsecase)

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()
	reviewpb.RegisterReviewServiceServer(grpcServer, reviewServer)

	fmt.Println("gRPC server listening on port 50053")

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}

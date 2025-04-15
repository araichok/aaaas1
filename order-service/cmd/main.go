package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"order-service/database"
	grpcHandler1 "order-service/internal/grpc"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/order-service/proto/orderpb"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	orderRepo := repository.NewOrderRepository(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	orderGRPCServer := grpcHandler1.NewOrderServer(orderUC)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, orderGRPCServer)

	log.Println("Order gRPC server listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

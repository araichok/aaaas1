package main

import (
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"order-service/database"
	"order-service/internal/events"
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

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	orderRepo := repository.NewOrderRepository(db)
	publisher := events.NewEventPublisher(nc)
	orderUC := usecase.NewOrderUsecase(orderRepo, publisher)
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

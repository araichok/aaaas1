package main

import (
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"inventory-service/database"
	"inventory-service/events"
	"inventory-service/internal/cache"
	grpcHandler "inventory-service/internal/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/inventory-service/proto/inventorypb"
)

func main() {
	db := database.InitMongo()
	repo := repository.NewMongoProductRepo(db)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	productCache := cache.NewProductCache()
	publisher := events.NewEventPublisher(nc)
	uc := usecase.NewProductUsecase(repo, productCache, publisher)
	grpcServer := grpcHandler.NewInventoryServer(uc)

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(s, grpcServer)

	log.Println("Inventory gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

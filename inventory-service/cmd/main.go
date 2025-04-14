package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"inventory-service/database"
	grpcHandler "inventory-service/internal/grpc"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/inventory-service/proto/inventorypb"
)

func main() {
	db := database.InitMongo()
	repo := repository.NewMongoProductRepo(db)
	uc := usecase.NewProductUsecase(repo)
	grpcServer := grpcHandler.NewInventoryServer(uc)

	lis, err := net.Listen("tcp", ":50051")
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

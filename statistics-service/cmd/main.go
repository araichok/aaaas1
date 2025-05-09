package main

import (
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"statistics-service/events"
	"statistics-service/internal/delivery"
	"statistics-service/internal/repository"
	"statistics-service/internal/usecase"
	pb "statistics-service/statistics-service/proto"

	"google.golang.org/grpc"
)

func main() {
	repo := repository.NewMongoRepository("mongodb://localhost:27017")

	useCase := usecase.NewStatisticsUseCase(repo)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	events.SubscribeInventoryEvents(nc, useCase)
	events.SubscribeOrderEvents(nc, useCase)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(s, delivery.NewStatisticsHandler(useCase))

	log.Println("Statistics service listening on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

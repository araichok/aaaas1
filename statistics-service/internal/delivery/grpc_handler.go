package delivery

import (
	"context"

	"statistics-service/internal/usecase"
	pb "statistics-service/statistics-service/proto"
)

type StatisticsHandler struct {
	uc *usecase.StatisticsUseCase
	pb.UnimplementedStatisticsServiceServer
}

func NewStatisticsHandler(uc *usecase.StatisticsUseCase) *StatisticsHandler {
	return &StatisticsHandler{uc: uc}
}

func (s *StatisticsHandler) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	count := s.uc.GetInventoryCount()
	return &pb.UserStatisticsResponse{TotalInventoryCreated: count}, nil
}

func (s *StatisticsHandler) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {
	created, updated, deleted, hourlyStats := s.uc.GetOrderStats(req.UserId)

	hourlyStatsConverted := make(map[int32]int32)
	for hour, count := range hourlyStats {
		hourlyStatsConverted[int32(hour)] = count
	}

	return &pb.UserOrderStatisticsResponse{
		TotalOrdersCreated: created,
		TotalOrdersUpdated: updated,
		TotalOrdersDeleted: deleted,
		OrdersHourlyStats:  hourlyStatsConverted,
	}, nil
}

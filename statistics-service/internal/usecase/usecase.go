package usecase

import (
	"statistics-service/internal/repository"
)

type StatisticsUseCase struct {
	repo *repository.MongoRepository
}

func NewStatisticsUseCase(r *repository.MongoRepository) *StatisticsUseCase {
	return &StatisticsUseCase{repo: r}
}

func (u *StatisticsUseCase) IncrementInventoryCreated(userID string) {
	u.repo.SaveInventoryEvent(userID)
}

func (u *StatisticsUseCase) IncrementInventoryUpdated(userID string) {
	u.repo.SaveInventoryUpdate(userID)
}

func (u *StatisticsUseCase) IncrementInventoryDeleted(userID string) {
	u.repo.SaveInventoryDelete(userID)
}

func (u *StatisticsUseCase) GetInventoryCount(userID string) int32 {
	return u.repo.GetInventoryCount(userID)
}

func (u *StatisticsUseCase) IncrementOrderCreated(userID string, timeStr string) {
	u.repo.SaveOrderCreated(userID, timeStr)
}

func (u *StatisticsUseCase) IncrementOrderUpdated(userID string) {
	u.repo.SaveOrderUpdated(userID)
}

func (u *StatisticsUseCase) IncrementOrderDeleted(userID string) {
	u.repo.SaveOrderDeleted(userID)
}

func (u *StatisticsUseCase) GetOrderStats(userID string) (created, updated, deleted int32, hourlyOrders map[int]int32) {

	return u.repo.GetOrderStats(userID)
}

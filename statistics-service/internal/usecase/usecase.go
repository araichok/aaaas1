package usecase

import (
	"log"
	"statistics-service/internal/repository"
)

type StatisticsUseCase struct {
	repo *repository.MongoRepository
}

func NewStatisticsUseCase(r *repository.MongoRepository) *StatisticsUseCase {
	return &StatisticsUseCase{repo: r}
}

func (u *StatisticsUseCase) IncrementInventoryCreated() {
	u.repo.SaveInventoryEvent()
}

func (u *StatisticsUseCase) IncrementInventoryUpdated() {
	u.repo.SaveInventoryUpdate()
}

func (u *StatisticsUseCase) IncrementInventoryDeleted() {
	u.repo.SaveInventoryDelete()
}

func (u *StatisticsUseCase) GetInventoryCount() int32 {
	return u.repo.GetInventoryCount()
}

func (u *StatisticsUseCase) IncrementOrderCreated(userID, time string) {
	log.Println("[USECASE] IncrementOrderCreated for:", userID)
	u.repo.SaveOrderCreated(userID, time)
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

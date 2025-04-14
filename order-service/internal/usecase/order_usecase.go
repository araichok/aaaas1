package usecase

import (
	"order-service/internal/domain"
	"order-service/internal/repository"
)

type OrderUsecase struct {
	repo *repository.OrderRepository
}

func NewOrderUsecase(repo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (uc *OrderUsecase) Create(order *domain.Order) error {
	return uc.repo.Create(order)
}

func (uc *OrderUsecase) GetByID(id string) (*domain.Order, error) {
	return uc.repo.GetByID(id)
}

func (uc *OrderUsecase) UpdateStatus(id string, status domain.OrderStatus) error {
	return uc.repo.UpdateStatus(id, status)
}

func (uc *OrderUsecase) GetByUser(userID string) ([]domain.Order, error) {
	return uc.repo.GetByUser(userID)
}

package usecase

import (
	"order-service/internal/domain"
	"order-service/internal/events"
	"order-service/internal/repository"
)

type OrderUsecase struct {
	repo      *repository.OrderRepository
	publisher *events.EventPublisher
}

func NewOrderUsecase(repo *repository.OrderRepository, publisher *events.EventPublisher) *OrderUsecase {
	return &OrderUsecase{repo: repo, publisher: publisher}
}

func (uc *OrderUsecase) Create(order *domain.Order) error {
	err := uc.repo.Create(order)
	if err == nil {
		uc.publisher.PublishOrderEvent("created", order.ID.Hex(), order.UserID)
	}
	return err
}

func (uc *OrderUsecase) GetByID(id string) (*domain.Order, error) {
	return uc.repo.GetByID(id)
}

func (uc *OrderUsecase) UpdateOrderStatus(id string, status domain.OrderStatus) error {
	order, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	err = uc.repo.UpdateStatus(id, status)
	if err == nil {
		uc.publisher.PublishOrderEvent("updated", order.ID.Hex(), order.UserID)
	}
	return err
}

func (uc *OrderUsecase) GetByUser(userID string) ([]domain.Order, error) {
	return uc.repo.GetByUser(userID)
}

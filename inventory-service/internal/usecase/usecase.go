package usecase

import (
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type ProductUsecase interface {
	Create(p *domain.Product) error
	GetByID(id string) (*domain.Product, error)
	Update(p *domain.Product) error
	Delete(id string) error
	List() ([]domain.Product, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: r}
}

func (u *productUsecase) Create(p *domain.Product) error {
	return u.repo.Create(p)
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	return u.repo.GetByID(id)
}

func (u *productUsecase) Update(p *domain.Product) error {
	return u.repo.Update(p)
}

func (u *productUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *productUsecase) List() ([]domain.Product, error) {
	return u.repo.List()
}

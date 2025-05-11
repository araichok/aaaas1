package usecase

import (
	"fmt"
	"inventory-service/internal/cache"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
	"time"
)

type ProductUsecase interface {
	Create(p *domain.Product) error
	GetByID(id string) (*domain.Product, error)
	Update(p *domain.Product) error
	Delete(id string) error
	List() ([]domain.Product, error)
}

type productUsecase struct {
	repo  repository.ProductRepository
	cache *cache.ProductCache
}

func NewProductUsecase(r repository.ProductRepository, c *cache.ProductCache) ProductUsecase {
	uc := &productUsecase{repo: r, cache: c}

	products, _ := r.List()
	c.LoadFromDB(products)

	go func() {
		ticker := time.NewTicker(12 * time.Hour)

		for range ticker.C {
			products, _ := r.List()
			c.LoadFromDB(products)
			fmt.Println("Cache refreshed at", time.Now())
		}
	}()

	return uc
}

func (u *productUsecase) Create(p *domain.Product) error {
	err := u.repo.Create(p)
	if err == nil {
		u.cache.Set(*p)
	}
	return err
}

func (u *productUsecase) GetByID(id string) (*domain.Product, error) {
	if p, found := u.cache.Get(id); found {
		return &p, nil
	}
	return u.repo.GetByID(id)
}

func (u *productUsecase) Update(p *domain.Product) error {
	err := u.repo.Update(p)
	if err == nil {
		u.cache.Set(*p)
	}
	return err
}

func (u *productUsecase) Delete(id string) error {
	err := u.repo.Delete(id)
	if err == nil {

	}
	return err
}

func (u *productUsecase) List() ([]domain.Product, error) {
	return u.cache.GetAll(), nil
}

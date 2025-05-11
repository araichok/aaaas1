package cache

import (
	"fmt"
	"inventory-service/internal/domain"
	"sync"
)

type ProductCache struct {
	products map[string]domain.Product
	mu       sync.RWMutex
}

func NewProductCache() *ProductCache {
	return &ProductCache{
		products: make(map[string]domain.Product),
	}
}

func (c *ProductCache) Set(p domain.Product) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products[p.ID] = p
	fmt.Println("Product cached:", p.ID)
}

func (c *ProductCache) Get(id string) (domain.Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, found := c.products[id]
	if found {
		fmt.Println("Cache HIT for ID:", id)
	} else {
		fmt.Println("Cache MISS for ID:", id)
	}
	return p, found
}

func (c *ProductCache) GetAll() []domain.Product {
	c.mu.RLock()
	defer c.mu.RUnlock()
	list := make([]domain.Product, 0, len(c.products))
	for _, p := range c.products {
		list = append(list, p)
	}
	fmt.Printf("Cache returned %d products ", len(list))
	return list
}

func (c *ProductCache) LoadFromDB(products []domain.Product) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products = make(map[string]domain.Product)
	for _, p := range products {
		c.products[p.ID] = p
	}
	fmt.Printf("Cache initialized with %d products ", len(products))
}

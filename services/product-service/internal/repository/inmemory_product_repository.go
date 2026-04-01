package repository

import (
	"sort"
	"sync"
	"time"

	"github.com/your-org/project-business/services/product-service/internal/model"
)

type InMemoryProductRepository struct {
	mu       sync.RWMutex
	nextID   int64
	products map[int64]model.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		nextID:   1,
		products: make(map[int64]model.Product),
	}
}

func (r *InMemoryProductRepository) Create(product model.Product) (model.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now().UTC()
	product.ID = r.nextID
	product.CreatedAt = now
	product.UpdatedAt = now

	r.products[product.ID] = product
	r.nextID++

	return product, nil
}

func (r *InMemoryProductRepository) List() ([]model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	items := make([]model.Product, 0, len(r.products))
	for _, p := range r.products {
		items = append(items, p)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})

	return items, nil
}

func (r *InMemoryProductRepository) GetByID(id int64) (model.Product, bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.products[id]
	return p, ok, nil
}

func (r *InMemoryProductRepository) Update(id int64, product model.Product) (model.Product, bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.products[id]
	if !ok {
		return model.Product{}, false, nil
	}

	existing.Name = product.Name
	existing.SKU = product.SKU
	existing.Price = product.Price
	existing.Stock = product.Stock
	existing.UpdatedAt = time.Now().UTC()

	r.products[id] = existing
	return existing, true, nil
}

func (r *InMemoryProductRepository) Delete(id int64) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.products[id]; !ok {
		return false, nil
	}

	delete(r.products, id)
	return true, nil
}

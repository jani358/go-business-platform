package service

import (
	"errors"
	"strings"

	"github.com/your-org/project-business/services/product-service/internal/dto"
	"github.com/your-org/project-business/services/product-service/internal/model"
	"github.com/your-org/project-business/services/product-service/internal/repository"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("product not found")
)

type ProductService interface {
	Create(input dto.CreateProductRequest) (model.Product, error)
	List() ([]model.Product, error)
	GetByID(id int64) (model.Product, error)
	Update(id int64, input dto.UpdateProductRequest) (model.Product, error)
	Delete(id int64) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(input dto.CreateProductRequest) (model.Product, error) {
	if err := validateProductInput(input.Name, input.SKU, input.Price, input.Stock); err != nil {
		return model.Product{}, err
	}

	return s.repo.Create(model.Product{
		Name:  strings.TrimSpace(input.Name),
		SKU:   strings.TrimSpace(input.SKU),
		Price: input.Price,
		Stock: input.Stock,
	})
}

func (s *productService) List() ([]model.Product, error) {
	return s.repo.List()
}

func (s *productService) GetByID(id int64) (model.Product, error) {
	if id <= 0 {
		return model.Product{}, ErrInvalidInput
	}

	p, ok, err := s.repo.GetByID(id)
	if err != nil {
		return model.Product{}, err
	}
	if !ok {
		return model.Product{}, ErrNotFound
	}

	return p, nil
}

func (s *productService) Update(id int64, input dto.UpdateProductRequest) (model.Product, error) {
	if id <= 0 {
		return model.Product{}, ErrInvalidInput
	}
	if err := validateProductInput(input.Name, input.SKU, input.Price, input.Stock); err != nil {
		return model.Product{}, err
	}

	updated, ok, err := s.repo.Update(id, model.Product{
		Name:  strings.TrimSpace(input.Name),
		SKU:   strings.TrimSpace(input.SKU),
		Price: input.Price,
		Stock: input.Stock,
	})
	if err != nil {
		return model.Product{}, err
	}
	if !ok {
		return model.Product{}, ErrNotFound
	}

	return updated, nil
}

func (s *productService) Delete(id int64) error {
	if id <= 0 {
		return ErrInvalidInput
	}

	ok, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotFound
	}

	return nil
}

func validateProductInput(name, sku string, price float64, stock int) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(sku) == "" {
		return ErrInvalidInput
	}
	if price <= 0 {
		return ErrInvalidInput
	}
	if stock < 0 {
		return ErrInvalidInput
	}
	return nil
}

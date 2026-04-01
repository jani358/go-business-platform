package repository

import "github.com/your-org/project-business/services/product-service/internal/model"

type ProductRepository interface {
	Create(product model.Product) (model.Product, error)
	List() ([]model.Product, error)
	GetByID(id int64) (model.Product, bool, error)
	Update(id int64, product model.Product) (model.Product, bool, error)
	Delete(id int64) (bool, error)
}

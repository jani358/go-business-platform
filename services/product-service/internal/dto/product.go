package dto

type CreateProductRequest struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

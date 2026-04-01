package main

import (
	"log"
	"net/http"

	"github.com/your-org/project-business/services/product-service/internal/handler"
	"github.com/your-org/project-business/services/product-service/internal/repository"
	"github.com/your-org/project-business/services/product-service/internal/service"
)

func main() {
	repo := repository.NewInMemoryProductRepository()
	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("product-service: ok"))
	})
	mux.HandleFunc("/products", productHandler.HandleCollection)
	mux.HandleFunc("/products/", productHandler.HandleItem)

	log.Println("product-service running on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}

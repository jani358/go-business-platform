package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/your-org/project-business/services/product-service/internal/dto"
	"github.com/your-org/project-business/services/product-service/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) HandleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.list(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *ProductHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	id, ok := parseProductID(r.URL.Path)
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid product id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) list(w http.ResponseWriter, _ *http.Request) {
	products, err := h.service.List()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) getByID(w http.ResponseWriter, id int64) {
	product, err := h.service.GetByID(id)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) update(w http.ResponseWriter, r *http.Request, id int64) {
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	product, err := h.service.Update(id, req)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) delete(w http.ResponseWriter, id int64) {
	if err := h.service.Delete(id); err != nil {
		writeServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func parseProductID(path string) (int64, bool) {
	const prefix = "/products/"
	if !strings.HasPrefix(path, prefix) {
		return 0, false
	}

	idStr := strings.TrimSpace(strings.TrimPrefix(path, prefix))
	if idStr == "" || strings.Contains(idStr, "/") {
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, false
	}
	return id, true
}

func writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	case errors.Is(err, service.ErrNotFound):
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

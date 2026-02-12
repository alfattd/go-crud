package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alfattd/crud/internal/dto/request"
	"github.com/alfattd/crud/internal/dto/response"
	"github.com/alfattd/crud/internal/repository"
	"github.com/alfattd/crud/internal/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.svc.CreateProduct(req)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.CreateProductResponse{
		Product: response.ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Price:      product.Price,
			Stock:      product.Stock,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID required to be filled in", http.StatusBadRequest)
		return
	}

	var req request.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.svc.UpdateProduct(id, req)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Product or Category not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.UpdateProductResponse{
		Product: response.ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Price:      product.Price,
			Stock:      product.Stock,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID required to be filled in", http.StatusBadRequest)
		return
	}

	err := h.svc.DeleteProduct(id)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.DeleteProductResponse{
		ID:      id,
		Message: "Product Successfully Deleted",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID required to be filled in", http.StatusBadRequest)
		return
	}

	product, err := h.svc.GetProductByID(id)
	if err != nil {
		if err == repository.ErrNotFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := response.GetProductResponse{
		Product: response.ProductResponse{
			ID:         product.ID,
			Name:       product.Name,
			CategoryID: product.CategoryID,
			Price:      product.Price,
			Stock:      product.Stock,
			CreatedAt:  product.CreatedAt,
			UpdatedAt:  product.UpdatedAt,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) ListProduct(w http.ResponseWriter, r *http.Request) {

	products, err := h.svc.ListProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resList := make([]response.ProductResponse, len(products))

	for i, p := range products {
		resList[i] = response.ProductResponse{
			ID:         p.ID,
			Name:       p.Name,
			CategoryID: p.CategoryID,
			Price:      p.Price,
			Stock:      p.Stock,
			CreatedAt:  p.CreatedAt,
			UpdatedAt:  p.UpdatedAt,
		}
	}

	res := response.ListProductResponse{
		Products: resList,
		Total:    len(resList),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now(),
	})
}

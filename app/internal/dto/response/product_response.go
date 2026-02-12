package response

import "time"

type ProductResponse struct {
	ID         string            `json:"id"`
	Name       string            `json:"Name"`
	CategoryID string            `json:"category_id"`
	Category   *CategoryResponse `json:"category,omitempty"`
	Price      float64           `json:"price"`
	Stock      int               `json:"stock"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type DeleteProductResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type CreateProductResponse struct {
	Product ProductResponse `json:"product"`
}

type UpdateProductResponse struct {
	Product ProductResponse `json:"product"`
}

type ListProductResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int               `json:"total"`
}

type GetProductResponse struct {
	Product ProductResponse `json:"product"`
}

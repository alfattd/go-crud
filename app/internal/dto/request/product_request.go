package request

type CreateProductRequest struct {
	Name       string  `json:"name" binding:"required"`
	CategoryID string  `json:"category_id" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	Stock      int     `json:"stock" binding:"required"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name,omitempty"`
	CategoryID string  `json:"category_id,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Stock      int     `json:"stock,omitempty"`
}

type DeleteProductRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetProductRequest struct {
	ID string `json:"id" binding:"required"`
}

type ListProductRequest struct {
	Name       string `json:"name,omitempty"`
	CategoryID string `json:"category_id,omitempty"`
}

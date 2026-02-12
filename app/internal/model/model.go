package model

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CategoryID string    `json:"category_id"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category *Category `json:"category,omitempty"`
}

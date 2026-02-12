package repository

import "github.com/alfattd/crud/internal/model"

type ProductRepository interface {
	Create(p *model.Product) error
	GetByID(id string) (*model.Product, error)
	Update(p *model.Product) error
	Delete(id string) error
	List() ([]*model.Product, error)
}

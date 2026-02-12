package service

import (
	"time"

	"github.com/alfattd/crud/internal/dto/request"
	"github.com/alfattd/crud/internal/model"
	"github.com/alfattd/crud/internal/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	repo         repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(r repository.ProductRepository, kr repository.CategoryRepository) *ProductService {
	return &ProductService{
		repo:         r,
		categoryRepo: kr,
	}
}

func (s *ProductService) CreateProduct(req request.CreateProductRequest) (*model.Product, error) {
	_, err := s.categoryRepo.GetByID(req.CategoryID)
	if err != nil {
		return nil, err
	}

	product := &model.Product{
		ID:         uuid.NewString(),
		Name:       req.Name,
		CategoryID: req.CategoryID,
		Price:      req.Price,
		Stock:      req.Stock,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err = s.repo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(id string, req request.UpdateProductRequest) (*model.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.CategoryID != "" {
		_, err := s.categoryRepo.GetByID(req.CategoryID)
		if err != nil {
			return nil, err
		}
		product.CategoryID = req.CategoryID
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}

	product.UpdatedAt = time.Now()
	err = s.repo.Update(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetProductByID(id string) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) ListProduct() ([]*model.Product, error) {
	return s.repo.List()
}

package memory

import (
	"sync"

	"github.com/alfattd/crud/internal/model"
	"github.com/alfattd/crud/internal/repository"
)

type InMemoryProductRepo struct {
	mu       sync.RWMutex
	products map[string]*model.Product
}

func NewInMemoryProductRepo() *InMemoryProductRepo {
	return &InMemoryProductRepo{
		products: make(map[string]*model.Product),
	}
}

func (r *InMemoryProductRepo) Create(p *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[p.ID] = p
	return nil
}

func (r *InMemoryProductRepo) Update(p *model.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.products[p.ID]
	if !ok {
		return repository.ErrNotFound
	}
	r.products[p.ID] = p
	return nil
}

func (r *InMemoryProductRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.products[id]
	if !ok {
		return repository.ErrNotFound
	}
	delete(r.products, id)
	return nil
}

func (r *InMemoryProductRepo) GetByID(id string) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.products[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return p, nil
}

func (r *InMemoryProductRepo) List() ([]*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*model.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}

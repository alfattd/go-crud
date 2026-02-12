package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/alfattd/crud/internal/model"
	"github.com/alfattd/crud/internal/repository"
)

type postgresProductRepo struct {
	db *sql.DB
}

func NewPostgresProductRepo(db *sql.DB) repository.ProductRepository {
	return &postgresProductRepo{db: db}
}

func (r *postgresProductRepo) Create(p *model.Product) error {
	query := `
	INSERT INTO products (id, name, category_id, price, stock, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(query, p.ID, p.Name, p.CategoryID, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt)
	return err
}

func (r *postgresProductRepo) Update(p *model.Product) error {
	p.UpdatedAt = time.Now()
	query := `
	UPDATE products
	SET name = $2, category_id = $3, price = $4, stock = $5, updated_at = $6
	WHERE id = $1
	`
	res, err := r.db.Exec(query, p.ID, p.Name, p.CategoryID, p.Price, p.Stock, p.UpdatedAt)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *postgresProductRepo) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *postgresProductRepo) GetByID(id string) (*model.Product, error) {
	query := `
	SELECT id, name, category_id, price, stock, created_at, updated_at
	FROM products
	WHERE id = $1
	`

	row := r.db.QueryRow(query, id)
	p := &model.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return p, nil
}

func (r *postgresProductRepo) List() ([]*model.Product, error) {
	query := `
	SELECT id, name, category_id, price, stock, created_at, updated_at
	FROM products
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*model.Product
	for rows.Next() {
		p := &model.Product{}
		if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

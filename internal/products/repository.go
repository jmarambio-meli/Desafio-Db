package products

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(product *domain.Product) (int64, error)
	ReadAll() ([]domain.Product, error)
	LoadData(*domain.Product) error
	Top5Products() ([]domain.TopProducts, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product *domain.Product) (int64, error) {
	query := `INSERT INTO products (description, price) VALUES (?, ?)`
	row, err := r.db.Exec(query, &product.Description, &product.Price)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]domain.Product, error) {
	query := `SELECT id, description, price FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]domain.Product, 0)
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) LoadData(product *domain.Product) error {
	query := `INSERT INTO products (id, description, price) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, &product.Id, &product.Description, &product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Top5Products() ([]domain.TopProducts, error) {
	query := `
	SELECT 
		p.description, 
		sum(quantity) as cantidad 
	FROM products AS p
	INNER JOIN sales AS s ON s.product_id = p.id
	GROUP BY p.id
	ORDER BY cantidad DESC
	LIMIT 5;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	TotalTopProducts := make([]domain.TopProducts, 0)
	for rows.Next() {
		topProducts := domain.TopProducts{}
		err := rows.Scan(&topProducts.Description, &topProducts.Total)
		if err != nil {
			return nil, err
		}
		TotalTopProducts = append(TotalTopProducts, topProducts)
	}
	return TotalTopProducts, nil
}

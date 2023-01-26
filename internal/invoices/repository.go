package invoices

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	LoadData(invoice *domain.Invoices) error
	FixData(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) LoadData(invoice *domain.Invoices) error {
	query := `INSERT INTO invoices (id, datetime, customer_id, total) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, &invoice.Id, &invoice.Datetime, &invoice.CustomerId, &invoice.Total)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FixData(id int) error {
	query := `
	UPDATE invoices AS i SET 
	total = (SELECT 
		ROUND(sum(p.price * s.quantity),2) AS subtotal
	FROM sales AS s
	INNER JOIN products AS p ON p.id = s.product_id
	WHERE s.invoice_id = ?
	GROUP BY s.invoice_id)
	WHERE i.id = ?;	
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

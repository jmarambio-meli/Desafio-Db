package customers

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(customers *domain.Customers) (int64, error)
	ReadAll() ([]*domain.Customers, error)
	LoadData(customers *domain.Customers) error
	TotalCustomerByCondition() ([]domain.TotalCustomerByCondition, error)
	TotalAmountActiveCustomers() ([]domain.TotalAmountActiveCustomers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers *domain.Customers) (int64, error) {
	query := `INSERT INTO customers (first_name, last_name, customers.condition) VALUES (?, ?, ?, ?)`
	row, err := r.db.Exec(query, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *repository) LoadData(customers *domain.Customers) error {
	query := `INSERT INTO customers (id, first_name, last_name, customers.condition) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, &customers.Id, &customers.FirstName, &customers.LastName, &customers.Condition)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) TotalCustomerByCondition() ([]domain.TotalCustomerByCondition, error) {
	query := `
	SELECT 
	CASE 
		WHEN c.condition = 0 THEN "Inactivo (0)"
		WHEN c.condition = 1 THEN "Activo (1)"
	END AS "Condition",
		ROUND(sum(total),2) 
	FROM customers as c
	INNER JOIN invoices AS i ON i.customer_id = c.id
	GROUP BY c.condition;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	TotalCustomers := make([]domain.TotalCustomerByCondition, 0)
	for rows.Next() {
		TotalCustomerByCondition := domain.TotalCustomerByCondition{}
		err := rows.Scan(&TotalCustomerByCondition.Condition, &TotalCustomerByCondition.Total)
		if err != nil {
			return nil, err
		}
		TotalCustomers = append(TotalCustomers, TotalCustomerByCondition)
	}
	return TotalCustomers, nil
}

func (r *repository) TotalAmountActiveCustomers() ([]domain.TotalAmountActiveCustomers, error) {
	query := `
	SELECT 
		c.last_name, 
		c.first_name, 
		ROUND(sum(i.total), 2) as amount 
	FROM customers as c
	INNER JOIN invoices AS i ON i.customer_id = c.id
	WHERE c.condition = 1
	GROUP BY c.id
	ORDER BY amount DESC
	LIMIT 5;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	TotalCustomers := make([]domain.TotalAmountActiveCustomers, 0)
	for rows.Next() {
		TotalAmountActiveCustomer := domain.TotalAmountActiveCustomers{}
		err := rows.Scan(&TotalAmountActiveCustomer.Last_name, &TotalAmountActiveCustomer.First_name, &TotalAmountActiveCustomer.Amount)
		if err != nil {
			return nil, err
		}
		TotalCustomers = append(TotalCustomers, TotalAmountActiveCustomer)
	}
	return TotalCustomers, nil
}

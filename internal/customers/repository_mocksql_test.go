package customers

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestTotalCustomerByCondition_SQLMock_Ok(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	repo := NewRepository(db)
	assert.NoError(t, err)
	defer db.Close()
	expectedData := []domain.TotalCustomerByCondition{
		{
			Condition: "Inactivo (0)",
			Total:     605929.11,
		},
		{
			Condition: "Activo (1)",
			Total:     716792.33,
		},
	}

	rows := mock.NewRows([]string{"condition", "total"})

	for _, d := range expectedData {
		rows.AddRow(d.Condition, d.Total)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
	SELECT 
	CASE 
		WHEN c.condition = 0 THEN "Inactivo (0)"
		WHEN c.condition = 1 THEN "Activo (1)"
	END AS "Condition",
		ROUND(sum(total),2) 
	FROM customers as c
	INNER JOIN invoices AS i ON i.customer_id = c.id
	GROUP BY c.condition;
	`)).WillReturnRows(rows)

	//act
	customers, err := repo.TotalCustomerByCondition()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedData, customers)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTotalAmountActiveCustomers_SQLMock_Ok(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	repo := NewRepository(db)
	assert.NoError(t, err)
	defer db.Close()
	expectedData := []domain.TotalAmountActiveCustomers{
		{
			Last_name:  "Tortis",
			First_name: "Lannie",
			Amount:     58513.55,
		},
		{
			Last_name:  "Crowcum",
			First_name: "Jasen",
			Amount:     48291.03,
		},
		{
			Last_name:  "Anstis",
			First_name: "Lazaro",
			Amount:     40792.06,
		},
		{
			Last_name:  "Kieran",
			First_name: "Tomasina",
			Amount:     39162.4,
		},
		{
			Last_name:  "Penbarthy",
			First_name: "Cassondra",
			Amount:     33749.85,
		},
	}

	rows := mock.NewRows([]string{"last_name", "first_name", "amount"})

	for _, d := range expectedData {
		rows.AddRow(d.Last_name, d.First_name, d.Amount)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
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
	`)).WillReturnRows(rows)

	//act
	customers, err := repo.TotalAmountActiveCustomers()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedData, customers)
	assert.NoError(t, mock.ExpectationsWereMet())
}

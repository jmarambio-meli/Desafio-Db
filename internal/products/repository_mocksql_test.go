package products

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetAll_SQLMock_Ok(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	repo := NewRepository(db)
	assert.NoError(t, err)
	defer db.Close()
	expectedData := []domain.Product{
		{
			Id:          1,
			Description: "French Pastry - Mini Chocolate",
			Price:       97.01,
		},
		{
			Id:          2,
			Description: "Beans - Soya Bean",
			Price:       12.89,
		},
	}

	rows := mock.NewRows([]string{"id", "description", "price"})

	for _, d := range expectedData {
		rows.AddRow(d.Id, d.Description, d.Price)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, description, price FROM products`)).WillReturnRows(rows)

	//act
	products, err := repo.ReadAll()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, expectedData, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_SQLMock_Error(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(`SELECT id, description, price FROM products`).WillReturnError(sql.ErrConnDone)

	rp := NewRepository(db)
	//act
	products, err := rp.ReadAll()
	//assert
	assert.ErrorIs(t, err, sql.ErrConnDone)
	assert.Empty(t, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTop5Products_OK_SQLMock_Ok(t *testing.T) {
	//arrange
	db, mock, err := sqlmock.New()
	repo := NewRepository(db)
	assert.NoError(t, err)
	defer db.Close()
	ExpectedData := []domain.TopProducts{
		{
			Description: "Vinegar - Raspberry",
			Total:       660,
		},
		{
			Description: "Flour - Corn, Fine",
			Total:       521,
		},
		{
			Description: "Cookie - Oatmeal",
			Total:       467,
		},
		{
			Description: "Pepper - Red Chili",
			Total:       439,
		},
		{
			Description: "Chocolate - Milk Coating",
			Total:       436,
		},
	}

	rows := mock.NewRows([]string{"description", "price"})

	for _, d := range ExpectedData {
		rows.AddRow(d.Description, d.Total)
	}
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT 
			p.description, 
			sum(quantity) as cantidad 
		FROM products AS p
		INNER JOIN sales AS s ON s.product_id = p.id
		GROUP BY p.id
		ORDER BY cantidad DESC
		LIMIT 5;
	`)).WillReturnRows(rows)

	//act
	products, err := repo.Top5Products()
	//assert
	assert.NoError(t, err)
	assert.Equal(t, ExpectedData, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

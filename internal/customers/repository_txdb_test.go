package customers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	txdb.Register("txdb", "mysql", "meli_sprint_user:Meli_Sprint#123@/fantasy_products")
}

func TestTotalCustomerByCondition_Ok(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
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

	rp := NewRepository(db)
	customers, err := rp.TotalCustomerByCondition()
	assert.NoError(t, err)
	assert.Equal(t, expectedData, customers)
	assert.NotEmpty(t, customers)
}

func TestTotalAmountActiveCustomers_Ok(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer db.Close()

	rp := NewRepository(db)
	products, err := rp.TotalCustomerByCondition()
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}

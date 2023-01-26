package products

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

func TestTop5Products_OK(t *testing.T) {
	db, err := sql.Open("txdb", "identifier")
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

	rp := NewRepository(db)
	products, err := rp.Top5Products()
	assert.Equal(t, ExpectedData, products)
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}

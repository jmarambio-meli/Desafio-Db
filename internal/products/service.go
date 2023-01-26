package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(product *domain.Product) error
	ReadAll() ([]domain.Product, error)
	LoadData() error
	Top5Products() ([]domain.TopProducts, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(product *domain.Product) error {
	_, err := s.r.Create(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]domain.Product, error) {
	return s.r.ReadAll()
}

func (s *service) LoadData() error {

	dataJson, err := cargarProductsJson()
	if err != nil {
		return err
	}
	for _, v := range dataJson {
		fmt.Println(v)
		err := s.r.LoadData(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Top5Products() ([]domain.TopProducts, error) {
	return s.r.Top5Products()
}

func cargarProductsJson() ([]*domain.Product, error) {
	var dataJson []*domain.Product
	raw, err := ioutil.ReadFile("./datos/products.json")
	if err != nil {
		fmt.Println(err.Error())
		return []*domain.Product{}, err
	}
	json.Unmarshal(raw, &dataJson)
	return dataJson, nil
}

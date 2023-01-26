package sales

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(sales *domain.Sales) error
	ReadAll() ([]*domain.Sales, error)
	LoadData() error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(sales *domain.Sales) error {
	_, err := s.r.Create(sales)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Sales, error) {
	return s.r.ReadAll()
}

func (s *service) LoadData() error {

	dataJson, err := cargarSalesJson()
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

func cargarSalesJson() ([]*domain.Sales, error) {
	var dataJson []*domain.Sales
	raw, err := ioutil.ReadFile("./datos/sales.json")
	if err != nil {
		fmt.Println(err.Error())
		return []*domain.Sales{}, err
	}
	json.Unmarshal(raw, &dataJson)
	return dataJson, nil
}

package customers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(customers *domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
	LoadData() error
	TotalCustomerByCondition() ([]domain.TotalCustomerByCondition, error)
	TotalAmountActiveCustomers() ([]domain.TotalAmountActiveCustomers, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(customers *domain.Customers) error {
	_, err := s.r.Create(customers)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}

func (s *service) LoadData() error {

	dataJson, err := cargarCustomersJson()
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

func (s *service) TotalCustomerByCondition() ([]domain.TotalCustomerByCondition, error) {
	return s.r.TotalCustomerByCondition()
}

func (s *service) TotalAmountActiveCustomers() ([]domain.TotalAmountActiveCustomers, error) {
	return s.r.TotalAmountActiveCustomers()
}

func cargarCustomersJson() ([]*domain.Customers, error) {
	var dataJson []*domain.Customers
	raw, err := ioutil.ReadFile("./datos/customers.json")
	if err != nil {
		fmt.Println(err.Error())
		return []*domain.Customers{}, err
	}
	json.Unmarshal(raw, &dataJson)
	return dataJson, nil
}

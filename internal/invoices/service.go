package invoices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	LoadData() error
	FixData() error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) ReadAll() ([]*domain.Invoices, error) {
	return s.r.ReadAll()
}

func (s *service) LoadData() error {

	dataJson, err := cargarInvoicesJson()
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

func cargarInvoicesJson() ([]*domain.Invoices, error) {
	var dataJson []*domain.Invoices
	raw, err := ioutil.ReadFile("./datos/invoices.json")
	if err != nil {
		fmt.Println(err.Error())
		return []*domain.Invoices{}, err
	}
	json.Unmarshal(raw, &dataJson)
	return dataJson, nil
}

func (s *service) FixData() error {
	invoices, err := s.r.ReadAll()
	if err != nil {
		return err
	}
	for _, v := range invoices {
		err := s.r.FixData(v.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

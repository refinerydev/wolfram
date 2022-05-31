package customer

import (
	"github.com/refinerydev/wolfram/src/database/entity"
)

type Services interface {
	AddCustomer(req *request) error
}

type services struct {
	repo repository
}

func NewService() *services {
	repo := customerRepository()
	return &services{repo}
}

func (s *services) AddCustomer(req *request) error {
	customer := entity.Customer{}
	customer.Name = req.Name
	customer.Email = req.Email

	err := s.repo.create(customer)
	if err != nil {
		return err
	}

	return nil
}

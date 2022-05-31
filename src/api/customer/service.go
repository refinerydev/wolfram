package customer

import (
	"github.com/refinerydev/wolfram/src/database/entity"
)

type Services interface {
	AddNew(req *request) error
	GetAll() ([]entity.Customer, error)
	Edit(req *request) error
	Remove(customerId uint) error
}

type services struct {
	repo repository
}

func NewService() *services {
	repo := customerRepository()
	return &services{repo}
}

func (s *services) AddNew(req *request) error {
	customer := entity.Customer{}
	customer.Name = req.Name
	customer.Email = req.Email
	customer.PhoneNumber = req.PhoneNumber
	customer.Address = req.Address
	customer.Latitude = req.Latitude
	customer.Longitude = req.Longitude
	customer.CustomerType = req.CustomerType

	if err := s.repo.create(customer); err != nil {
		return err
	}

	return nil
}

func (s *services) GetAll() ([]entity.Customer, error) {
	customers, err := s.repo.read()
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (s *services) Edit(req *request, customerId uint) error {
	customer := entity.Customer{}
	customer.ID = customerId
	customer.Name = req.Name
	customer.Email = req.Email

	if err := s.repo.update(customer); err != nil {
		return err
	}

	return nil
}

func (s *services) Remove(customerId uint) error {
	customer := entity.Customer{}
	customer.ID = customerId

	if err := s.repo.delete(customer); err != nil {
		return err
	}

	return nil
}

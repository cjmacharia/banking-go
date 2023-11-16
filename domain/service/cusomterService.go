package service

import (
	"github.com/cjmash/banking/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	PostCustomer(customer *domain.Customer) error
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) PostCustomer(customer *domain.Customer) error {
	return s.repo.CreateCustomer(customer)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

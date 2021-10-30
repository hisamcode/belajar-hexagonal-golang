package service

import "github.com/hisamcode/belajar-hexagonal-golang/domain"

type CustomerService interface {
  GetAllCustomer() ([]domain.Customer, error)
  GetCustomer(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
  repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
  return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error) {
  return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
  return DefaultCustomerService{repo: repository}
}

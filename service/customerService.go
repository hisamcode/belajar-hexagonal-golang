package service

import (
  "github.com/hisamcode/belajar-hexagonal-golang/domain"
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
)

type CustomerService interface {
  GetAllCustomer() ([]domain.Customer, *errs.AppError)
  GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
  repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
  return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
  return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
  return DefaultCustomerService{repo: repository}
}

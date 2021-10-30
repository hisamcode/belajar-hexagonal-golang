package service

import (
  "github.com/hisamcode/belajar-hexagonal-golang/domain"
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
)

type CustomerService interface {
  GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
  GetCustomer(id string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
  repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
  if status == "active" {
    status = "1"
  } else if status == "inactive" {
    status = "0"
  } else {
    status = ""
  }
  return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
  return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
  return DefaultCustomerService{repo: repository}
}

package service

import "github.com/"

type CustomerService interface {
  GetAllCustomer() ([]domain.Customer, error)
}
package domain

import (
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
  ById(string) (*Customer, *errs.AppError)
}

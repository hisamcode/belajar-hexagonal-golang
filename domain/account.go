package domain

import (
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
  "github.com/hisamcode/belajar-hexagonal-golang/dto"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
  Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
  return dto.NewAccountResponse{a.AccountId}
}
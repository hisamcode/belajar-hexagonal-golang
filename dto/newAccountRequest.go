package dto

import (
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
)

type NewAccountRequest struct {
  CustomerId string `json:"customer_id"`
  AccountType string `json:"account_type"`
  Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppErro {
  if (r.Amount < 5000) {
    return errs.NewValidationError()
  }
}


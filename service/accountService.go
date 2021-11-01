package service

import (
  "github.com/hisamcode/belajar-hexagonal-golang/dto"
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
  "github.com/hisamcode/belajar-hexagonal-golang/domain"
)

type AccountService interface {
  NewAccount(dto.NewAccountRequest) (*dtoNewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
  repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount (dto.NewAccountRequest) (*dto.newAccountResponse, errs.AppError) {
  a := domain.Account{
    AccountId: "",
    CustomerId: req.CustomerId,
    OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
    AccountType: req.AccountType,
    Amount: req.Amount,
    Status: "1",
  }
  newAccount, err := s.repo.Save(a)

  if err != nil {
    return nil, err
  }

  response := newAccount.ToNewAccountResponseDto()

  return &response, nil
}

func NewAccountService(repo AccountRepository) DefaultAccountService {
  return DefaultAccountService{repo}
}
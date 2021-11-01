package domain

import (
  "strconv"
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
  "github.com/hisamcode/belajar-hexagonal-golang/logger"
  "github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
  client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
  sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

  result, err := d.client.Exec(sqlInsert, a.CustumerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
  if err != nil {
    logger.Error("Error while creating new account : " + err.Error())
    return nil, errs.NewUnexpectedError("Unexpected error from database")
  }

  id, err := result.LastInsertId()
  if err != nil {
    logger.Error("Error while getting last insert id for new account : " + err.Error())
    return nil, errs.NewUnexpectedError("Unexpected error from database")
  }

  a.AccountId = strconv.FormatInt(id, 10)
  return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
  return AccountRepositoryDb{dbClient}
}
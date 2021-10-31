package domain

import (
  "database/sql"
  "time"

  _ "github.com/go-sql-driver/mysql"
  "github.com/hisamcode/belajar-hexagonal-golang/errs"
  "github.com/hisamcode/belajar-hexagonal-golang/logger"
)

type CustomerRepositoryDb struct {
  client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

  var rows *sql.Rows
  var err error
  
  if status == "" {
    findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
    rows, err = d.client.Query(findAllSql)
  } else {
    findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
    rows, err = d.client.Query(findAllSql, status)
  }
  
  if err != nil {
    logger.Error("Error while querying customers table " + err.Error())
    return nil, errs.NewUnexpectedError("Unexpected database error")
  }

  customers := make([]Customer, 0)
  for rows.Next() {
    var c Customer

    err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

    if err != nil {
      logger.Error("Error while scanning customers table " + err.Error())
      return nil, errs.NewUnexpectedError("Unexpected database error")
    }

    customers = append(customers, c)
  }

  return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
  customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

  row := d.client.QueryRow(customerSql, id)

  var c Customer
  err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

  if err != nil {
    if err == sql.ErrNoRows {
      return nil, errs.NewNotFoundError("Customer not found")
    } else {
      logger.Error("error while scanning customer " + err.Error())
      return nil, errs.NewUnexpectedError("Unexpected database error")
    }
  }

  return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
    client, err := sql.Open("mysql", "uppxago76tywvdd8:xwYZw66BhKm0lIAkJKn4@tcp(bvjh6qwrrq8wqlmgcp8s-mysql.services.clever-cloud.com:3306)/bvjh6qwrrq8wqlmgcp8s")
    if err != nil {
      panic(err)
    }
    // See "Important settings" section.
    client.SetConnMaxLifetime(time.Minute * 3)
    client.SetMaxOpenConns(10)
    client.SetMaxIdleConns(10)

    return CustomerRepositoryDb{client}
}
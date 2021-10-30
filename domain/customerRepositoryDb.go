package domain

import (
  "database/sql"
  "time"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
  client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

  findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

  rows, err := d.client.Query(findAllSql)
  if err != nil {
    log.Println("Error while querying customers table " + err.Error())
    return nil, err
  }

  customers := make([]Customer, 0)
  for rows.Next() {
    var c Customer

    err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

    if err != nil {
      log.Println("Error while scanning customers table " + err.Error())
      return nil, err
    }

    customers = append(customers, c)
  }

  return customers, nil
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
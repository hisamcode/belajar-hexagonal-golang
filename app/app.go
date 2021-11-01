package app

import (
	"log"
	"net/http"
  "fmt"
  "os"
  "time"

	"github.com/gorilla/mux"
  "github.com/jmoiron/sqlx"

  "github.com/hisamcode/belajar-hexagonal-golang/service"
  "github.com/hisamcode/belajar-hexagonal-golang/domain"
  "github.com/hisamcode/belajar-hexagonal-golang/logger"
)

func sanityCheck() {
  if os.Getenv("SERVER_ADDRESS") == "" ||
    os.Getenv("SERVER_PORT") == "" {
      logger.Error("environment SERVER variable not defined...")
    }

  if os.Getenv("DB_USER") == "" ||
    os.Getenv("DB_PASSWORD") == "" ||
    os.Getenv("DB_ADDRESS") == "" ||
    os.Getenv("DB_PORT") == "" ||
    os.Getenv("DB_NAME") == "" {
      logger.Error("environment DB variable not defined...")
    }
}

func Start() {

  sanityCheck()

	router := mux.NewRouter()

  // wiring
  // ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
  dbClient := getDbClient()
  customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
  accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

  ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
  ah := AccountHandler{service: service.NewAccountService(accountRepositoryDb)}

	// define routes
  router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
  router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
  router.HandleFunc("/customers/{customer_id:[0-9]+/account}", ah.NewAccount).Methods(http.MethodPost)

	// starting server
  address := os.Getenv("SERVER_ADDRESS")
  port := os.Getenv("SERVER_PORT")
  replit := os.Getenv("REPLIT")
  if replit == "1" {
    log.Fatal(http.ListenAndServe(":0", router))
  } else {
    log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
  }
}

func getDbClient() *sqlx.DB {
  dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbAddress := os.Getenv("DB_ADDRESS")
  dbPort := os.Getenv("DB_PORT")
  dbName := os.Getenv("DB_NAME")
  dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
  client, err := sqlx.Open("mysql", dataSource)
	// client, err := sqlx.Open("mysql", "uppxago76tywvdd8:xwYZw66BhKm0lIAkJKn4@tcp(bvjh6qwrrq8wqlmgcp8s-mysql.services.clever-cloud.com:3306)/bvjh6qwrrq8wqlmgcp8s")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

  return client
}

package app

import (
	"log"
	"net/http"
  "fmt"
  "os"
	"github.com/gorilla/mux"
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
  ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
  router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
  router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

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
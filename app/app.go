package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
  "github.com/hisamcode/belajar-hexagonal-golang/service"
  "github.com/hisamcode/belajar-hexagonal-golang/domain"
)

func Start() {

	router := mux.NewRouter()

  // wiring
  // ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
  ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
  router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":0", router))
}
package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	Router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//Defines routes
	Router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")

	//Starting server
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe("localhost:8080", Router))
}

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

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//Defines routes
	Router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	Router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	//Starting server
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe("localhost:3000", Router))
}

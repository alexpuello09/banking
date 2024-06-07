package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USERB") == "" ||
		os.Getenv("DB_PASSB") == "" ||
		os.Getenv("DB_ADDRB") == "" ||
		os.Getenv("DB_NAMEB") == "" ||
		os.Getenv("DB_PORTB") == "" {
		log.Fatal("There is an enviroment variable Not defined")
	}
}

func Start() {

	sanityCheck()
	Router := mux.NewRouter()

	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//Defines routes
	Router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	Router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	//Starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), Router))
}

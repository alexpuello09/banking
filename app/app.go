package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
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

	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	//Defines routes
	Router.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	Router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")
	Router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//Starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), Router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USERB")
	dbPass := os.Getenv("DB_PASSB")
	dbAddr := os.Getenv("DB_ADDRB")
	dbName := os.Getenv("DB_NAMEB")
	dbPort := os.Getenv("DB_PORTB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

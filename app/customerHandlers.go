package app

import (
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customers)
	}

	//if r.Header.Get("Content-Type") != "application/json" {
	//	w.Header().Add("Content-Type", "application/json")
	//	err := json.NewEncoder(w).Encode(customers)
	//	if err != nil {
	//		panic(err)
	//	}
	//} else {
	//	err := json.NewEncoder(w).Encode(customers)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writerResponse(w, err.Code, err.AsMessage())
	} else {
		writerResponse(w, http.StatusOK, customer)
	}
}

func writerResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}

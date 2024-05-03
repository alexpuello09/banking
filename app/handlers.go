package app

import (
	"banking/service"
	"encoding/json"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (s CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := s.service.GetAllCustomers()

	if r.Header.Get("Content-Type") != "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		json.NewEncoder(w).Encode(customers)
	}
}

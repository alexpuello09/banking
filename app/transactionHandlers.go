package app

import (
	"banking/service"
	"net/http"
)

type TransactionsHandler struct {
	service service.TransactionService
}

func (h TransactionsHandler) NewTransaction(r http.ResponseWriter, request *http.Request) {

}

package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionsHandler struct {
	service service.TransactionService
}

func (h TransactionsHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.AccountId = accountId
		responseDto, appErr := h.service.MakeTransaction(req)
		if appErr != nil {
			writerResponse(w, http.StatusInternalServerError, appErr.Message)
		} else {

			result := dto.TransactionResponse{
				TransactionId: responseDto.TransactionId,
				Amount:        responseDto.Amount,
			}
			writerResponse(w, http.StatusOK, result)
		}
	}
}

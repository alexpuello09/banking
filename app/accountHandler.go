package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writerResponse(w, http.StatusBadRequest, err.Error())
	} else {
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writerResponse(w, appError.Code, appError.Message)
		} else {
			writerResponse(w, http.StatusCreated, account)
		}
	}
}

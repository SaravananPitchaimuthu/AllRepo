package app

import (
	"encoding/json"
	"net/http"

	"github.com/SaravananPitchaimuthu/pub-api/dto"
	"github.com/SaravananPitchaimuthu/pub-api/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, AppError := ah.service.NewAccount(request)
		if AppError != nil {
			WriteResponse(w, AppError.Code, AppError.Message)
		} else {
			WriteResponse(w, http.StatusOK, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			WriteResponse(w, appError.Code, appError.Message)
		} else {
			WriteResponse(w, http.StatusOK, account)
		}

	}

}

package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"time"
)

type TransactionService struct {
	repo domain.TransactionRepositoryDb
}

type ItransactionService interface {
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

func (t TransactionService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err1 := req.ValidateTransactionType()
	if err1 != nil {
		return nil, err1
	}
	err2 := req.ValidateNegativeAmount()
	if err2 != nil {
		return nil, err2
	}

	transactionObject := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err3 := t.repo.SaveTransaction(transactionObject)
	if err3 != nil {
		logger.Error("TransactionService - SaveTransaction Error" + err3.Message)
		return nil, err3
	}

	response := dto.TransactionResponse{
		TransactionId: newTransaction.TransactionId,
		Amount:        newTransaction.Amount,
	}
	return &response, nil
}

func NewHelperTransaction(repo domain.TransactionRepositoryDb) TransactionService {
	return TransactionService{repo: repo}
}

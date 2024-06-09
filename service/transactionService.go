package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"time"
)

type TransactionService struct {
	repo domain.TransactionRepository
}

type ItransactionService interface {
	makeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

func (t TransactionService) makeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	transactionObject := domain.Transaction{
		TransactionId:   0,
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := t.repo.SaveTransaction(transactionObject)
	if err != nil {
		return nil, err
	}

	response := dto.TransactionResponse{
		TransactionId: newTransaction.TransactionId,
		Amount:        newTransaction.Amount,
	}
	return &response, nil
}

func NewHelperTransaction(repo *domain.TransactionRepository) TransactionService {
	return TransactionService{repo: repo}
}

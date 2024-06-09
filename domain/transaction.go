package domain

import "banking/errs"

type Transaction struct {
	TransactionId   int     `json:"transaction_id"`
	AccountId       int     `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

type ITransactionRepository interface {
	SaveTransaction(transaction *Transaction) (*Transaction, *errs.AppError)
}

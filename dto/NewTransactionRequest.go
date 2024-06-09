package dto

import (
	"banking/errs"
	"strings"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (t TransactionRequest) ValidateTransactionType() *errs.AppError {
	if strings.ToLower(t.TransactionType) != "deposit" && strings.ToLower(t.TransactionType) != "withdrawal" {
		return errs.NewValidationError("The transaction type should be withdrawal or deposit")
	}
	return nil
}

func (t TransactionRequest) ValidateNegativeAmount() *errs.AppError {
	if t.Amount < 0 {
		return errs.NewValidationError("The amount must be a positive number")
	}
	return nil
}

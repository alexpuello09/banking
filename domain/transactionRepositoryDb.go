package domain

import (
	"banking/errs"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TransactionRepository struct {
	dbClient *sqlx.DB
}

func (tr TransactionRepository) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	conn := tr.dbClient
	var transId int
	var AmountAccount float64
	var sqlUpdateAccount string

	if strings.ToLower(t.TransactionType) == "withdrawal" {

		var Amount float64
		query := "SELECT amount FROM accounts WHERE account_id = ?"
		err1 := conn.Get(Amount, query, t.AccountId)
		if err1 != nil {
			return nil, errs.NewUnexpectedError("Error while validating account amount")
		}
		if Amount >= t.Amount {
			sqlUpdateAccount = "UPDATE accounts SET amount = amount - ? WHERE account_id = ? RETURNING amount"
		} else {
			return nil, errs.NewUnexpectedError("Error Account amount is too little")
		}

	}

	if strings.ToLower(t.TransactionType) == "deposit" {
		sqlUpdateAccount = "UPDATE accounts SET amount = amount + ? WHERE account_id = ? RETURNING amount"
	}

	sqlxInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES  (?,?,?,?) RETURNING transaction_id"
	err := conn.Get(transId, sqlxInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		return nil, errs.NewNotFoundError("Error while Inserting New transaction")
	}

	err2 := conn.Get(AmountAccount, sqlUpdateAccount, t.Amount, t.AccountId)
	if err2 != nil {
		return nil, errs.NewUnexpectedError("Unexpected Error while updating account")
	}
	t.TransactionId = transId
	t.Amount = AmountAccount
	return &t, nil
}

func NewTransactionDb(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{dbClient: db}
}

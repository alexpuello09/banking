package domain

import (
	"banking/errs"
	"banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (t TransactionRepositoryDb) SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError) {
	conn := t.dbClient
	var transId string
	var AmountAccount float64
	var sqlUpdateAccount string

	if strings.ToLower(transaction.TransactionType) == "withdrawal" {

		var Amount float64
		query := "SELECT amount FROM accounts WHERE account_id = ?"
		err1 := conn.Get(&Amount, query, transaction.AccountId)
		if err1 != nil {
			logger.Error("Error while getting account amount" + err1.Error())
			return nil, errs.NewUnexpectedError("Error while validating account amount")
		}
		if Amount >= transaction.Amount {
			sqlUpdateAccount = "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"
		} else {
			return nil, errs.NewUnexpectedError("Error Account amount is too little")
		}

	}

	if strings.ToLower(transaction.TransactionType) == "deposit" {
		sqlUpdateAccount = "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	}

	sqlxInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES  (?, ?, ?, ?)"
	result, err := conn.Exec(sqlxInsert, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewNotFoundError("Error while Inserting New transaction")
	}
	lastInsertedId, _ := result.LastInsertId()
	transId = strconv.FormatInt(lastInsertedId, 10)

	_, err2 := conn.Exec(sqlUpdateAccount, transaction.Amount, transaction.AccountId)
	if err2 != nil {
		logger.Error("Error while Updating Account: " + err2.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error while updating account")
	}
	selectAccountId := "SELECT amount FROM accounts WHERE account_id = ?"
	err3 := conn.Get(&AmountAccount, selectAccountId, transaction.AccountId)
	if err3 != nil {
		return nil, errs.NewUnexpectedError("Error while Selecting amount from updated account")
	}
	transaction.TransactionId = transId
	transaction.Amount = AmountAccount
	return &transaction, nil
}

func NewTransactionDb(db *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient: db}
}

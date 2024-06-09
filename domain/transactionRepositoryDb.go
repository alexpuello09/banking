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
	dbConn := t.dbClient
	var transId string
	var AmountAccount float64
	var sqlUpdateAccount string

	switch strings.ToLower(transaction.TransactionType) {
	case "withdrawal":
		var Amount float64
		query := "SELECT amount FROM accounts WHERE account_id = ?"
		err1 := dbConn.Get(&Amount, query, transaction.AccountId)
		if err1 != nil {
			logger.Error("Error while getting account amount" + err1.Error())
			return nil, errs.NewUnexpectedError("Error while validating account amount")
		}
		if Amount < transaction.Amount {
			return nil, errs.NewUnexpectedError("Error Account amount is too little")
		}
		sqlUpdateAccount = "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"

	case "deposit":
		sqlUpdateAccount = "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	}

	sqlxInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES  (?, ?, ?, ?)"
	result, err := dbConn.Exec(sqlxInsert, transaction.AccountId, transaction.Amount, transaction.TransactionType, transaction.TransactionDate)
	if err != nil {
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewNotFoundError("Error while Inserting New transaction")
	}
	lastInsertedId, _ := result.LastInsertId()
	transId = strconv.FormatInt(lastInsertedId, 10)

	_, err2 := dbConn.Exec(sqlUpdateAccount, transaction.Amount, transaction.AccountId)
	if err2 != nil {
		logger.Error("Error while Updating Account: " + err2.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error while updating account")
	}

	selectAccountId := "SELECT amount FROM accounts WHERE account_id = ?"
	err3 := dbConn.Get(&AmountAccount, selectAccountId, transaction.AccountId)
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

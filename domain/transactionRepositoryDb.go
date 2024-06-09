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
	var AccountBalance float64
	var sqlUpdateAccount string

	switch strings.ToLower(transaction.TransactionType) {
	case "withdrawal":
		var Amount float64
		query := "SELECT amount FROM accounts WHERE account_id = ?"
		_ = dbConn.Get(&Amount, query, transaction.AccountId)
		if Amount < transaction.Amount {
			return nil, errs.NewUnexpectedError("There are not enough funds in your account, check your balance and try again.")
		}
		sqlUpdateAccount = "UPDATE accounts SET amount = amount - ? WHERE account_id = ?"

	case "deposit":
		sqlUpdateAccount = "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	}

	//Insert transaction in the db
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

	//Fetching the updated amount from the account
	selectAccountId := "SELECT amount FROM accounts WHERE account_id = ?"
	err3 := dbConn.Get(&AccountBalance, selectAccountId, transaction.AccountId)
	if err3 != nil {
		return nil, errs.NewUnexpectedError("Error while Selecting amount from updated account")
	}
	transaction.TransactionId = transId
	transaction.Amount = AccountBalance
	return &transaction, nil
}

func NewTransactionDb(db *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient: db}
}

package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sqlx.Rows
	var err error
	var customers []Customer
	if status == "" {
		query := "SELECT * from customers"
		err = d.client.Select(&customers, query)
	} else {
		query := "SELECT * from customers where status = ?"
		err = d.client.Select(&customers, query, status)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Error while querying database")
		} else {
			logger.Error("unexpected database error" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer
	err := d.client.Get(&c, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USERB")
	dbPass := os.Getenv("DB_PASSB")
	dbAddr := os.Getenv("DB_ADDRB")
	dbName := os.Getenv("DB_NAMEB")
	dbPort := os.Getenv("DB_PORTB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}

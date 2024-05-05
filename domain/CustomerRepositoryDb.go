package domain

import (
	"banking/errs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	query := "SELECT * from customers "

	if strings.ToLower(status) == "active" {
		query += "where status = 1;"
	} else if strings.ToLower(status) == "inactive" {
		query += "where status = 0;"
	}

	rows, err := d.client.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Error while querying database")
		} else {
			log.Println("unexpected database error", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(query, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning customers", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}

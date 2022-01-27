package domain

import (
	"database/sql"

	"github.com/SaravananPitchaimuthu/pub-api/errs"
	"github.com/SaravananPitchaimuthu/pub-api/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

func (c CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "SELECT customer_id,name,city,zipcode,date_of_birth,status FROM customers"
		err = c.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT customer_id,name,city,zipcode,date_of_birth,status FROM customers where status = ?"
		err = c.client.Select(&customers, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying the customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database Error")
	}
	return customers, nil
}

func (c CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var customer Customer
	err := c.client.Get(&customer, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("no records found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &customer, nil
}

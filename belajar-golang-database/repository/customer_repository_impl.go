package repository

import (
	"belajar_golang_database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type customerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImpl{DB: db}
}

func (repository *customerRepositoryImpl) Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	sqlString := "insert into customer (name, email, balance, rating, married) values (?,?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx, sqlString, customer.Name, customer.Email, customer.Balance, customer.Rating, customer.Married)
	if err != nil {
		return customer, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return customer, err
	}
	customer.Id = int32(id)
	return customer, nil
}

func (repository *customerRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Customer, error) {
	sqlString := "select id, name, email, balance, rating, married from customer where id = ?"
	rows, err := repository.DB.QueryContext(ctx, sqlString, id)
	customer := entity.Customer{}
	if err != nil {
		return customer, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.Rating, &customer.Married)
		return customer, err
	} else {
		return customer, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *customerRepositoryImpl) FindAll(ctx context.Context) ([]entity.Customer, error) {
	sqlString := "select id, name, email, balance, rating, married from customer"
	rows, err := repository.DB.QueryContext(ctx, sqlString)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Balance, &customer.Rating, &customer.Married)
		customers = append(customers, customer)
	}

	return customers, err

}

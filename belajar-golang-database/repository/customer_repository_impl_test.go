package repository

import (
	"belajar_golang_database"
	"belajar_golang_database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCustomerInsert(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	customer := entity.Customer{
		Name:    "Aloha",
		Email:   "aloha@gmail.com",
		Balance: 90,
		Rating:  77,
		Married: false,
	}

	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestCustomerFindById(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang_database.GetConnection())
	customer, err := customerRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestCustomerFindAll(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang_database.GetConnection())
	customers, err := customerRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}
}

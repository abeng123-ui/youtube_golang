package repository

import (
	"belajar_golang_database"
	"belajar_golang_database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "akbar@gmail.com",
		Comment: "Test Repo",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

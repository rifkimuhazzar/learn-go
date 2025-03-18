package repository

import (
	"context"
	"fmt"
	"testing"

	gomysql "go-mysql"
	"go-mysql/entity"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetMysqlConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "react@example.com",
		Comment: "Hello react",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetMysqlConnection())
	comment, err := commentRepository.FindById(context.Background(), 120)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetMysqlConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for i, comment := range comments {
		fmt.Println(i, ":", comment)
	}
}
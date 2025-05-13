package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"go_restful_api/helper"
	"go_restful_api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func main() {
	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

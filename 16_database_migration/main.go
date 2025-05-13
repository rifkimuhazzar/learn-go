package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/helper"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
)

func main() {
	categoryRepository := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

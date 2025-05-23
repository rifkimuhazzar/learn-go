// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	categoryRepository := repository.NewCategoryRepository()
	db := app.NewDB()
	v := ProvideValidatorOptions()
	validate := validator.New(v...)
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

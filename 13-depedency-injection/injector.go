//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"

	"go_restful_api/app"
	"go_restful_api/controller"
	"go_restful_api/middleware"
	"go_restful_api/repository"
	"go_restful_api/service"
)

// var NewRouterSet = wire.NewSet(
// 	app.NewRouter,
// 	wire.Bind(new(http.Handler), new(*httprouter.Router)),
// )

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		ProvideValidatorOptions,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}

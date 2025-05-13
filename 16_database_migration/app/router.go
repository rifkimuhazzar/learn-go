package app

import (
	"github.com/julienschmidt/httprouter"

	"go_restful_api/controller"
	"go_restful_api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
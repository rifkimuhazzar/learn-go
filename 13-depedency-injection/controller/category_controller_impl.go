package controller

import (
	"go_restful_api/helper"
	"go_restful_api/models/dto"
	"go_restful_api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, 
request *http.Request, params httprouter.Params) {
	createCategoryRequest := dto.CreateCategoryRequest{}
	helper.ReadFromRequestBody(request, &createCategoryRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), 
	createCategoryRequest)

	apiResponse := dto.APIResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, 
request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	apiResponse := dto.APIResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, 
request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())

	apiResponse := dto.APIResponse{
		Code: 200,
		Status: "OK",
		Data: categoriesResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, 
request *http.Request, params httprouter.Params) {
	updateCategoryRequest := dto.UpdateCategoryRequest{}
	helper.ReadFromRequestBody(request, &updateCategoryRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	updateCategoryRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), 
	updateCategoryRequest)

	apiResponse := dto.APIResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, 
request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)

	apiResponse := dto.APIResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

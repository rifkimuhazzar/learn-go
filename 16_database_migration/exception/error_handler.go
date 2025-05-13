package exception

import (
	"go_restful_api/helper"
	"go_restful_api/models/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func notFoundError(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if !ok {
		return false
	}

	writer.WriteHeader(http.StatusNotFound)
	writer.Header().Set("Content-Type", "application/json")

	apiResponse := dto.APIResponse{
		Code: http.StatusNotFound,
		Status: "NOT FOUND",
		Data: exception.Error,
	}

	helper.WriteToResponseBody(writer, apiResponse)

	return true
}

func validationErrors(writer http.ResponseWriter, _ *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "application/json")

	apiResponse := dto.APIResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data: exception.Error(),
	}

	helper.WriteToResponseBody(writer, apiResponse)

	return true
}

func internalServerError(writer http.ResponseWriter, _ *http.Request, err any) {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Set("Content-Type", "application/json")

	apiResponse := dto.APIResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err.(error).Error(),
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

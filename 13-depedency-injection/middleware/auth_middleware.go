package middleware

import (
	"go_restful_api/helper"
	"go_restful_api/models/dto"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, 
request *http.Request) {
	if request.Header.Get("X-API-KEY") == "SECRET" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Header().Set("Content-Type", "application/json")

		apiResponse := dto.APIResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data: nil,
		}

		helper.WriteToResponseBody(writer, apiResponse)
	}
}
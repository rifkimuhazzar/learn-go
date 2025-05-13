package helper

import (
	"go_restful_api/models/domain"
	"go_restful_api/models/dto"
)

func ToCategoryResponse(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoriesResponse(categories []domain.Category) []dto.CategoryResponse {
	categoriesResponse := []dto.CategoryResponse{}
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}
	return categoriesResponse
}
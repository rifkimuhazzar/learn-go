package service

import (
	"context"
	"go_restful_api/models/dto"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CreateCategoryRequest) dto.CategoryResponse

	FindById(ctx context.Context, categoryId int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse

	Update(ctx context.Context, request dto.UpdateCategoryRequest) dto.CategoryResponse

	Delete(ctx context.Context, categoryId int)
}

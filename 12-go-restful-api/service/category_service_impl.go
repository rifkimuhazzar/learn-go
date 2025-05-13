package service

import (
	"context"
	"database/sql"
	"go_restful_api/exception"
	"go_restful_api/helper"
	"go_restful_api/models/domain"
	"go_restful_api/models/dto"
	"go_restful_api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB 								 *sql.DB
	Validate					 *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, 
db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: db,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, 
request dto.CreateCategoryRequest) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}
	category = service.CategoryRepository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, 
categoryId int) dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context,
) []dto.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoriesResponse(categories)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, 
request dto.UpdateCategoryRequest) dto.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, 
categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

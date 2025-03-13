package service

import (
	"fmt"
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err  := categoryService.Get("1")
	assert.Nil(t, category, nil)
	assert.NotNil(t, err)
	fmt.Println("--------------------")
	fmt.Println(category, err)
	fmt.Println("--------------------")
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
	fmt.Println("--------------------")
	fmt.Println(category, err)
	fmt.Println("--------------------")
}
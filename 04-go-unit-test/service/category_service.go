package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (s CategoryService) Get(id string) (*entity.Category, error) {
	category := s.Repository.FindById(id)
	if category == nil {
		return category, errors.New("category is not found")
	} else {
		return category, nil
	}
}
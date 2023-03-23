package service

import (
	"errors"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("Categpry Not Found")
	} else {
		return category, nil  
	}
}

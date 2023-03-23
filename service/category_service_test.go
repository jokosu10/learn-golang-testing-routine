package service

import (
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	cat, err := categoryService.Get("1")
	assert.Nil(t, cat)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	cat := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(cat)

	res, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Name, res.Name)
}

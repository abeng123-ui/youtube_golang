package service

import (
	"testing"
	"unit-test/entity"
	"unit-test/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	//program mock, set result yg diharapkan nanti,
	// yaitu func findById, param 1, menghasilkan Nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	//test case
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	// data yg di set sbg result
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	//program mock, set result yg diharapkan nanti,
	// yaitu func findById, param 2, menghasilkan entity category diatas
	categoryRepository.Mock.On("FindById", "2").Return(category)

	// test case
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	t.Run("category not found", func(t *testing.T) {
		categoryRepository.On("FindById", 1).Return(nil)
		category, err := categoryService.Get(1)
		assert.Nil(t, category)
		assert.NotNil(t, err)
	})

	t.Run("category found", func(t *testing.T) {
		// Here we return a Category, not its pointer,
		// because in the mock, we do Category assertion on a return object
		categoryRepository.On("FindById", 2).Return(Category{})
		category, err := categoryService.Get(2)
		assert.NotNil(t, category)
		assert.Nil(t, err)
	})
}

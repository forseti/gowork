package database

import (
	"errors"
)

type CategoryService struct {
	repository CategoryRepository
}

func (s *CategoryService) Get(id int) (*Category, error) {
	category := s.repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	}

	return category, nil
}

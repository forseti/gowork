package database

type CategoryRepository interface {
	FindById(id int) *Category
}

package category

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.Category, error)
	Search(query string) ([]*entity.Category, error)
	List() ([]*entity.Category, error)
}

type Writer interface {
	Create(e *entity.Category) (uint, error)
	Update(e *entity.Category) error
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetCategory(id uint) (*entity.Category, error)
	SearchCategories(query string) ([]*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(name string) (uint, error)
	UpdateCategory(e *entity.Category) error
	DeleteCategory(id uint) error
}

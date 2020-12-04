package product

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.Product, error)
	Search(query string) ([]*entity.Product, error)
	List() ([]*entity.Product, error)
}

type Writer interface {
	Create(e *entity.Product) (uint, error)
	Update(e *entity.Product) error
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetProduct(id uint) (*entity.Product, error)
	SearchProducts(query string) ([]*entity.Product, error)
	ListProducts() ([]*entity.Product, error)
	CreateProduct(name string) (uint, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id uint) error
}

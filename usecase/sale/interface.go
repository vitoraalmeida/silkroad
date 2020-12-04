package sale

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.Sale, error)
	Search(id uint) ([]*entity.Sale, error)
	List() ([]*entity.Sale, error)
}

type Writer interface {
	Create(e *entity.Sale) (uint, error)
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetSale(id uint) (*entity.Sale, error)
	SearchSales(customerID uint) ([]*entity.Sale, error)
	ListSales() ([]*entity.Sale, error)
	CreateSale(customerID uint, totalAmount float64) (uint, error)
	DeleteSale(id uint) error
}

package saleitem

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.SaleItem, error)
	SearchBySaleID(saleID uint) ([]*entity.SaleItem, error)
	List() ([]*entity.SaleItem, error)
}

type Writer interface {
	Create(e *entity.SaleItem) (uint, error)
	Update(e *entity.SaleItem) error
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetSaleItem(id uint) (*entity.SaleItem, error)
	ListSaleItems() ([]*entity.SaleItem, error)
	SearchItemsBySaleID(saleID uint) ([]*entity.SaleItem, error)
	CreateSaleItem(saleID uint, productID uint, quantity uint, itemAmount float64) (uint, error)
	UpdateSaleItem(e *entity.SaleItem) error
	DeleteSaleItem(id uint) error
}

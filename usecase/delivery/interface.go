package delivery

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.Delivery, error)
	Search(saleID uint) ([]*entity.Delivery, error)
	List() ([]*entity.Delivery, error)
}

type Writer interface {
	Create(e *entity.Delivery) (uint, error)
	Update(e *entity.Delivery) error
	Delete(id uint) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetDelivery(id uint) (*entity.Delivery, error)
	ListDeliveries() ([]*entity.Delivery, error)
	SearchDelivery(saleID uint) ([]*entity.Delivery, error)
	CreateDelivery(saleID uint, customerID uint, address string, status string) (uint, error)
	UpdateDelivery(e *entity.Delivery) error
	DeleteDelivery(id uint) error
}

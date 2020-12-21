package customer

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type Reader interface {
	Get(id uint) (*entity.Customer, error)
	GetByEmail(email string) (*entity.Customer, error)
	Search(query string) ([]*entity.Customer, error)
	List() ([]*entity.Customer, error)
}

type Writer interface {
	Create(e *entity.Customer) (uint, error)
	Delete(id uint) error
	Update(*entity.Customer) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetCustomer(id uint) (*entity.Customer, error)
	GetCustomerByEmail(email string) (*entity.Customer, error)
	SearchCustomers(name string) ([]*entity.Customer, error)
	ListCustomers() ([]*entity.Customer, error)
	CreateCustomer(name, email, cpf, password string) (uint, error)
	UpdateCustomer(e *entity.Customer) error
	DeleteCustomer(id uint) error
}

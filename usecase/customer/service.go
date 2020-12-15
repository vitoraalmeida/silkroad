package customer

import (
	"time"

	"github.com/vitoraalmeida/silkroad/entity"
)

var cID uint = 1

//Service customer usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateCustomer create a customer
func (s *Service) CreateCustomer(name, email, cpf, password string) (uint, error) {
	//validar email (unico), cpf, senha
	b, err := entity.NewCustomer(name, email, cpf, password)
	if err != nil {
		return b.ID, err
	}
	b.ID = cID
	cID += 1
	return s.repo.Create(b)
}

//GetCustomer get a customer
func (s *Service) GetCustomer(id uint) (*entity.Customer, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchCustomers search sales
func (s *Service) SearchCustomers(name string) ([]*entity.Customer, error) {
	sales, err := s.repo.Search(name)
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, entity.ErrNotFound
	}
	return sales, nil
}

//ListCustomers list sales
func (s *Service) ListCustomers() ([]*entity.Customer, error) {
	sales, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, entity.ErrNotFound
	}
	return sales, nil
}

//DeleteCustomer Delete a customer
func (s *Service) DeleteCustomer(id uint) error {
	_, err := s.GetCustomer(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCustomer Update a customer
func (s *Service) UpdateCustomer(e *entity.Customer) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

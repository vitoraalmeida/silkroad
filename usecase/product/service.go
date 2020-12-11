package product

import (
	"errors"
	"strings"
	"time"

	"github.com/vitoraalmeida/silkroad/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

var ErrProductOutOfStock = errors.New("Out of stock")

//CreateCategory create a book
func (s *Service) CreateProduct(name string, categoryID uint, price float64, stock uint, available bool) (uint, error) {
	b, err := entity.NewProduct(name, categoryID, price, stock, available)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetCategory get a book
func (s *Service) GetProduct(id uint) (*entity.Product, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchCategorys search products
func (s *Service) SearchProducts(query string) ([]*entity.Product, error) {
	products, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

//ListCategorys list products
func (s *Service) ListProducts() ([]*entity.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

//DeleteCategory Delete a book
func (s *Service) DeleteProduct(id uint) error {
	_, err := s.GetProduct(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCategory Update a book
func (s *Service) UpdateProduct(e *entity.Product) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

func (s *Service) DecrementProductStock(id, quantity uint) error {
	p, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	if p.Stock-quantity < 0 {
		return ErrProductOutOfStock
	}
	return s.repo.DecrementStock(id, quantity)
}

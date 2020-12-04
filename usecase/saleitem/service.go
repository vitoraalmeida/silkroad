package saleitem

import (
	"github.com/vitoraalmeida/silkroad/entity"
	"time"
)

var cID uint = 1

//Service saleitem usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateCategory create a saleitem
func (s *Service) CreateSaleItem(saleID uint, productID uint, quantity uint, itemAmount float64) (uint, error) {
	b, err := entity.NewSaleItem(saleID, productID, quantity, itemAmount)
	if err != nil {
		return b.ID, err
	}
	b.ID = cID
	cID += 1
	return s.repo.Create(b)
}

//GetCategory get a saleitem
func (s *Service) GetSaleItem(id uint) (*entity.SaleItem, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchCategorys search sales
//func (s *Service) SearchSaleItems(customerID uint) ([]*entity.SaleItem, error) {
//	sales, err := s.repo.Search(customerID)
//	if err != nil {
//		return nil, err
//	}
//	if len(sales) == 0 {
//		return nil, entity.ErrNotFound
//	}
//	return sales, nil
//}

//ListCategorys list sales
func (s *Service) ListSaleItems() ([]*entity.SaleItem, error) {
	sales, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, entity.ErrNotFound
	}
	return sales, nil
}

//DeleteCategory Delete a saleitem
func (s *Service) DeleteSaleItem(id uint) error {
	_, err := s.GetSaleItem(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCategory Update a saleitem
func (s *Service) UpdateSaleItem(e *entity.SaleItem) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

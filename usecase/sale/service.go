package sale

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

var cID uint = 1

//Service sale usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateCategory create a sale
func (s *Service) CreateSale(customerID uint, totalAmount float64) (uint, error) {
	b, err := entity.NewSale(customerID, totalAmount)
	if err != nil {
		return b.ID, err
	}
	b.ID = cID
	cID += 1
	return s.repo.Create(b)
}

//GetCategory get a sale
func (s *Service) GetSale(id uint) (*entity.Sale, error) {
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
func (s *Service) SearchSales(customerID uint) ([]*entity.Sale, error) {
	sales, err := s.repo.Search(customerID)
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, entity.ErrNotFound
	}
	return sales, nil
}

//ListCategorys list sales
func (s *Service) ListSales() ([]*entity.Sale, error) {
	sales, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(sales) == 0 {
		return nil, entity.ErrNotFound
	}
	return sales, nil
}

//DeleteCategory Delete a sale
func (s *Service) DeleteSale(id uint) error {
	_, err := s.GetSale(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCategory Update a sale
//func (s *Service) UpdateSale(e *entity.Sale) error {
//	err := e.Validate()
//	if err != nil {
//		return err
//	}
//	e.UpdatedAt = time.Now()
//	return s.repo.Update(e)
//}

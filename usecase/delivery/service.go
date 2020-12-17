package delivery

import (
	"github.com/vitoraalmeida/silkroad/entity"
	"time"
)

//Service delivery usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateDelivery create a delivery
func (s *Service) CreateDelivery(saleID uint, customerID uint, address string, status string) (uint, error) {
	d, err := entity.NewDelivery(saleID, customerID, address, status)
	if err != nil {
		return d.ID, err
	}
	return s.repo.Create(d)
}

//GetDelivery get a delivery
func (s *Service) GetDelivery(id uint) (*entity.Delivery, error) {
	d, err := s.repo.Get(id)
	if d == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return d, nil
}

//SearchSalesItems search sale items by saleID
func (s *Service) SearchDelivery(saleID uint) ([]*entity.Delivery, error) {
	deliveries, err := s.repo.Search(saleID)
	if err != nil {
		return nil, err
	}
	if len(deliveries) == 0 {
		return nil, entity.ErrNotFound
	}
	return deliveries, nil
}

//ListDeliverys list deliveries
func (s *Service) ListDeliveries() ([]*entity.Delivery, error) {
	deliveries, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(deliveries) == 0 {
		return nil, entity.ErrNotFound
	}
	return deliveries, nil
}

//DeleteDelivery Delete a delivery
func (s *Service) DeleteDelivery(id uint) error {
	_, err := s.GetDelivery(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateDelivery Update a delivery
func (s *Service) UpdateDelivery(e *entity.Delivery) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

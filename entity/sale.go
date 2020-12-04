package entity

import (
	"time"
)

type Sale struct {
	ID          uint
	CustomerID  uint
	TotalAmount float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewSale(customerID uint, totalAmount float64) (*Sale, error) {
	s := &Sale{
		CustomerID:  customerID,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now(),
	}
	err := s.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return s, nil
}

func (s *Sale) Validate() error {
	if s.CustomerID <= 0 || s.TotalAmount <= 0.00 {
		return ErrInvalidEntity
	}
	return nil
}

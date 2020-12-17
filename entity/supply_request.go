package entity

import (
	"time"
)

type SupplyRequest struct {
	ID        uint
	ProductID uint
	Quantity  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSupplyRequest(productID uint, quantity uint) (*SupplyRequest, error) {
	sr := &SupplyRequest{
		ProductID: productID,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}
	err := sr.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return sr, nil
}

func (sr *SupplyRequest) Validate() error {
	if sr.ProductID <= 0 || sr.Quantity <= 0 {
		return ErrInvalidEntity
	}
	return nil
}

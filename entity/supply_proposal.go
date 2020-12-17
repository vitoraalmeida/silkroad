package entity

import (
	"time"
)

type SupplyProposal struct {
	ID           uint
	RequestID    uint
	Quantity     uint
	UnitaryPrice float64
	TotalAmount  float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewSupplyProposal(requestID uint, quantity uint, unitaryPrice float64, totalAmount float64) (*SupplyProposal, error) {
	sp := &SupplyProposal{
		RequestID:    requestID,
		Quantity:     quantity,
		UnitaryPrice: unitaryPrice,
		TotalAmount:  totalAmount,
		CreatedAt:    time.Now(),
	}
	err := sp.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return sp, nil
}

func (sp *SupplyProposal) Validate() error {
	if sp.RequestID == 0 || sp.Quantity == 0 || sp.UnitaryPrice == 0.00 || sp.TotalAmount == 0.00 {
		return ErrInvalidEntity
	}
	return nil
}

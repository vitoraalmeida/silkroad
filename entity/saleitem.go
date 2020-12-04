package entity

import (
	"time"
)

type SaleItem struct {
	ID         uint
	SaleID     uint
	ProductID  uint
	Quantity   uint
	ItemAmount float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewSaleItem(saleID, productID, quantity uint, itemAmount float64) (*SaleItem, error) {
	si := &SaleItem{
		SaleID:     saleID,
		ProductID:  productID,
		Quantity:   quantity,
		ItemAmount: itemAmount,
		CreatedAt:  time.Now(),
	}
	err := si.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return si, nil
}

func (si *SaleItem) Validate() error {
	if si.SaleID <= 0 || si.ProductID <= 0 || si.Quantity <= 0 {
		return ErrInvalidEntity
	}
	return nil
}

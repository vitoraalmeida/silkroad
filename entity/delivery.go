package entity

import (
	"time"
)

type Delivery struct {
	ID         uint
	SaleID     uint
	CustomerID uint
	Address    string
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewDelivery(saleID uint, customerID uint, address string, status string) (*Delivery, error) {
	d := &Delivery{
		SaleID:     saleID,
		CustomerID: customerID,
		Address:    address,
		Status:     status,
		CreatedAt:  time.Now(),
	}
	err := d.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return d, nil
}

func (d *Delivery) Validate() error {
	if d.SaleID == 0 || d.CustomerID == 0 || d.Address == "" || d.Status == "" {
		return ErrInvalidEntity
	}
	return nil
}

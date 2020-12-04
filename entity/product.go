package entity

import (
	"time"
)

var cID uint = 1

type Product struct {
	ID         uint
	Name       string
	CategoryID uint
	Price      float64
	Stock      uint
	Available  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewProduct(name string, categoryID uint, price float64, stock uint, available bool) (*Product, error) {
	p := &Product{
		ID:         cID,
		Name:       name,
		CategoryID: categoryID,
		Available:  available,
		Price:      price,
		Stock:      stock,
		CreatedAt:  time.Now(),
	}
	cID = cID + 1
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" || p.CategoryID <= 0 || p.Price <= 0.00 || p.Stock <= 0 {
		return ErrInvalidEntity
	}
	if p.Stock == 0 && p.Available == true {
		return ErrInvalidEntity
	}
	return nil
}

package entity

import (
	"time"
)

var createID uint = 1

type Category struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) (*Category, error) {
	c := &Category{
		ID:        createID,
		Name:      name,
		CreatedAt: time.Now(),
	}
	createID = createID + 1
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}

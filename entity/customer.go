package entity

import (
	"time"
)

type Customer struct {
	ID        uint
	Name      string
	Email     string
	CPF       string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name, email, cpf, password string) (*Customer, error) {
	c := &Customer{
		Name:      name,
		Email:     email,
		CPF:       cpf,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

// adiciona um pedido
// busca pedidos

func (c *Customer) Validate() error {
	if c.Name == "" || c.Email == "" || c.CPF == "" || c.Password == "" {
		return ErrInvalidEntity
	}
	return nil
}

// valida senha
// gera hash senha

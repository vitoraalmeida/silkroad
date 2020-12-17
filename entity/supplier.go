package entity

import (
	"time"
)

type Supplier struct {
	ID        uint
	Name      string
	Email     string
	CNPJ      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSupplier(name, email, cnpj, password string) (*Supplier, error) {
	s := &Supplier{
		Name:      name,
		Email:     email,
		CNPJ:      cnpj,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := s.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return s, nil
}

func (s *Supplier) Validate() error {
	if s.Name == "" || s.Email == "" || s.CNPJ == "" || s.Password == "" {
		return ErrInvalidEntity
	}
	return nil
}

// valida senha
// gera hash senha

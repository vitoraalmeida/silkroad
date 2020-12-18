package entity

type Admin struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func NewAdmin(name, email, password string) (*Admin, error) {
	a := &Admin{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := a.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return a, nil
}

func (a *Admin) Validate() error {
	if a.Name == "" || a.Email == "" || a.Password == "" {
		return ErrInvalidEntity
	}
	return nil
}

//validar senha
//gerar hash de senha

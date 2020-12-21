package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type CustomerPSQL struct {
	db *sql.DB
}

func NewCustomerPQSL(db *sql.DB) *CustomerPSQL {
	return &CustomerPSQL{
		db: db,
	}
}

func (cs *CustomerPSQL) Create(c *entity.Customer) (uint, error) {
	var id uint
	err := cs.db.QueryRow(`insert into customer (name, email, cpf, password, 
	created_at) values ($1,$2,$3,$4,$5) RETURNING id`,
		c.Name, c.Email, c.CPF, c.Password, c.CreatedAt,
	).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create customer psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (cs *CustomerPSQL) List() ([]*entity.Customer, error) {
	stmt, err := cs.db.Prepare(`select id, name, email, cpf, password, 
	created_at, updated_at from customer`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var customers []*entity.Customer
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Customer
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.CPF,
			&c.Password,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("List customers psql: %v", err)
		}
		customers = append(customers, &c)

	}
	if len(customers) == 0 {
		return nil, fmt.Errorf("List customers psql: %v", "customer table is empty")
	}
	return customers, nil
}

func (cs *CustomerPSQL) Get(id uint) (*entity.Customer, error) {
	stmt, err := cs.db.Prepare(`select id, name, email, cpf, password, 
	created_at, updated_at from customer where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var c entity.Customer
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.CPF,
			&c.Password,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Get customer psql: %v", err)
		}

	}

	// check if any customer was found
	if c.ID == 0 {
		return nil, fmt.Errorf("Get customer psql: customer not found")
	}

	return &c, nil
}

func (cs *CustomerPSQL) GetByEmail(email string) (*entity.Customer, error) {
	stmt, err := cs.db.Prepare(`select id, name, email, cpf, password, 
	created_at, updated_at from customer where email = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var c entity.Customer
	rows, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.CPF,
			&c.Password,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Get customer psql: %v", err)
		}

	}

	// check if any customer was found
	if c.ID == 0 {
		return nil, fmt.Errorf("Get customer psql: customer not found")
	}

	return &c, nil
}

func (cs *CustomerPSQL) Search(query string) ([]*entity.Customer, error) {
	stmt, err := cs.db.Prepare(`select id, name, email, cpf, password, 
	created_at, updated_at from customer where name like $1`)
	if err != nil {
		return nil, err
	}
	var customers []*entity.Customer
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Customer
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.CPF,
			&c.Password,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Search customer psql: %v", err)
		}
		customers = append(customers, &c)
	}

	return customers, nil
}

func (cs *CustomerPSQL) Update(e *entity.Customer) error {
	_, err := cs.db.Exec(`update customer set name = $1, email = $2, cpf = $3, 
	password = $4 where id = $5`, e.Name, e.Email, e.CPF, e.Password, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CustomerPSQL) Delete(id uint) error {
	_, err := cs.db.Exec("delete from customer where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

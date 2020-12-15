package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type SalePSQL struct {
	db *sql.DB
}

func NewSalePQSL(db *sql.DB) *SalePSQL {
	return &SalePSQL{
		db: db,
	}
}

func (sr *SalePSQL) Create(s *entity.Sale) (uint, error) {
	var id uint
	err := sr.db.QueryRow(`
		insert into sale (customer_id, total_amount) values($1,$2) RETURNING id`,
		s.CustomerID,
		s.TotalAmount,
	).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create sale psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (sr *SalePSQL) List() ([]*entity.Sale, error) {
	stmt, err := sr.db.Prepare(`select id, customer_id, total_amount, created_at, updated_at from sale`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var sales []*entity.Sale
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var s entity.Sale
		err := rows.Scan(&s.ID, &s.CustomerID, &s.TotalAmount, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("List sales psql: %v", err)
		}
		sales = append(sales, &s)

	}
	if len(sales) == 0 {
		return nil, fmt.Errorf("List sales psql: %v", "sale table is empty")
	}
	return sales, nil
}

func (sr *SalePSQL) Get(id uint) (*entity.Sale, error) {
	stmt, err := sr.db.Prepare(`select id, customer_id, total_amount, created_at, 
	updated_at from sale where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var s entity.Sale
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&s.ID, &s.CustomerID, &s.TotalAmount, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Get sale psql: %v", err)
		}
	}
	return &s, nil
}

func (sr *SalePSQL) Search(customerID uint) ([]*entity.Sale, error) {
	stmt, err := sr.db.Prepare(`select id, customer_id, total_amount, created_at, 
	updated_at from sale where customer_id = $1`)
	if err != nil {
		return nil, err
	}
	var sales []*entity.Sale
	rows, err := stmt.Query(customerID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var s entity.Sale
		err = rows.Scan(&s.ID, &s.CustomerID, &s.TotalAmount, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &s)
	}

	return sales, nil
}

func (sr *SalePSQL) Delete(id uint) error {
	_, err := sr.db.Exec("delete from sale where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

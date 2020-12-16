package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type SaleItemPSQL struct {
	db *sql.DB
}

func NewSaleItemPQSL(db *sql.DB) *SaleItemPSQL {
	return &SaleItemPSQL{
		db: db,
	}
}

func (sr *SaleItemPSQL) Create(s *entity.SaleItem) (uint, error) {
	var id uint
	err := sr.db.QueryRow(`
		insert into sale_item (sale_id, product_id, quantity, item_amount) 
		values($1,$2,$3,$4) RETURNING id`,
		s.SaleID,
		s.ProductID,
		s.Quantity,
		s.ItemAmount,
	).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create sale_item psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (sr *SaleItemPSQL) List() ([]*entity.SaleItem, error) {
	stmt, err := sr.db.Prepare(`select id, sale_id, product_id, quantity, item_amount, 
	created_at, updated_at from sale_item`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var saleItems []*entity.SaleItem
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var s entity.SaleItem
		err := rows.Scan(
			&s.ID,
			&s.SaleID,
			&s.ProductID,
			&s.Quantity,
			&s.ItemAmount,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("List saleItems psql: %v", err)
		}
		saleItems = append(saleItems, &s)

	}
	if len(saleItems) == 0 {
		return nil, fmt.Errorf("List saleItems psql: %v", "sale_item table is empty")
	}
	return saleItems, nil
}

func (sr *SaleItemPSQL) Get(id uint) (*entity.SaleItem, error) {
	stmt, err := sr.db.Prepare(`select id, sale_id, product_id, quantity, item_amount, 
	created_at, updated_at from sale_item where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var s entity.SaleItem
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&s.ID,
			&s.SaleID,
			&s.ProductID,
			&s.Quantity,
			&s.ItemAmount,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Get sale_item psql: %v", err)
		}
	}
	return &s, nil
}

func (sr *SaleItemPSQL) Search(customerID uint) ([]*entity.SaleItem, error) {
	stmt, err := sr.db.Prepare(`select id, sale_id, product_id, quantity, item_amount, 
	created_at, updated_at from sale_item where sale_id = $1`)
	if err != nil {
		return nil, err
	}
	var saleItems []*entity.SaleItem
	rows, err := stmt.Query(customerID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var s entity.SaleItem
		err := rows.Scan(
			&s.ID,
			&s.SaleID,
			&s.ProductID,
			&s.Quantity,
			&s.ItemAmount,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		saleItems = append(saleItems, &s)
	}

	return saleItems, nil
}

func (sr *SaleItemPSQL) Update(e *entity.SaleItem) error {
	_, err := sr.db.Exec(`update product set sale_id = $1, product_id = $2, 
	quantity = $3, item_amount = $4 where id = $5`,
		e.SaleID, e.ProductID, e.Quantity, e.ItemAmount, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SaleItemPSQL) Delete(id uint) error {
	_, err := sr.db.Exec("delete from sale_item where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

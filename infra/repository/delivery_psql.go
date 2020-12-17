package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type DeliveryPSQL struct {
	db *sql.DB
}

func NewDeliveryPSQL(db *sql.DB) *DeliveryPSQL {
	return &DeliveryPSQL{
		db: db,
	}
}

func (d *DeliveryPSQL) Create(e *entity.Delivery) (uint, error) {
	var id uint
	err := d.db.QueryRow(`
		insert into delivery (sale_id, customer_id, address, status) 
		values($1,$2,$3,$4) RETURNING id`,
		e.SaleID,
		e.CustomerID,
		e.Address,
		e.Status,
	).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create delivery psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (d *DeliveryPSQL) List() ([]*entity.Delivery, error) {
	stmt, err := d.db.Prepare(`select id, sale_id, customer_id, address, status, 
	created_at, updated_at from delivery`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var deliveries []*entity.Delivery
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e entity.Delivery
		err := rows.Scan(
			&e.ID,
			&e.SaleID,
			&e.CustomerID,
			&e.Address,
			&e.Status,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("List deliveries psql: %v", err)
		}
		deliveries = append(deliveries, &e)

	}
	if len(deliveries) == 0 {
		return nil, fmt.Errorf("List deliveries psql: %v", "delivery table is empty")
	}
	return deliveries, nil
}

func (d *DeliveryPSQL) Get(id uint) (*entity.Delivery, error) {
	stmt, err := d.db.Prepare(`select id, sale_id, customer_id, address, status, 
	created_at, updated_at from delivery where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var e entity.Delivery
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&e.ID,
			&e.SaleID,
			&e.CustomerID,
			&e.Address,
			&e.Status,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Get delivery psql: %v", err)
		}
	}
	return &e, nil
}

func (d *DeliveryPSQL) Search(saleID uint) ([]*entity.Delivery, error) {
	stmt, err := d.db.Prepare(`select id, sale_id, customer_id, address, status, 
	created_at, updated_at from delivery where sale_id = $1`)
	if err != nil {
		return nil, err
	}
	var deliveries []*entity.Delivery
	rows, err := stmt.Query(saleID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e entity.Delivery
		err := rows.Scan(
			&e.ID,
			&e.SaleID,
			&e.CustomerID,
			&e.Address,
			&e.Status,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		deliveries = append(deliveries, &e)
	}

	return deliveries, nil
}

func (d *DeliveryPSQL) Update(e *entity.Delivery) error {
	_, err := d.db.Exec(`update delivery set sale_id = $1, customer_id = $2, 
	address = $3, status = $4 where id = $5`,
		e.SaleID, e.CustomerID, e.Address, e.Status, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *DeliveryPSQL) Delete(id uint) error {
	_, err := d.db.Exec("delete from delivery where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

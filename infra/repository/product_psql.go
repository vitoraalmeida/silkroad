package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type ProductPSQL struct {
	db *sql.DB
}

func NewProductPQSL(db *sql.DB) *ProductPSQL {
	return &ProductPSQL{
		db: db,
	}
}

func (pr *ProductPSQL) Create(p *entity.Product) (uint, error) {
	var id uint
	err := pr.db.QueryRow(`
		insert into product (name, category_id, price, stock, available)
		values($1,$2,$3,$4,$5) RETURNING id`,
		p.Name, p.CategoryID, p.Price, p.Stock, p.Available).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create product psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (pr *ProductPSQL) List() ([]*entity.Product, error) {
	stmt, err := pr.db.Prepare(`select id, name, category_id, price, stock,
	available, created_at, updated_at from product`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var products []*entity.Product
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.CategoryID,
			&p.Price,
			&p.Stock,
			&p.Available,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("List products psql: %v", err)
		}
		products = append(products, &p)

	}
	if len(products) == 0 {
		return nil, fmt.Errorf("List products psql: %v", "product table is empty")
	}
	return products, nil
}

func (pr *ProductPSQL) Get(id uint) (*entity.Product, error) {
	stmt, err := pr.db.Prepare(`select id, name, category_id, price, stock, 
	available, created_at, updated_at from product where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p entity.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.CategoryID,
			&p.Price,
			&p.Stock,
			&p.Available,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Create product psql: %v", err)
		}
	}
	return &p, nil
}

func (pr *ProductPSQL) Search(query string) ([]*entity.Product, error) {
	stmt, err := pr.db.Prepare(`select id, name, category_id, price, stock, 
	available, created_at, updated_at from product where name like $1`)
	if err != nil {
		return nil, err
	}
	var products []*entity.Product
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price, &p.Stock, &p.Available, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}

func (pr *ProductPSQL) Update(e *entity.Product) error {
	_, err := pr.db.Exec("update product set name = $1, category_id = $2, price = $3, stock = $4, available = $5 where id = $6", e.Name, e.CategoryID, e.Price, e.Stock, e.Available, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductPSQL) Delete(id uint) error {
	_, err := pr.db.Exec("delete from product where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProductPSQL) DecrementStock(id, quantity uint) error {
	_, err := pr.db.Exec("update product set stock = stock - $1 where id = $2", id, quantity)
	if err != nil {
		return err
	}
	return nil
}

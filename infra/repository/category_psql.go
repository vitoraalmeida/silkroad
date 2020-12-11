package repository

import (
	"database/sql"
	"fmt"

	"github.com/vitoraalmeida/silkroad/entity"
)

type CategoryPSQL struct {
	db *sql.DB
}

func NewCategoryPQSL(db *sql.DB) *CategoryPSQL {
	return &CategoryPSQL{
		db: db,
	}
}

func (cs *CategoryPSQL) Create(c *entity.Category) (uint, error) {
	var id uint
	err := cs.db.QueryRow(`
		insert into category (name, created_at)
		values($1,$2) RETURNING id`,
		c.Name,
		c.CreatedAt).Scan(&id)

	if err != nil {
		if id == 0 {
			return 0, fmt.Errorf("Create category psql: %v", err)
		}
		return 0, err
	}

	return id, nil
}

func (cs *CategoryPSQL) List() ([]*entity.Category, error) {
	stmt, err := cs.db.Prepare(`select id, name, created_at, updated_at from category`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var categories []*entity.Category
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("List categories psql: %v", err)
		}
		categories = append(categories, &c)

	}
	if len(categories) == 0 {
		return nil, fmt.Errorf("List categories psql: %v", "category table is empty")
	}
	return categories, nil
}

func (cs *CategoryPSQL) Get(id uint) (*entity.Category, error) {
	stmt, err := cs.db.Prepare(`select id, name, created_at, updated_at from category where id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var c entity.Category
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Create category psql: %v", err)
		}
	}
	return &c, nil
}

func (cs *CategoryPSQL) Search(query string) ([]*entity.Category, error) {
	stmt, err := cs.db.Prepare(`select id, name, created_at, updated_at from category where name like $1`)
	if err != nil {
		return nil, err
	}
	var categories []*entity.Category
	rows, err := stmt.Query("%" + query[1:] + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c entity.Category
		err = rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}

	return categories, nil
}

func (cs *CategoryPSQL) Update(e *entity.Category) error {
	_, err := cs.db.Exec("update category set name = $1 where id = $2", e.Name, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CategoryPSQL) Delete(id uint) error {
	_, err := cs.db.Exec("delete from category where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

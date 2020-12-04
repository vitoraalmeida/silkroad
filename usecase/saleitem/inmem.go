package saleitem

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.SaleItem
}

//newInmem cre,te new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.SaleItem{}
	return &inmem{
		m: m,
	}
}

//Create a saleitem
func (r *inmem) Create(e *entity.SaleItem) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a saleitem
func (r *inmem) Get(id uint) (*entity.SaleItem, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a saleitem
func (r *inmem) Update(e *entity.SaleItem) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search saleitem
//func (r *inmem) Search(id uint) ([]*entity.SaleItem, error) {
//	var d []*entity.SaleItem
//	for _, j := range r.m {
//		if id == j.ID {
//			d = append(d, j)
//		}
//	}
//	return d, nil
//}

//List saleitem
func (r *inmem) List() ([]*entity.SaleItem, error) {
	var d []*entity.SaleItem
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a saleitem
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

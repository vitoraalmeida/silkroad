package sale

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.Sale
}

//newInmem cre,te new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.Sale{}
	return &inmem{
		m: m,
	}
}

//Create a sale
func (r *inmem) Create(e *entity.Sale) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a sale
func (r *inmem) Get(id uint) (*entity.Sale, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a sale
func (r *inmem) Update(e *entity.Sale) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search products
func (r *inmem) Search(id uint) ([]*entity.Sale, error) {
	var d []*entity.Sale
	for _, j := range r.m {
		if id == j.ID {
			d = append(d, j)
		}
	}
	return d, nil
}

//List products
func (r *inmem) List() ([]*entity.Sale, error) {
	var d []*entity.Sale
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a sale
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

package product

import (
	"strings"

	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.Product
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.Product{}
	return &inmem{
		m: m,
	}
}

//Create a product
func (r *inmem) Create(e *entity.Product) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a product
func (r *inmem) Get(id uint) (*entity.Product, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a product
func (r *inmem) Update(e *entity.Product) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search products
func (r *inmem) Search(query string) ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List products
func (r *inmem) List() ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a product
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

func (r *inmem) DecrementStock(id, quantity uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id].Stock -= quantity
	return nil
}

package customer

import (
	"strings"

	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.Customer
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.Customer{}
	return &inmem{
		m: m,
	}
}

//Create a customer
func (r *inmem) Create(e *entity.Customer) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a customer
func (r *inmem) Get(id uint) (*entity.Customer, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a customer
func (r *inmem) Update(e *entity.Customer) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search customers
func (r *inmem) Search(query string) ([]*entity.Customer, error) {
	var d []*entity.Customer
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List customers
func (r *inmem) List() ([]*entity.Customer, error) {
	var d []*entity.Customer
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a customer
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

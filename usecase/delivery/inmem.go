package delivery

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.Delivery
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.Delivery{}
	return &inmem{
		m: m,
	}
}

//Create a delivery
func (r *inmem) Create(e *entity.Delivery) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a delivery
func (r *inmem) Get(id uint) (*entity.Delivery, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a delivery
func (r *inmem) Update(e *entity.Delivery) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search delivery by saleID
func (r *inmem) SearchBySaleID(saleID uint) ([]*entity.Delivery, error) {
	var d []*entity.Delivery
	for _, j := range r.m {
		if saleID == j.SaleID {
			d = append(d, j)
		}
	}
	return d, nil
}

//List delivery
func (r *inmem) List() ([]*entity.Delivery, error) {
	var d []*entity.Delivery
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a delivery
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

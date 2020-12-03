package category

import (
	"strings"

	"github.com/vitoraalmeida/silkroad/entity"
)

//inmem in memory repo
type inmem struct {
	m map[uint]*entity.Category
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[uint]*entity.Category{}
	return &inmem{
		m: m,
	}
}

//Create a category
func (r *inmem) Create(e *entity.Category) (uint, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a category
func (r *inmem) Get(id uint) (*entity.Category, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update a category
func (r *inmem) Update(e *entity.Category) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search books
func (r *inmem) Search(query string) ([]*entity.Category, error) {
	var d []*entity.Category
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List books
func (r *inmem) List() ([]*entity.Category, error) {
	var d []*entity.Category
	for _, j := range r.m {
		if j == nil {
			continue
		}
		d = append(d, j)
	}
	return d, nil
}

//Delete a category
func (r *inmem) Delete(id uint) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

package category

import (
	"fmt"
	"github.com/vitoraalmeida/silkroad/entity"
	"strings"
	"time"
)

//Service category usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateCategory create a category
func (s *Service) CreateCategory(name string) (uint, error) {
	c, err := entity.NewCategory(name)
	if err != nil {
		return c.ID, err
	}
	return s.repo.Create(c)
}

//GetCategory get a category
func (s *Service) GetCategory(id uint) (*entity.Category, error) {
	c, err := s.repo.Get(id)
	if c == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

//SearchCategorys search categories
func (s *Service) SearchCategories(query string) ([]*entity.Category, error) {
	fmt.Println("query: ", query)
	categories, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, entity.ErrNotFound
	}
	return categories, nil
}

//ListCategorys list categories
func (s *Service) ListCategories() ([]*entity.Category, error) {
	categories, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, entity.ErrNotFound
	}
	return categories, nil
}

//DeleteCategory Delete a category
func (s *Service) DeleteCategory(id uint) error {
	_, err := s.GetCategory(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCategory Update a category
func (s *Service) UpdateCategory(e *entity.Category) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

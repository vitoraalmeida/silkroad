package category

import (
	"strings"
	"time"

	"github.com/vitoraalmeida/silkroad/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateCategory create a book
func (s *Service) CreateCategory(name string) (uint, error) {
	b, err := entity.NewCategory(name)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetCategory get a book
func (s *Service) GetCategory(id uint) (*entity.Category, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchCategorys search categories
func (s *Service) SearchCategories(query string) ([]*entity.Category, error) {
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

//DeleteCategory Delete a book
func (s *Service) DeleteCategory(id uint) error {
	_, err := s.GetCategory(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateCategory Update a book
func (s *Service) UpdateCategory(e *entity.Category) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

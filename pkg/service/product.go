package service

import (
	"github.com/ks6088ts/rest-go/pkg/entity"
	"github.com/ks6088ts/rest-go/pkg/repository"
)

// Service ...
type Service struct {
	session *repository.Session
}

// NewService creates a service
func NewService(session *repository.Session) (*Service, error) {
	service := &Service{
		session: session,
	}
	service.MigrateProduct()
	return service, nil
}

// CreateProduct ...
func (s *Service) CreateProduct(p *entity.Product) (*entity.Product, error) {
	if err := s.session.CreateProduct(p); err != nil {
		return nil, err
	}

	return p, nil
}

// ReadProduct ...
func (s *Service) ReadProduct(id string) (*entity.Product, error) {
	p, err := s.session.ReadProduct(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ReadProducts ...
func (s *Service) ReadProducts() ([]entity.Product, error) {
	p, err := s.session.ReadProducts()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// MigrateProduct ...
func (s *Service) MigrateProduct() {
	s.session.MigrateProduct()
}

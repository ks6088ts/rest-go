package service

import (
	"github.com/gin-gonic/gin"
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
func (s *Service) CreateProduct(c *gin.Context) (*entity.Product, error) {
	var p entity.Product

	if err := c.BindJSON(&p); err != nil {
		return nil, err
	}

	if err := s.session.Db.Create(&p).Error; err != nil {
		return nil, err
	}

	return &p, nil
}

// GetProduct ...
func (s *Service) GetProduct(id string) (*entity.Product, error) {
	var p entity.Product
	if err := s.session.Db.Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// MigrateProduct ...
func (s *Service) MigrateProduct() {
	s.session.Db.AutoMigrate(&entity.Product{})
}
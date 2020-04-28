package repository

import (
	"github.com/ks6088ts/rest-go/pkg/entity"
)

// CreateProduct ...
func (s *Session) CreateProduct(p *entity.Product) error {
	if err := s.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

// GetProduct ...
func (s *Session) GetProduct(id string) (*entity.Product, error) {
	var p entity.Product
	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// ReadProducts ...
func (s *Session) ReadProducts() ([]entity.Product, error) {
	products := []entity.Product{}
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// MigrateProduct ...
func (s *Session) MigrateProduct() {
	s.db.AutoMigrate(&entity.Product{})
}

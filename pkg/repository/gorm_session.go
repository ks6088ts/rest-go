package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/ks6088ts/rest-go/pkg/entity"
)

// NewGormSession creates a GormSession
func NewGormSession(dbms, connect string) (*GormSession, error) {
	gormDB, err := gorm.Open(dbms, connect)
	if err != nil {
		return nil, err
	}

	return &GormSession{
		db: gormDB,
	}, nil
}

// Close GormSession
func (s *GormSession) Close() error {
	return s.db.Close()
}

// GormSession is a type definition for GormSession
type GormSession struct {
	db *gorm.DB
}

// CreateProduct ...
func (s *GormSession) CreateProduct(p *entity.Product) error {
	if err := s.db.Create(p).Error; err != nil {
		return err
	}
	return nil
}

// ReadProduct ...
func (s *GormSession) ReadProduct(id string) (*entity.Product, error) {
	var p entity.Product
	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// ReadProducts ...
func (s *GormSession) ReadProducts() ([]entity.Product, error) {
	products := []entity.Product{}
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// MigrateProduct ...
func (s *GormSession) MigrateProduct() {
	s.db.AutoMigrate(&entity.Product{})
}

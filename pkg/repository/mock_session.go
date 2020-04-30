package repository

import (
	"strconv"

	"github.com/ks6088ts/rest-go/pkg/entity"
)

// MockSession for test
type MockSession struct {
	db     map[int64]entity.Product
	lastID int64
}

// NewMockSession creates a MockSession
func NewMockSession() (*MockSession, error) {
	return &MockSession{
		db:     make(map[int64]entity.Product),
		lastID: 0,
	}, nil
}

// Close ...
func (s *MockSession) Close() error {
	return nil
}

// CreateProduct ...
func (s *MockSession) CreateProduct(p *entity.Product) error {
	s.db[s.lastID] = *p
	s.lastID++
	return nil
}

// ReadProduct ...
func (s *MockSession) ReadProduct(id string) (*entity.Product, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	p := s.db[idInt]
	return &p, nil
}

// ReadProducts ...
func (s *MockSession) ReadProducts() ([]entity.Product, error) {
	var products []entity.Product
	for i := int64(0); i < s.lastID; i++ {
		products = append(products, s.db[i])
	}
	return products, nil
}

// MigrateProduct ...
func (s *MockSession) MigrateProduct() {
}

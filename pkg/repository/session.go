package repository

import (

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ks6088ts/rest-go/pkg/entity"
)

// Session ...
type Session interface {
	Close() error
	CreateProduct(p *entity.Product) error
	ReadProduct(id string) (*entity.Product, error)
	ReadProducts() ([]entity.Product, error)
	MigrateProduct()
}

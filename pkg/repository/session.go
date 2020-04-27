package repository

import (
	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewSession creates a session
func NewSession(dbms, connect string) (*Session, error) {
	gormDB, err := gorm.Open(dbms, connect)
	if err != nil {
		return nil, err
	}

	return &Session{
		Db: gormDB,
	}, nil
}

// Close session
func (s *Session) Close() error {
	return s.Db.Close()
}

// Session is a type definition for session
type Session struct {
	// TODO: 非公開化
	Db *gorm.DB
}

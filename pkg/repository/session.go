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
		db: gormDB,
	}, nil
}

// Close session
func (s *Session) Close() {
	s.db.Close()
}

// Session is a type definition for session
type Session struct {
	db *gorm.DB
}

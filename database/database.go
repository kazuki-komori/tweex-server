package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewDB() (*SqlHandler, error) {

	db, err := gorm.Open("postgres", "postgres://admin:pass@db:5432/db?sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("failed to connect db= %w", err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler, nil
}

package datastore

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgres driver
)

// NewPostgres creates a connection with Postgres database
func NewPostgres() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_ADDRS"))
	if err != nil {
		return nil, err
	}

	db.Debug()
	db.LogMode(true)

	return db, nil
}

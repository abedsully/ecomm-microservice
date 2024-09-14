package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/ecomm?parseTime=true")

	if err != nil {
		return nil, fmt.Errorf("Error opening database : %w", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
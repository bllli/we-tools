package db

import (
	"github.com/jmoiron/sqlx"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(driver, uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(driver, uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

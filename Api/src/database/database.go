package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Database Connect
func DatabaseConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DATABASE_FULL_ADD)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

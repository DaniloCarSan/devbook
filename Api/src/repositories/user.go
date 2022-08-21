package repositories

import (
	"api/src/models"
	"database/sql"
)

type user struct {
	db *sql.DB
}

// Instancia um novo repository
func UserRepository(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

// Create
func (u *user) Create(user models.User) (uint64, error) {

	return 0, nil
}

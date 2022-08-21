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
func (repository *user) Create(user models.User) (uint64, error) {

	stmt, err := repository.db.Prepare(
		"INSERT INTO users (name,nickname,email,password)VALUES(?,?,?,?)",
	)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nickane, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

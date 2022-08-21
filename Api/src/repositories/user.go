package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository *user) All(nameOrNickname string) ([]models.User, error) {

	var rows *sql.Rows
	var err error

	if nameOrNickname != "" {

		nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

		rows, err = repository.db.Query(
			"SELECT id,name,nickname,email,createAt FROM users WHERE name LIKE ? OR nickname LIKE ?",
			nameOrNickname, nameOrNickname,
		)

	} else {

		rows, err = repository.db.Query("SELECT id,name,nickname,email,createAt FROM users")

	}

	if err != nil {
		return nil, err
	}

	var users []models.User

	var user models.User

	for rows.Next() {

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickane,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

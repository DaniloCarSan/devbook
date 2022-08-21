package models

import (
	"errors"
	"net/mail"
	"strings"
	"time"
)

// User
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nickane  string    `json:"nickane,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (u *User) validate(step string) error {

	if u.Name == "" {
		return errors.New("field name required")
	}

	if u.Nickane == "" {
		return errors.New("field nickane required")
	}

	if step == "create" && u.Email == "" {
		return errors.New("field email required")
	}

	if step == "create" {
		_, err := mail.ParseAddress(u.Email)

		if err != nil {
			return errors.New("email invalid")
		}
	}

	if step == "create" && u.Password == "" {
		return errors.New("field password required")
	}

	if step == "create" && len(u.Password) < 6 {
		return errors.New("require  > 6 ")
	}

	return nil
}

// Validar e preparar os dados de entrada
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	u.format()

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickane = strings.TrimSpace(u.Nickane)
	u.Email = strings.TrimSpace(u.Email)
}

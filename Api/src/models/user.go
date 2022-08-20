package models

import "time"

// User
type User struct {
	ID       uint      `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nickane  string    `json:"nickane,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

package security

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordToHash
func PasswordToHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CompareHashWithPassword
func CompareHashWithPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

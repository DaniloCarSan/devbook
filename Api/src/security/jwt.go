package security

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Generate token jwt
func GenerateJwtToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secret"))
}

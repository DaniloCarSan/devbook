package security

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	return token.SignedString([]byte(config.JWT_SECRET_KEY))
}

// Validate jwt token
func ValidateJwtToken(r *http.Request) error {

	tokenString := extractJwtToken(r)

	token, err := jwt.Parse(tokenString, funcSecretJwtKeyVerifySignature)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalid")
}

func extractJwtToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func funcSecretJwtKeyVerifySignature(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method of signature unexpected  %v", token.Header["alg"])
	}

	return config.JWT_SECRET_KEY, nil
}

func ExtractUserIdOfJwtToken(r *http.Request) (uint64, error) {

	tokenString := extractJwtToken(r)

	token, err := jwt.Parse(tokenString, funcSecretJwtKeyVerifySignature)

	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, errors.New("token invalid")
}

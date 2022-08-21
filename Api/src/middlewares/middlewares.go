package middlewares

import (
	"api/src/responses"
	"api/src/security"
	"net/http"
)

// Verify user authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := security.ValidateJwtToken(r); err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

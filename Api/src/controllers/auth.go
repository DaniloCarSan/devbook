package controllers

import (
	"api/src/database"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"net/http"
)

// SignIn
func SignIn(w http.ResponseWriter, r *http.Request) {

	var email string = r.FormValue("email")
	var password string = r.FormValue("password")

	db, err := database.DatabaseConnect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.User(db)

	user, err := repository.ByEmail(email)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if security.CompareHashWithPassword(user.Password, password) != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("huhuhuhuhuhu está logado !"))
}

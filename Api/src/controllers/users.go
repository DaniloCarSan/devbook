package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Create
func Create(w http.ResponseWriter, r *http.Request) {

	user := models.User{
		Name:     r.FormValue("name"),
		Nickane:  r.FormValue("nickname"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err := user.Prepare(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DatabaseConnect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.User(db)

	user.ID, err = repository.Create(user)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// All
func All(w http.ResponseWriter, r *http.Request) {

	var nameOrNickname string = strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.DatabaseConnect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.User(db)

	users, err := repository.All(nameOrNickname)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// ById
func ById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.DatabaseConnect()

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.User(db)

	user, err := repository.ById(id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if user == (models.User{}) {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// Update
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}

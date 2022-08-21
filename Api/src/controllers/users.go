package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"net/http"
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

	repository := repositories.UserRepository(db)

	user.ID, err = repository.Create(user)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// All
func All(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All"))
}

// ById
func ById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ById"))
}

// Update
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}

// Delete
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}

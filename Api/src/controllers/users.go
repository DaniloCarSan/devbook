package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"fmt"
	"log"
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

	db, err := database.DatabaseConnect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UserRepository(db)

	id, err := repository.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Create user id:%d", id)))
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

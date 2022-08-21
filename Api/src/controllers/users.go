package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Create
func Create(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.DatabaseConnect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UserRepository(db)

	repository.Create(user)

	w.Write([]byte("Create"))
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

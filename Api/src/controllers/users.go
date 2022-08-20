package controllers

import "net/http"

// Create
func Create(w http.ResponseWriter, r *http.Request) {
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

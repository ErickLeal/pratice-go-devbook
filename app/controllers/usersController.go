package controllers

import "net/http"

func StoreUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("creating user"))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("listing usesr"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting user"))
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showing user"))
}

package controllers

import "net/http"

func ApiHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api ok!"))
}

func DbHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("db ok!"))
}

package controllers

import (
	"api/app/models"
	"api/app/repositories"
	"api/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error to read StoreUser request body")
		return
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		fmt.Println("Error to Unmarshal user")
		return
	}

	db, err := config.ConnectMysql()
	if err != nil {
		log.Fatal("error to connect mysql")
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		log.Fatal("error to create in mysql", err)
		return
	}

	w.Write([]byte(fmt.Sprintf("creating usesr %d", user.ID)))
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

package controllers

import (
	"api/app/models"
	"api/app/repositories"
	"api/config"
	"api/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid json format",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	if err = json.Unmarshal(request, &user); err != nil {
		utils.JSONResponse(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "Failed to read payload",
			"error":   err.Error(),
		})
		return
	}

	db, err := config.ConnectMysql()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "Unexpected error",
		})
		log.Println("error to connect mysql - ", err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	user.ID, err = repo.Create(user)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Error to create user",
			"error":   err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user_id": user.ID,
	})
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

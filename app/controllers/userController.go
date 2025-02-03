package controllers

import (
	models "api/app/models/user"
	"api/app/services"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {

	var userRequest models.UserCreateRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid json format",
			"error":   err.Error(),
		})
		return
	}

	validationErrors := utils.ValidateStuctRequest(userRequest)
	if len(validationErrors) > 0 {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	createdUser, err := services.CreateUser(userRequest)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "error to create user",
			"error":   err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"user":    createdUser.ToResponse(),
	})
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.ListAllUsers()

	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "error to list users",
			"error":   err.Error(),
		})
		return
	}

	usersResponse := make([]models.UserResponse, len(users))
	for i, user := range users {
		usersResponse[i] = user.ToResponse()
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "Success!",
		"user":    usersResponse,
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)

	userID, err := strconv.ParseUint(urlVars["id"], 10, 64)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}

	var userRequest models.UserUpdateRequest

	err = json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid json format",
			"error":   err.Error(),
		})
		return
	}

	validationErrors := utils.ValidateStuctRequest(userRequest)
	if len(validationErrors) > 0 {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	user, err := services.UpdateUser(uint64(userID), userRequest)

	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "error to update user",
			"error":   err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User updated!",
		"user":    user.ToResponse(),
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)

	userID, err := strconv.ParseUint(urlVars["id"], 10, 64)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}

	err = services.DeleteUser(uint64(userID))

	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "error to delete user",
			"error":   err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, nil)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)

	userID, err := strconv.ParseUint(urlVars["id"], 10, 64)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
			"error":   err.Error(),
		})
		return
	}

	user, err := services.GetUserById(uint64(userID))

	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"message": "error to find user",
			"error":   err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "Success!",
		"user":    user.ToResponse(),
	})
}

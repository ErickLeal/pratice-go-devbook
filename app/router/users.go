package router

import (
	"api/app/controllers"

	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.StoreUser).Methods("POST")
	router.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.ShowUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}

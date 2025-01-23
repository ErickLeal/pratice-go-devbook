package router

import (
	"api/app/controllers"

	"github.com/gorilla/mux"
)

func InitHealthRoutes(router *mux.Router) {
	router.HandleFunc("/api-health", controllers.ApiHealth).Methods("GET")
	router.HandleFunc("/db-health", controllers.DbHealth).Methods("GET")
}

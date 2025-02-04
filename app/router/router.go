package router

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

func GetRouter() *mux.Router {

	if router == nil {
		router = initializeRoutes()
	}
	return router
}

func initializeRoutes() *mux.Router {
	router := mux.NewRouter()

	InitUserRoutes(router)
	InitHealthRoutes(router)

	return router
}

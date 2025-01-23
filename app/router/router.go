package router

import (
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	InitUserRoutes(router)
	InitHealthRoutes(router)

	return router
}

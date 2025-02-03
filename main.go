package main

import (
	"api/app/router"
	"api/config"
	"net/http"
)

func main() {
	r := router.InitializeRoutes()

	config.LoadEnvs()
	config.ConnectMysql()

	http.ListenAndServe(":8080", r)
}

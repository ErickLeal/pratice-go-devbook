package main

import (
	"api/app/router"
	"api/config"
	"net/http"
)

func main() {
	r := router.GetRouter()

	config.LoadEnvs()

	http.ListenAndServe(":8080", r)
}

package main

import (
	"api/app/router"
	"api/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("é isso")

	r := router.InitializeRoutes()

	config.LoadEnvs()
	config.ConnectMysql()

	log.Fatal(http.ListenAndServe(":5000", r))

}

package main

import (
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/router"
)

func main() {
	router := router.Initialize()
	err := postgres.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(config.PORT, router)
}

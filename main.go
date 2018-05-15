package main

import (
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/postgres"
)

func main() {
	router := InitRouter()
	err := postgres.CreateTables()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(PORT, router)
}

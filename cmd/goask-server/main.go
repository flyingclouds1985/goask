package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/router"
)

func main() {
	store := postgres.New()
	err := store.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	router := router.Initialize(store)

	fmt.Println("App is running...")

	http.ListenAndServe(config.PORT, router)
}

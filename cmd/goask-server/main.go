package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/server"
)

var Ali string

func main() {
	store := postgres.New()
	err := store.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	server := &server.Server{}
	server.Store = store
	router := server.SetupRoute()

	fmt.Println("App is running...")
	http.ListenAndServe(config.PORT, router)
}

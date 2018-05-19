package main

import (
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/store"
)

func main() {
	router := InitRouter()

	err := store.CreateSchema()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(PORT, router)
}

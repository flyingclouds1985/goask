package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Setup()
	storeConf := &postgres.Config{Password: "secret"}
	store := postgres.New(storeConf)
	err := store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(store, gin.DebugMode)

	fmt.Printf("App is running on %s\n", server.Config.Port)
	http.ListenAndServe(server.Config.Port, server.Router)
}

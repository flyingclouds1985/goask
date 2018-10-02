package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/server"
	"github.com/gin-gonic/gin"
)

const (
	DBNAME     = "GOASK"
	DBPASSWORD = "secret"
	DBUSERNAME = "postgres"
)

func main() {
	storeConf := &postgres.Config{Password: "secret"}
	store := postgres.New(storeConf)
	err := store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	serverConf := &server.Config{RouterRealm: "goask.com"}
	server := server.New(store, gin.DebugMode, serverConf)

	fmt.Println("App is running...")
	http.ListenAndServe(server.Config.Port, server.Router)
}

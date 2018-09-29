package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/server"
)

const (
	DBNAME     = "GOASK"
	DBPASSWORD = "secret"
	DBUSERNAME = "postgres"
)

func main() {
	store := postgres.New(DBUSERNAME, DBPASSWORD, DBNAME)
	err := store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	server := &server.Server{
		Config: &server.Config{
			Port:            "localhost:9090",
			Domain:          "http://localhost:9090",
			RouterRealm:     "Question.com",
			RouterSecretKey: "asd!#@@#$nd189ehas-sS@mda",
		},
		Store: store,
	}
	router := server.SetupRouter(gin.DebugMode)

	fmt.Println("App is running...")
	http.ListenAndServe(server.Config.Port, router)
}

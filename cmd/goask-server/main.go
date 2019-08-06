package main

import (
	"fmt"
	"github.com/Alireza-Ta/goask/pkg/config"
	"github.com/Alireza-Ta/goask/postgres"
	"github.com/Alireza-Ta/goask/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	root, err := os.Getwd()
	c, err := config.Load(root + "/configuration.json")
	if err != nil {
		panic(err)
	}

	storeConf := postgres.Config{Password: c.GetString("database.password")}
	store := postgres.New(storeConf)
	err = store.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(store, gin.DebugMode, c)

	fmt.Printf("App is running on %d\n", c.GetInt("server.port"))
	http.ListenAndServe(string(c.GetInt("server.port")), s.Router)
}

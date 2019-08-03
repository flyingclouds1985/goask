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
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	root := filepath.Dir(filepath.Dir(currentDir))

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

	server := server.NewServer(store, gin.DebugMode, c)

	fmt.Printf("App is running on %s\n", server.Config.GetString("server.port"))
	http.ListenAndServe(server.Config.GetString("server.port"), server.Router)
}

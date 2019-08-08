package main

import (
	"fmt"
	"github.com/Alireza-Ta/goask/pkg/config"
	"github.com/Alireza-Ta/goask/postgres"
	"github.com/Alireza-Ta/goask/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"runtime"
)

var(
	_, currentFilePath, _, _ = runtime.Caller(0)
	rootPath = path.Dir(path.Dir(path.Dir(currentFilePath)))
)

func main() {
	c, err := config.Load(rootPath + "/configuration.json")
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
	err = http.ListenAndServe(c.GetString("server.domain"), s.Router)
	if err != nil {
		panic(err)
	}
}

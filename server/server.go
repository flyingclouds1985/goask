package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
)

// Config is server configuration parameters.
type Config struct {
	Port            string
	Domain          string
	RouterRealm     string
	RouterSecretKey string
}

// Server type.
type Server struct {
	Config *Config
	Router *gin.Engine
	Store  *postgres.Store
}

// New retunrs a new server.
func New(store *postgres.Store, routerMode string, conf *Config) *Server {
	if conf.Port == "" {
		conf.Port = "localhost:9090"
	}
	if conf.Domain == "" {
		conf.Domain = "http://localhost:9090"
	}
	if conf.RouterRealm == "" {
		conf.RouterRealm = "example.com"
	}
	if conf.RouterSecretKey == "" {
		k := os.Getenv("RouterSecretKey")
		if k == "" {
			s := RandomString(20)
			err := os.Setenv("RouterSecretKey", s)
			if err != nil {
				log.Println(err)
			}
			conf.RouterSecretKey = s
		} else {
			conf.RouterSecretKey = k
		}
	}

	router := NewRouter(routerMode)
	server := &Server{
		Config: conf,
		Store:  store,
		Router: router,
	}
	server.Routes()

	return server
}

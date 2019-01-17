package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
)

// Config is server configuration parameters.
type Config struct {
	Port   string
	Domain string
}

// Server type.
type Server struct {
	Config Config
	Router *gin.Engine
	Store  *postgres.Store
}

// NewServer is the entry point of the system.
func NewServer(store *postgres.Store, routerMode string, config ...Config) *Server {
	conf := initServerConfig(config)

	router := NewRouter(routerMode)
	server := &Server{
		Config: conf,
		Store:  store,
		Router: router,
	}
	server.Routes()

	return server
}

func initServerConfig(config []Config) Config {
	defaultConfig := Config{
		Port:   "localhost:9090",
		Domain: "http://localhost:9090",
	}
	switch len(config) {
	case 0:
		return defaultConfig
	case 1:
		conf := config[0]
		if conf.Port == "" {
			conf.Port = defaultConfig.Port
		}
		if conf.Domain == "" {
			conf.Domain = defaultConfig.Domain
		}
		return conf
	default:
		panic("too much argument!")
	}
}

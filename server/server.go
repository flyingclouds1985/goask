package server

import (
	"github.com/Alireza-Ta/GOASK/pkg/config"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/gin-gonic/gin"
)

// Server type.
type Server struct {
	Config *config.Config
	Router *gin.Engine
	Store  *postgres.Store
}

// NewServer is the entry point of the system.
func NewServer(store *postgres.Store, routerMode string, config *config.Config) *Server {

	router := NewRouter(routerMode)
	server := &Server{
		Config: config,
		Store:  store,
		Router: router,
	}
	server.Routes()

	return server
}



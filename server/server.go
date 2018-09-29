package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
)

type Config struct {
	Port            string
	Domain          string
	RouterRealm     string
	RouterSecretKey string
}

// Server type
type Server struct {
	Config *Config
	Router *gin.Engine
	Store  *postgres.Store
}

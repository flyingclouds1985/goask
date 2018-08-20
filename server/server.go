package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Alireza-Ta/GOASK/postgres"
)

// Server type
type Server struct {
	Router *gin.Engine
	Store  *postgres.Store
}

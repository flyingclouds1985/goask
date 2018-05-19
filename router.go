package main

import (
	"net/http"

	"github.com/Alireza-Ta/GOASK/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routeList(r)

	return r
}

func routeList(router *gin.Engine) {
	router.POST("/login", AuthMiddleware().LoginHandler)

	auth := router.Group("/auth")
	auth.GET("/refresh_token", AuthMiddleware().RefreshHandler)

	q := router.Group("questions")
	{
		q.GET("/ask", api.GetAskQuestion)
	}
}

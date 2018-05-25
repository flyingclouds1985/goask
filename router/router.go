package router

import (
	"net/http"

	"github.com/Alireza-Ta/GOASK/api"
	"github.com/gin-gonic/gin"
)

func Initialize() http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routeList(r)

	return r
}

func routeList(router *gin.Engine) {
	router.POST("/login", AuthMiddleware().LoginHandler)

	// for testing purposes
	router.GET("/test", func(c *gin.Context) {

	})

	auth := router.Group("/auth")
	auth.GET("/refresh_token", AuthMiddleware().RefreshHandler)

	q := router.Group("questions")
	{
		q.GET("/", api.GetQuestionList)
		q.GET("/:id/:question", api.GetQuestion)
		q.PATCH("/:id", api.PatchQuestion)
		q.POST("/ask", api.PostAskQuestion)
	}

}

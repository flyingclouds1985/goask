package router

import (
	"net/http"

	"github.com/Alireza-Ta/GOASK/api"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/gin-gonic/gin"
)

func Initialize(store *postgres.Store) http.Handler {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("web/template/*")
	r.Static("/assets", "web/static")

	routeList(r, store)

	return r
}

func routeList(router *gin.Engine, store *postgres.Store) {
	router.POST("/login", AuthMiddleware().LoginHandler)
	auth := router.Group("/auth")
	auth.GET("/refresh_token", AuthMiddleware().RefreshHandler)

	api := &api.Api{
		Store: store,
	}

	// for testing purposes
	router.GET("/test", func(c *gin.Context) {

	})

	q := router.Group("questions")
	{
		q.GET("/", api.GetQuestionList)
		q.POST("/", api.PostQuestion)
		q.GET("/:id", api.GetQuestion)
		q.PATCH("/:id", api.PatchQuestion)
		q.GET("/:id/:question", api.GetQuestion)
		q.PATCH("/:id/:vote", api.PatchVoteQuestion)
	}

	c := router.Group("comments")
	{
		c.GET("/questions/:question_id", api.GetQuestionCommentList)
		c.POST("/questions/:question_id", api.PostQuestionComment)
		c.GET("/replies/:reply_id", api.GetReplyCommentList)
		c.POST("/replies/:reply_id", api.PostReplyComment)
	}

	r := router.Group("replies")
	{
		r.GET("/questions/:question_id", api.GetReplyList)
		r.POST("/questions/:question_id", api.PostReply)
		r.PATCH("/:reply_id/questions/:question_id/", api.PatchReply)
	}
}

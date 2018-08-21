package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoute setups gin with common middlewares.
func (s *Server) SetupRoute() http.Handler {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	s.Router = r
	s.routeList(r)

	return r
}

func (s *Server) routeList(router *gin.Engine) {
	s.Router.POST("/login", AuthMiddleware().LoginHandler)
	auth := s.Router.Group("/auth")
	auth.GET("/refresh_token", AuthMiddleware().RefreshHandler)

	// for testing purposes
	s.Router.GET("/test", func(c *gin.Context) {

	})

	q := s.Router.Group("questions")
	{
		q.GET("/", s.GetQuestionList)
		q.POST("/", s.PostQuestion)
		q.GET("/:id", s.GetQuestion)
		q.PATCH("/:id", s.PatchQuestion)
		q.GET("/:id/:question", s.GetQuestion)
		q.PATCH("/:id/:vote", s.PatchVoteQuestion)
	}

	c := s.Router.Group("comments")
	{
		c.GET("/questions/:question_id", s.GetQuestionCommentList)
		c.POST("/questions/:question_id", s.PostQuestionComment)
		c.GET("/replies/:reply_id", s.GetReplyCommentList)
		c.POST("/replies/:reply_id", s.PostReplyComment)
	}

	r := s.Router.Group("replies")
	{
		r.GET("/questions/:question_id", s.GetReplyList)
		r.POST("/questions/:question_id", s.PostReply)
		r.PATCH("/:reply_id/questions/:question_id/", s.PatchReply)
	}
}

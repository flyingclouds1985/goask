package server

import (
	"net/http"

	"github.com/Alireza-Ta/GOASK/server/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoute setups gin with common middlewares.
func (s *Server) SetupRoute(mode string) http.Handler {
	gin.SetMode(mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	s.Router = r
	s.routeList(r)

	return r
}

func (s *Server) routeList(router *gin.Engine) {
	public := s.Router.Group("/")
	{
		q := public.Group("questions")
		{
			q.GET("/", s.GetQuestionList)
			q.GET("/:id", s.GetQuestion)
			q.GET("/:id/:question", s.GetQuestion)
		}
		c := public.Group("comments")
		{
			c.GET("/questions/:question_id", s.GetQuestionCommentList)
			c.GET("/replies/:reply_id", s.GetReplyCommentList)
		}
		r := public.Group("replies")
		{
			r.GET("/questions/:question_id", s.GetReplyList)
		}

		// for testing purposes
		public.GET("/test", func(c *gin.Context) {

		})
	}

	private := s.Router.Group("/")
	{
		private.POST("/login", middleware.Auth().LoginHandler)
		private.GET("/refresh_token", middleware.Auth().RefreshHandler)
		q := private.Group("questions")
		{
			q.POST("/", s.PostQuestion)
			q.PATCH("/:id", s.PatchQuestion)
			q.PATCH("/:id/:vote", s.PatchVoteQuestion)
		}

		c := private.Group("comments")
		{
			c.POST("/questions/:question_id", s.PostQuestionComment)
			c.POST("/replies/:reply_id", s.PostReplyComment)
		}

		r := private.Group("replies")
		{
			r.POST("/questions/:question_id", s.PostReply)
			r.PATCH("/:reply_id/questions/:question_id/", s.PatchReply)
		}

		u := private.Group("users")
		{
			u.POST("/", s.PostUser)
		}
	}

}

package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter setups gin with common middlewares.
func NewRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	return r
}

func (s *Server) Routes() {
	public := s.Router.Group("/")
	{
		public.GET("login")
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
		u := public.Group("users")
		{
			u.GET("/:username", s.GetUser)
		}
		// for testing purposes
		public.GET("/test", func(c *gin.Context) {

		})
	}

	private := s.Router.Group("/")
	{
		private.POST("login", s.Auth().LoginHandler)
		auth := private.Group("auth")
		auth.Use(s.Auth().MiddlewareFunc())
		{
			auth.GET("refresh_token", s.Auth().RefreshHandler)
			auth.GET("hello", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"data": "asd",
					"c":    c.Keys,
				})
			})
		}
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
			u.PATCH("/:id", s.PatchUser)
		}
	}

}

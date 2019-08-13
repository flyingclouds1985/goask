package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter setups gin with common middlewares.
func NewRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.RouterGroup.Use()

	return r
}

// Routes is list of routes.
func (s *Server) Routes() {
	authHandler := AuthAPI{
		jwtRealm:     s.Config.GetString("router.realm"),
		jwtSecretKey: s.Config.GetString("router.secretKey"),
		store:        s.Store,
	}

	questionHandler := QuestionAPI{
		store:  s.Store,
		domain: s.Config.GetString("server.domain"),
	}

	commentHandler := CommentAPI{
		store: s.Store,
	}

	replyHandler := ReplyAPI{
		store: s.Store,
	}

	userHandler := UserAPI{
		store: s.Store,
	}

	public := s.Router.Group("/")
	{
		public.POST("login", authHandler.Auth().LoginHandler)
		q := public.Group("questions")
		{
			q.GET("/", questionHandler.GetQuestionList)
			q.GET("/:id", questionHandler.GetQuestion)
			q.GET("/:id/:question", questionHandler.GetQuestion)
		}
		c := public.Group("comments")
		{
			c.GET("/questions/:question_id", commentHandler.GetQuestionCommentList)
			c.GET("/replies/:reply_id", commentHandler.GetReplyCommentList)
		}
		r := public.Group("replies")
		{
			r.GET("/questions/:question_id", replyHandler.GetReplyList)
		}
		u := public.Group("users")
		{
			u.GET("/:username", userHandler.GetUser)
		}
		// for testing purposes
		public.GET("/test", func(c *gin.Context) {
			u, err := s.Store.FindUserByLoginCredentials(c.Query("username"), c.Query("password"))
			if err != nil {
				panic(err)
			}
			c.JSON(200, gin.H{
				"data": u,
			})
		})
	}

	private := s.Router.Group("/")
	// private.Use(authHandler.Auth().MiddlewareFunc())
	private.Use()
	{
		auth := private.Group("auth")
		auth.Use(authHandler.Auth().MiddlewareFunc())
		{
			auth.GET("refresh_token", authHandler.Auth().RefreshHandler)
			auth.GET("/hello", func(c *gin.Context) {
				claims := jwt.ExtractClaims(c)
				user, _ := c.Get("id")
				c.JSON(200, gin.H{
					"userID":   claims,
					"userName": user,
					"text":     "Hello World.",
				})
			})
		}
		q := private.Group("questions")
		{
			q.POST("/", questionHandler.PostQuestion)
			q.PATCH("/:id", questionHandler.PatchQuestion)
			q.PATCH("/:id/:vote", questionHandler.PatchVoteQuestion)
		}

		c := private.Group("comments")
		{
			c.POST("/questions/:question_id", commentHandler.PostQuestionComment)
			c.POST("/replies/:reply_id", commentHandler.PostReplyComment)
		}

		r := private.Group("replies")
		{
			r.POST("/questions/:question_id", replyHandler.PostReply)
			r.PATCH("/:reply_id/questions/:question_id/", replyHandler.PatchReply)
		}

		u := private.Group("users")
		{
			u.POST("/", userHandler.PostUser)
			u.PATCH("/:id", userHandler.PatchUser)
		}
	}

}

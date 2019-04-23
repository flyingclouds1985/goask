package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// Auth is the Authentication middleware.
func (s *Server) Auth() *jwt.GinJWTMiddleware {
	realm, err := config.GetString("router.realm")
	if err != nil {
		log.Fatal(err)
	}
	secretKey, err := config.GetString("router.secretKey")
	if err != nil {
		log.Fatal(err)
	}

	return &jwt.GinJWTMiddleware{
		Realm:      realm,
		Key:        []byte(secretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		SigningAlgorithm: "HS256",

		Authenticator: authenticator,

		Authorizator: authorizator,

		Unauthorized: unauthorized,

		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"context": c.Keys,
			})
		},

		TokenLookup: "header:Authorization",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,

		SendAuthorization: true,
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	var data model.User
	c.ShouldBindJSON(&data)
	if data.Username == "admin" && data.Password == "admin" {
		return "admin", nil
	}

	return "admin", jwt.ErrFailedAuthentication
}

func authorizator(data interface{}, c *gin.Context) bool {
	// if v, ok := data.(string); ok && v == "admin" {
	// 	return true
	// }

	// return false
	return true
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"context": c.Keys,
	})
}

package router

import (
	"time"

	"github.com/Alireza-Ta/GOASK/config"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      config.DOMAIN,
		Key:        []byte(config.SECRET_KEY),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,

		Authenticator: authenticator,

		Authorizator: authorizator,

		Unauthorized: unauthorized,

		TokenLookup: "header:Authorization",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	}
}

func authenticator(username, password string, c *gin.Context) (string, bool) {
	if username == "admin" && password == "admin" {
		return username, true
	}

	return username, false
}

func authorizator(username string, c *gin.Context) bool {
	if username == "admin" {
		return true
	}

	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

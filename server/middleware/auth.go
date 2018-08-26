package middleware

import (
	"time"

	"github.com/Alireza-Ta/GOASK/config"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Auth() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      config.REALM,
		Key:        []byte(config.SECRETKEY),
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

func authenticator(c *gin.Context) (interface{}, error) {
	if c.Param("username") == "admin" && c.Param("password") == "admin" {
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
	})
}

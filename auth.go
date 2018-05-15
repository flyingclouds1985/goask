package main

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

const (
	PORT       = "localhost:9090"
	DOMAIN     = "Question.com"
	SECRET_KEY = "asd!#@@#$nd189ehas-sS@mda"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      DOMAIN,
		Key:        []byte(SECRET_KEY),
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

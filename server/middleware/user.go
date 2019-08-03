package middleware

import (
	"github.com/gin-gonic/gin"
)

// MustUser is a middleware that checks if the user is logged in or not.
func MustUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

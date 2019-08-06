package server

import (
	"net/http"
	"time"

	"github.com/Alireza-Ta/goask/model"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//AuthAPI manages authentication stuffs.
type AuthAPI struct {
	jwtRealm     string
	jwtSecretKey string
}

// Auth is the Authentication middleware.
func (a *AuthAPI) Auth() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      a.jwtRealm,
		Key:        []byte(a.jwtSecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: "username",
		SigningAlgorithm: "HS256",

		Authenticator: authenticator,

		Authorizator: authorizator,

		Unauthorized: unauthorized,

		IdentityHandler: func(c *gin.Context) interface{} {
			// claims := jwt.ExtractClaims(c)
			// return &model.User{
			// 	Username: claims["id"].(string),
			// }
			claims := jwt.ExtractClaims(c)
			return claims[jwt.IdentityKey]
		},

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time){
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},

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

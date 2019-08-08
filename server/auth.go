package server

import (
	"log"
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
	store UserStore
}

// Auth is the Authentication middleware.
func (a *AuthAPI) Auth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            a.jwtRealm,
		Key:              []byte(a.jwtSecretKey),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      "username",
		SigningAlgorithm: "HS256",

		Authenticator: a.authenticator,

		Authorizator: a.authorizator,

		Unauthorized: a.unauthorized,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if u, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"user": u,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			// claims := jwt.ExtractClaims(c)
			// return &model.User{
			// 	Username: claims["id"].(string),
			// }
			claims := jwt.ExtractClaims(c)
			return claims[jwt.IdentityKey]
		},

		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
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
	})

	if err != nil {
		log.Fatal("JWT Error: ", err)
	}

	return authMiddleware
}

func (a *AuthAPI) authenticator(c *gin.Context) (interface{}, error) {
	var loginValues model.User

	if err := c.ShouldBindJSON(&loginValues); err != nil {
		return err, jwt.ErrMissingLoginValues
	}
	username := loginValues.Username
	password := loginValues.Password
	user, err := a.store.FindUserByLoginCredentials(username, password)
	if err !=nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return user, nil
}

func (a *AuthAPI) authorizator(data interface{}, c *gin.Context) bool {
	// if v, ok := data.(string); ok && v == "admin" {
	// 	return true
	// }

	// return false
	return true
}

func (a *AuthAPI) unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

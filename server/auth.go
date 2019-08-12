package server

import (
	"fmt"
	"github.com/Alireza-Ta/goask/model"
	"github.com/Alireza-Ta/goask/validation"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"time"
)

var validationErr map[string]string

// AuthStore manages encapsulated database access.
type AuthStore interface {
	FindUserByLoginCredentials(username, password string) (*model.User, error)
}

//AuthAPI manages authentication stuffs.
type AuthAPI struct {
	jwtRealm     string
	jwtSecretKey string
	store        UserStore
	errors       map[string]interface{}
}

// Auth is the Authentication middleware.
func (a *AuthAPI) Auth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            a.jwtRealm,
		Key:              []byte(a.jwtSecretKey),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      "id",
		SigningAlgorithm: "HS256",

		Authenticator: a.authenticator,

		Authorizator: a.authorizator,

		Unauthorized: a.unauthorized,

		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			if ve, ok := e.(validator.ValidationErrors); ok {
				validationErr = validation.Messages(ve)
			}

			return e.Error()
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if u, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"user": u,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := claims["user"].(map[string]interface{})
			return &model.User{
				Id:       int(user["id"].(float64)),
				Username: user["username"].(string),
				Email:    user["email"].(string),
			}
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

	err := c.ShouldBindJSON(&loginValues)

	// Ignore required on ConfirmPassword and Email fields.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.StructExcept(loginValues, "ConfirmPassword", "Email")
		//Ignore eqfield on Password field.
		if ve, ok := err.(validator.ValidationErrors); ok {
			if ve["User.Password"].Tag == "eqfield" {
				fmt.Println("ccc", ve)
				delete(ve, "User.Password")
				err = ve
			}
		}
	}

	if err != nil && len(err.(validator.ValidationErrors)) != 0 {
		return nil, err
	}

	username := loginValues.Username
	password := loginValues.Password
	user, err := a.store.FindUserByLoginCredentials(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// data is the returned value of IdentityHandler.
func (a *AuthAPI) authorizator(data interface{}, c *gin.Context) bool {
	// if v, ok := data.(string); ok && v == "admin" {
	// 	return true
	// }
	fmt.Println("data: g", data)
	// return false
	return true
}

func (a *AuthAPI) unauthorized(c *gin.Context, code int, message string) {
	if validationErr != nil {
		JSONValidation(validationErr, c)
	} else {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

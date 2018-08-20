package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONBadRequestError expresses bad rquest error in json format.
func JSONBadRequestError(message string, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  message + err.Error(),
		"status": 400,
	})
	return
}

// RenderByContentType represents multiple way of rendering based on header content-type.
func RenderByContentType(status int, c *gin.Context, data interface{}, tmplName string) {
	if c.GetHeader("Content-type") == "application/json" {
		c.JSON(status, data)
	} else {
		c.HTML(status, tmplName, data)
	}
}

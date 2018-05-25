package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONBadRequestError(message string, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  message + err.Error(),
		"status": 400,
	})
	return
}

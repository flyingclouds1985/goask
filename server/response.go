package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONBadRequest responds bad rquest error in json format.
func JSONBadRequest(customErr string, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": map[string]interface{}{
			"status":  400,
			"message": customErr + err.Error(),
		},
	})
}

// JSONValidation responds validation errors in json format.
func JSONValidation(messages map[string]string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": map[string]interface{}{
			"status":  400,
			"message": messages,
		},
	})
}

// JSONNotFound responds not found error in json format.
func JSONNotFound(customErr string, err error, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"errors": map[string]interface{}{
			"status":  404,
			"message": customErr + err.Error(),
		},
	})
}

// JSONInternalServer responds not found error in json format.
func JSONInternalServer(customErr string, err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"errors": map[string]interface{}{
			"status":  500,
			"message": customErr + " Actual Error: " + err.Error(),
		},
	})
}

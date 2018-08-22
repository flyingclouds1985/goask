package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONBadRequestError expresses bad rquest error in json format.
func JSONBadRequestError(customErr error, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  customErr.Error() + err.Error(),
		"status": 400,
	})
	return
}

// JSONNotFoundError expresses not found request error in json format.
func JSONNotFoundError(customErr error, err error, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  customErr.Error() + err.Error(),
		"status": 404,
	})
	return
}

func NotFoundErr(context string) error {
	return errors.New(context + " not found.")
}

func BindErr(context string) error {
	return errors.New(context + " error in binding.")
}

func InsertErr(context string) error {
	return errors.New(context + " error in inserting.")
}

func UpdateErr(context string) error {
	return errors.New(context + " error in updating.")
}

func VoteErr(context string) error {
	return errors.New(context + " voting error.")
}

func ListErr(context string) error {
	return errors.New(context + " error in getting list.")
}
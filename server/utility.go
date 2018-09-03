package server

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// JSONBadRequestError expresses bad rquest error in json format.
func JSONBadRequestError(customErr error, err error, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  customErr.Error() + err.Error(),
		"status": 400,
	})
}

// JSONNotFoundError expresses not found request error in json format.
func JSONNotFoundError(customErr error, err error, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  customErr.Error() + err.Error(),
		"status": 404,
	})
}

func JSONValidationError(messages map[string]string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": map[string]interface{}{
			"status":   "400",
			"messages": messages,
		},
	})
}

func NotFoundErr(context string) error {
	return errors.New(context + " not found. ")
}

func BindErr(context string) error {
	return errors.New(context + " error in binding. ")
}

func InsertErr(context string) error {
	return errors.New(context + " error in inserting. ")
}

func UpdateErr(context string) error {
	return errors.New(context + " error in updating. ")
}

func VoteErr(context string) error {
	return errors.New(context + " voting error. ")
}

func ListErr(context string) error {
	return errors.New(context + " error in getting list. ")
}

func checkNil(item interface{}, message string) {
	var err string
	if e, ok := item.(error); ok {
		err = e.Error()
	}
	if item != nil {
		fmt.Printf("Error: %s, Message %s", err, message)
	}
}

// HashPassword makes password in bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash checks hashed password against string password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

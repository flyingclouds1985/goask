package api

import (
	"fmt"
	"net/http"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/gin-gonic/gin"
)

func GetAskQuestion(c *gin.Context) {
	fmt.Print("asd")
}

func PostAskQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Binding...", err, c)
		return
	}

	q := &model.Question{
		Title: in.Title,
		Post: model.Post{
			Body:     in.Body,
			AuthorID: in.AuthorID,
		},
	}

	if err = postgres.CreateQuestion(q); err != nil {
		JSONBadRequestError("Inserting...", err, c)
		return
	}

	c.JSON(http.StatusOK, q)
}

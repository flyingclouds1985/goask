package api

import (
	"net/http"
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/gin-gonic/gin"
)

func GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	q, err := postgres.QuestionFind(id)

	if err != nil {
		JSONBadRequestError("Error in finding question. ", err, c)
	}

	// check question param if does not equal to title redirect and edit it.
	c.JSON(http.StatusOK, q)
}

func PostAskQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding question. ", err, c)
	}

	q := &model.Question{
		Title: in.Title,
		Post: model.Post{
			Body:     in.Body,
			AuthorID: in.AuthorID,
		},
	}

	if err = postgres.CreateQuestion(q); err != nil {
		JSONBadRequestError("Error in inserting question. ", err, c)
	}

	c.JSON(http.StatusOK, q)
}

func GetQuestionList(c *gin.Context) {
	list, err := postgres.QuestionList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError("Error in getting the questions list. ", err, c)
	}

	c.JSON(http.StatusOK, list)
}

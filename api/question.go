package api

import (
	"net/http"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/gin-gonic/gin"
)

// GetQuestion returns a question based on id and title.
func GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Param("question")
	q, err := postgres.QuestionFind(id)

	if err != nil {
		JSONBadRequestError("Error in finding question. ", err, c)
		return
	}
	// rewrite url if the question title isn't correct.
	s := slug.Make(q.Title)
	if title != s {
		c.Redirect(http.StatusTemporaryRedirect, config.DOMAIN+"/questions/"+c.Param("id")+"/"+s)
	}

	c.JSON(200, q)
}

// PostAskQuestion creates a question.
func PostAskQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding question. ", err, c)
	}

	q := new(model.Question)
	q.Title = in.Title
	q.Body = in.Body
	q.AuthorID = in.AuthorID

	if err = postgres.CreateQuestion(q); err != nil {
		JSONBadRequestError("Error in inserting question. ", err, c)
	}

	c.JSON(200, q)
}

// PatchQuestion upadte a question.
func PatchQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	// If client didn't specified the id in the request.
	if in.Id == 0 {
		id, _ := strconv.Atoi(c.Param("id"))
		in.Id = id
	}

	if err != nil {
		JSONBadRequestError("Error in binding question. ", err, c)
		return
	}

	if err = postgres.QuestionUpdate(in); err != nil {
		JSONBadRequestError("Error in updating question. ", err, c)
		return
	}

	c.JSON(200, in)
}

// PatchVoteQuestion gives a vote to a question.
func PatchVoteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := c.Param("vote")
	q, err := postgres.QuestionFind(id)

	// Check if there's such a question.
	if err != nil {
		JSONBadRequestError("Error in finding question. ", err, c)
		return
	}

	if v == "upvote" {
		q.Vote++
	} else if v == "downvote" {
		q.Vote--
	}

	err = postgres.QuestionVoteUpdate(q)
	if err != nil {
		JSONBadRequestError("Error in voting question. ", err, c)
		return
	}

	s := slug.Make(q.Title)
	c.Redirect(http.StatusSeeOther, config.DOMAIN+"/questions/"+c.Param("id")+"/"+s)
}

// GetQuestionList returns a list of questions.
func GetQuestionList(c *gin.Context) {
	list, err := postgres.QuestionsList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError("Error in getting the questions list. ", err, c)
	}

	c.JSON(200, list)
}

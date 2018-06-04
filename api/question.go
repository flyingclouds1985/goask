package api

import (
	"net/http"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

// GetQuestion returns a question based on id and title.
func (a *Api) GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Param("question")
	q, err := a.Store.QuestionFind(id)

	if err != nil {
		JSONBadRequestError("Error in finding question. ", err, c)
	}
	// rewrite url if the question title isn't correct.
	s := slug.Make(q.Title)
	if title != s {
		c.Redirect(http.StatusTemporaryRedirect, config.DOMAIN+"/questions/"+c.Param("id")+"/"+s)
	}

	c.JSON(200, q)
}

// PostAskQuestion creates a question.
func (a *Api) PostQuestion(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Question)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding question. ", err, c)
	}

	q := new(model.Question)
	q.Title = in.Title
	q.Body = in.Body
	// q.AuthorID = claims["id"]

	if err = a.Store.CreateQuestion(q); err != nil {
		JSONBadRequestError("Error in inserting question. ", err, c)
	}

	c.JSON(200, q)
}

// PatchQuestion upadte a question.
func (a *Api) PatchQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	// If client didn't specified the id in the request.
	if in.Id == 0 {
		id, _ := strconv.Atoi(c.Param("id"))
		in.Id = id
	}

	if err != nil {
		JSONBadRequestError("Error in binding question. ", err, c)
	}

	if err = a.Store.QuestionUpdate(in); err != nil {
		JSONBadRequestError("Error in updating question. ", err, c)
	}

	c.JSON(200, in)
}

// PatchVoteQuestion gives a vote to a question.
func (a *Api) PatchVoteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := c.Param("vote")
	q, err := a.Store.QuestionFind(id)

	// Check if there's such a question.
	if err != nil {
		JSONBadRequestError("Error in finding question. ", err, c)
	}

	if v == "upvote" {
		q.Vote++
	} else if v == "downvote" {
		q.Vote--
	}

	err = a.Store.QuestionVoteUpdate(q)
	if err != nil {
		JSONBadRequestError("Error in voting question. ", err, c)
	}

	s := slug.Make(q.Title)
	c.Redirect(http.StatusSeeOther, config.DOMAIN+"/questions/"+c.Param("id")+"/"+s)
}

// GetQuestionList returns a list of questions.
func (a *Api) GetQuestionList(c *gin.Context) {
	list, err := a.Store.QuestionsList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError("Error in getting the questions list. ", err, c)
	}

	c.JSON(200, list)
}

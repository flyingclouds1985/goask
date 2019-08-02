package server

import (
	"net/http"
	"strconv"
	"net/url"
	"github.com/Alireza-Ta/goask/model"
	"github.com/Alireza-Ta/goask/validation"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

//QuestionStore manages encapsulated database access.
type QuestionStore interface {
	QuestionWithRelations(id int) (*model.Question, error)
	CreateQuestion(q *model.Question) error 
	CreateTag(tags []*model.Tag, qid int)
	UpdateQuestion(q *model.Question) error
	FindQuestion(id int) (*model.Question, error)
	UpdateVote(q *model.Question) error
	ListQuestion(query url.Values) (model.Questions, error)
}

//QuestionAPI provides handlers for managing questions.
type QuestionAPI struct {
	store QuestionStore
	domain string
}

// GetQuestion returns a question based on id and title.
func (qapi *QuestionAPI) GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	q, err := qapi.store.QuestionWithRelations(id)
	if err != nil {
		JSONNotFound("Error question not found. ", err, c)
		return
	}

	// rewrite url if the question title does not provide or isn't in correct format.
	title := c.Param("question")
	titleSlug := slug.Make(q.Title)
	if title != titleSlug {
		urlStr := qapi.domain + "/questions/" + c.Param("id") + "/" + titleSlug
		c.Redirect(http.StatusSeeOther, urlStr)
		return
	}

	c.JSON(http.StatusOK, q)
}

// PostQuestion creates a question.
func (qapi *QuestionAPI) PostQuestion(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Question)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	q := new(model.Question)
	q.Title = in.Title
	q.Body = in.Body
	q.Tags = in.Tags
	// q.AuthorID = claims["id"]
	if err := qapi.store.CreateQuestion(q); err != nil {
		JSONInternalServer("Error inserting question. ", err, c)
		return
	}
	qapi.store.CreateTag(in.Tags, q.Id)

	c.JSON(http.StatusOK, q)
}

// PatchQuestion upadtes a question.
func (qapi *QuestionAPI) PatchQuestion(c *gin.Context) {
	in := new(model.Question)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}
	// If client didn't specified the id in the request.
	// if in.Id == 0 {
	// 	id, _ := strconv.Atoi(c.Param("id"))
	// 	in.Id = id
	// }

	if err := qapi.store.UpdateQuestion(in); err != nil {
		JSONInternalServer("Error updating question. ", err, c)
		return
	}

	c.JSON(http.StatusOK, in)
}

// PatchVoteQuestion gives a vote to a question.
func (qapi *QuestionAPI) PatchVoteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := c.Param("vote")
	q, err := qapi.store.FindQuestion(id)
	if err != nil {
		JSONNotFound("Error question not found. ", err, c)
		return
	}
	if v == "upvote" {
		q.Vote++
	} else if v == "downvote" {
		q.Vote--
	}

	err = qapi.store.UpdateVote(q)
	if err != nil {
		JSONInternalServer("Error in question voting. ", err, c)
		return
	}

	titleSlug := slug.Make(q.Title)
	c.Redirect(http.StatusSeeOther, qapi.domain+"/questions/"+c.Param("id")+"/"+titleSlug)
}

// GetQuestionList returns a list of questions.
func (qapi *QuestionAPI) GetQuestionList(c *gin.Context) {
	list, err := qapi.store.ListQuestion(c.Request.URL.Query())

	if err != nil {
		JSONNotFound("Error finding questions list. ", err, c)
		return
	}

	c.JSON(http.StatusOK, list)
}

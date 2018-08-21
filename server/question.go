package server

import (
	"net/http"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

// GetQuestion returns a question based on id and title.
func (s *Server) GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.Param("question")
	q, err := s.Store.QuestionSingleWithRelations(id)

	if err != nil {
		JSONNotFoundError(NotFoundErr("question"), err, c)
	}

	// rewrite url if the question title isn't correct.
	titleSlug := slug.Make(q.Title)
	if title != titleSlug {
		urlStr := config.DOMAIN + "/questions/" + c.Param("id") + "/" + titleSlug
		c.Writer.Header().Set("Location", urlStr)
	}

	c.JSON(200, q)
}

// PostQuestion creates a question.
func (s *Server) PostQuestion(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Question)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError(BindErr("question"), err, c)
	}

	tags, _ := s.Store.TagCreate(in.TagString)

	q := new(model.Question)

	q.Title = in.Title
	q.Body = in.Body
	q.Tags = tags
	// q.AuthorID = claims["id"]

	if err = s.Store.QuestionCreate(q); err != nil {
		JSONBadRequestError(InsertErr("question"), err, c)
	}

	c.JSON(200, q)
}

// PatchQuestion upadte a question.
func (s *Server) PatchQuestion(c *gin.Context) {
	in := new(model.Question)
	err := c.ShouldBind(in)

	// If client didn't specified the id in the request.
	if in.Id == 0 {
		id, _ := strconv.Atoi(c.Param("id"))
		in.Id = id
	}

	if err != nil {
		JSONBadRequestError(BindErr("question"), err, c)
	}

	if err = s.Store.QuestionUpdate(in); err != nil {
		JSONBadRequestError(UpdateErr("question"), err, c)
	}

	c.JSON(200, in)
}

// PatchVoteQuestion gives a vote to a question.
func (s *Server) PatchVoteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := c.Param("vote")
	q, err := s.Store.QuestionFind(id)

	// Check if there's such a question.
	if err != nil {
		JSONNotFoundError(NotFoundErr("question"), err, c)
	}

	if v == "upvote" {
		q.Vote++
	} else if v == "downvote" {
		q.Vote--
	}

	err = s.Store.QuestionVoteUpdate(q)
	if err != nil {
		JSONBadRequestError(VoteErr("question"), err, c)
	}

	titleSlug := slug.Make(q.Title)
	c.Redirect(http.StatusSeeOther, config.DOMAIN+"/questions/"+c.Param("id")+"/"+titleSlug)
}

// GetQuestionList returns a list of questions.
func (s *Server) GetQuestionList(c *gin.Context) {
	list, err := s.Store.QuestionsList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError(ListErr("question"), err, c)
	}

	c.JSON(200, list)
}

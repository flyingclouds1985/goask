package server

import (
	"strconv"

	"github.com/Alireza-Ta/GOASK/config"
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/validation"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

// GetQuestion returns a question based on id and title.
func (s *Server) GetQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	q, err := s.Store.QuestionWithRelations(id)
	if err != nil {
		JSONNotFound("Error question not found. ", err, c)
		return
	}

	// rewrite url if the question title does not provide or isn't in correct format.
	title := c.Param("question")
	titleSlug := slug.Make(q.Title)
	if title != titleSlug {
		urlStr := config.DOMAIN + "/questions/" + c.Param("id") + "/" + titleSlug
		c.Redirect(301, urlStr)
		return
	}
	c.JSON(200, q)
}

// PostQuestion creates a question.
func (s *Server) PostQuestion(c *gin.Context) {
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
	if err := s.Store.QuestionCreate(q); err != nil {
		JSONInternalServer("Error inserting question. ", err, c)
		return
	}
	s.Store.TagCreate(in.Tags, q.Id)

	c.JSON(200, q)
}

// PatchQuestion upadtes a question.
func (s *Server) PatchQuestion(c *gin.Context) {
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

	if err := s.Store.QuestionUpdate(in); err != nil {
		JSONInternalServer("Error updating question. ", err, c)
		return
	}

	c.JSON(200, in)
}

// PatchVoteQuestion gives a vote to a question.
func (s *Server) PatchVoteQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := c.Param("vote")
	q, err := s.Store.QuestionFind(id)
	if err != nil {
		JSONNotFound("Error question not found. ", err, c)
		return
	}
	if v == "upvote" {
		q.Vote++
	} else if v == "downvote" {
		q.Vote--
	}

	err = s.Store.QuestionVoteUpdate(q)
	if err != nil {
		JSONInternalServer("Error in question voting. ", err, c)
		return
	}

	titleSlug := slug.Make(q.Title)
	c.Redirect(301, config.DOMAIN+"/questions/"+c.Param("id")+"/"+titleSlug)
}

// GetQuestionList returns a list of questions.
func (s *Server) GetQuestionList(c *gin.Context) {
	list, err := s.Store.QuestionsList(c.Request.URL.Query())

	if err != nil {
		JSONNotFound("Error finding questions list. ", err, c)
		return
	}

	c.JSON(200, list)
}

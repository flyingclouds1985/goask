package server

import (
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

// GetQuestionCommentList returns a list consists of comments for the question.
func (s *Server) GetQuestionCommentList(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Set("question_id", c.Param("question_id"))
	list, err := s.Store.QuestionCommentList(query)

	if err != nil {
		JSONBadRequestError("Error in getting comments list. ", err, c)
	}

	c.JSON(200, list)
}

// PostQuestionComment creates a comment for the question.
func (s *Server) PostQuestionComment(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Comment)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding comment. ", err, c)
	}

	comment := new(model.Comment)
	comment.Body = in.Body
	// comment.AuthorID = claims["id"]
	qid, _ := strconv.Atoi(c.Param("question_id"))
	comment.TrackableId = qid
	comment.TrackableType = "Question"

	if err = s.Store.CommentCreate(comment); err != nil {
		JSONBadRequestError("Error in inserting comment. ", err, c)
	}

	c.JSON(200, comment)
}

// GetReplyCommentList returns a list consists of comments for the reply.
func (s *Server) GetReplyCommentList(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Set("reply_id", c.Param("reply_id"))
	list, err := s.Store.ReplyCommentList(query)

	if err != nil {
		JSONBadRequestError("Error in getting reply comment list. ", err, c)
	}

	c.JSON(200, list)
}

// PostReplyComment creates a comment for the reply.
func (s *Server) PostReplyComment(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Comment)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding comment. ", err, c)
	}

	comment := new(model.Comment)
	comment.Body = in.Body
	// comment.AuthorID = claims["id"]
	rid, _ := strconv.Atoi(c.Param("reply_id"))
	comment.TrackableId = rid
	comment.TrackableType = "Reply"

	if err = s.Store.CommentCreate(comment); err != nil {
		JSONBadRequestError("Error in inserting comment. ", err, c)
	}

	c.JSON(200, comment)
}

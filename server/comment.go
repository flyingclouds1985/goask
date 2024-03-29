package server

import (
	"github.com/Alireza-Ta/goask/model"
	"github.com/Alireza-Ta/goask/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

//CommentStore manages encapsulated database access.
type CommentStore interface {
	QuestionCommentList(query url.Values) (model.Comments, error)
	CommentCreate(c *model.Comment) error
	ReplyCommentList(query url.Values) (model.Comments, error)
}

//CommentAPI provides handler for managing comments.
type CommentAPI struct {
	store CommentStore
}

// GetQuestionCommentList returns a list consists of comments for the question.
func (capi *CommentAPI) GetQuestionCommentList(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Set("question_id", c.Param("question_id"))
	list, err := capi.store.QuestionCommentList(query)
	if err != nil {
		JSONNotFound("Error finding question comments list. ", err, c)
		return
	}

	c.JSON(http.StatusOK, list)
}

// PostQuestionComment creates a comment for the question.
func (capi *CommentAPI) PostQuestionComment(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Comment)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	comment := new(model.Comment)
	comment.Body = in.Body
	// comment.AuthorID = claims["id"]
	qid, _ := strconv.Atoi(c.Param("question_id"))
	comment.TrackableId = qid
	comment.TrackableType = "Question"

	if err := capi.store.CommentCreate(comment); err != nil {
		JSONInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetReplyCommentList returns a list consists of comments for the reply.
func (capi *CommentAPI) GetReplyCommentList(c *gin.Context) {
	query := c.Request.URL.Query()
	query.Set("reply_id", c.Param("reply_id"))
	list, err := capi.store.ReplyCommentList(query)

	if err != nil {
		JSONInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, list)
}

// PostReplyComment creates a comment for the reply.
func (capi *CommentAPI) PostReplyComment(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Comment)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	comment := new(model.Comment)
	comment.Body = in.Body
	// comment.AuthorID = claims["id"]
	rid, _ := strconv.Atoi(c.Param("reply_id"))
	comment.TrackableId = rid
	comment.TrackableType = "Reply"

	if err := capi.store.CommentCreate(comment); err != nil {
		JSONInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, comment)
}

package api

import (
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

func (a *Api) GetQuestionCommentList(c *gin.Context) {
	list, err := a.Store.CommentList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError("Error in getting comments list. ", err, c)
	}

	c.JSON(200, list)
}

func (a *Api) PostQuestionComment(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Comment)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding comment. ", err, c)
	}

	comment := new(model.Comment)
	comment.Body = in.Body
	// comment.AuthorID = claims["id"]

	if err = a.Store.CommentCreate(comment); err != nil {
		JSONBadRequestError("Error in inserting comment. ", err, c)
	}

	qid, _ := strconv.Atoi(c.Param("question_id"))
	cq := new(model.CommentsQuestion)
	cq.CommentId = comment.Id
	cq.QuestionId = qid

	if err = a.Store.CommentQuestionCreate(cq); err != nil {
		JSONBadRequestError("Error in inserting comment_question relation. ", err, c)
	}

	c.JSON(200, comment)
}

func (a *Api) GetReplyCommentList(c *gin.Context) {
	c.JSON(200, "okok")
}

func (a *Api) PostReplyComment(c *gin.Context) {
	c.JSON(200, "okok")
}

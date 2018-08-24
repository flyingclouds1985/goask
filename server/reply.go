package server

import (
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

// GetReplyList returns list of replies.
func (s *Server) GetReplyList(c *gin.Context) {
	list, err := s.Store.ReplyList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError(ListErr("reply"), err, c)
		return
	}

	c.JSON(200, list)
}

// PostReply creates a reply.
func (s *Server) PostReply(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Reply)
	err := c.ShouldBindJSON(in)

	if err != nil {
		JSONBadRequestError(BindErr("reply"), err, c)
		return
	}
	qid, _ := strconv.Atoi(c.Param("question_id"))
	r := new(model.Reply)
	r.Body = in.Body
	r.QuestionId = qid
	// r.AuthorID = claims["id"]

	if err = s.Store.ReplyCreate(r); err != nil {
		JSONBadRequestError(InsertErr("reply"), err, c)
		return
	}

	c.JSON(200, r)
}

// PatchReply updates a reply.
func (s *Server) PatchReply(c *gin.Context) {
	in := new(model.Reply)
	err := c.ShouldBindJSON(in)

	if in.Id == 0 {
		rid, _ := strconv.Atoi(c.Param("reply_id"))
		in.Id = rid
	}

	if err != nil {
		JSONBadRequestError(BindErr("reply"), err, c)
		return
	}

	if err = s.Store.ReplyUpdate(in); err != nil {
		JSONBadRequestError(UpdateErr("reply"), err, c)
		return
	}

	c.JSON(200, in)
}

package server

import (
	"net/http"
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/validation"
	"github.com/gin-gonic/gin"
)

type ReplyAPI struct {
	store *postgres.Store
}

// GetReplyList returns list of replies.
func (rapi *ReplyAPI) GetReplyList(c *gin.Context) {
	list, err := rapi.store.ReplyList(c.Request.URL.Query())

	if err != nil {
		JSONNotFound("Error finding replies list. ", err, c)
		return
	}

	c.JSON(http.StatusOK, list)
}

// PostReply creates a reply.
func (rapi *ReplyAPI) PostReply(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Reply)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	qid, _ := strconv.Atoi(c.Param("question_id"))
	r := new(model.Reply)
	r.Body = in.Body
	r.QuestionId = qid
	// r.AuthorID = claims["id"]

	if err := rapi.store.ReplyCreate(r); err != nil {
		JSONInternalServer("Error inserting reply. ", err, c)
		return
	}

	c.JSON(http.StatusOK, r)
}

// PatchReply updates a reply.
func (rapi *ReplyAPI) PatchReply(c *gin.Context) {
	in := new(model.Reply)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	// if in.Id == 0 {
	// 	rid, _ := strconv.Atoi(c.Param("reply_id"))
	// 	in.Id = rid
	// }

	if err := rapi.store.ReplyUpdate(in); err != nil {
		JSONInternalServer("Error updating reply. ", err, c)
		return
	}

	c.JSON(http.StatusOK, in)
}

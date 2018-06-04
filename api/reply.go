package api

import (
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

func (a *Api) GetReplyList(c *gin.Context) {
	list, err := a.Store.ReplyList(c.Request.URL.Query())

	if err != nil {
		JSONBadRequestError("Error in getting reply list. ", err, c)
	}

	c.JSON(200, list)
}

func (a *Api) PostReply(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	in := new(model.Reply)
	err := c.ShouldBind(in)

	if err != nil {
		JSONBadRequestError("Error in binding reply. ", err, c)
	}

	r := new(model.Reply)
	r.Body = in.Body
	// r.AuthorID = claims["id"]

	if err = a.Store.ReplyCreate(r); err != nil {
		JSONBadRequestError("Error in inserting reply. ", err, c)
	}

	c.JSON(200, r)
}

func (a *Api) PatchReply(c *gin.Context) {
	in := new(model.Reply)
	err := c.ShouldBind(in)

	if in.Id == 0 {
		rid, _ := strconv.Atoi(c.Param("reply_id"))
		in.Id = rid
	}

	if err != nil {
		JSONBadRequestError("Error in binding reply. ", err, c)
	}

	if err = a.Store.ReplyUpdate(in); err != nil {
		JSONBadRequestError("Error in updating question. ", err, c)
	}

	c.JSON(200, in)
}

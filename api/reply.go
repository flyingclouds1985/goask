package api

import (
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

	if err = a.Store.CreateReply(r); err != nil {
		JSONBadRequestError("Error in inserting reply. ", err, c)
	}

	c.JSON(200, r)
}

package server

import (
	"errors"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/gin-gonic/gin"
)

var PasswordHashErr = errors.New("Problem in creating user, try again!")

func (s *Server) GetUser(c *gin.Context) {
	username := c.Param("username")
	u, err := s.Store.UserFind(username)
	if err != nil {
		JSONNotFoundError(NotFoundErr("user"), err, c)
		return
	}

	c.JSON(200, u.Copy())
}

func (s *Server) PostUser(c *gin.Context) {
	in := new(model.User)

	if err := c.ShouldBindJSON(in); err != nil {
		JSONBadRequestError(BindErr("user"), err, c)
		return
	}

	u := new(model.User)
	u.Username = in.Username
	pass, err := HashPassword(in.Password)
	if err != nil {
		JSONBadRequestError(PasswordHashErr, err, c)
		return
	}
	u.Password = pass
	u.Email = in.Email
	u.Bio = in.Bio
	if err = u.ValidateUsername(); err != nil {
		JSONBadRequestError(InsertErr("user"), err, c)
		return
	}

	if err = s.Store.UserCreate(u); err != nil {
		JSONBadRequestError(InsertErr("user"), err, c)
		return
	}

	c.JSON(200, u.Copy())
}

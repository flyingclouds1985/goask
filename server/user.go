package server

import (
	"net/http"

	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/Alireza-Ta/GOASK/validation"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	store *postgres.Store
}

// GetUser responds user by username.
func (uapi *UserAPI) GetUser(c *gin.Context) {
	username := c.Param("username")
	u, err := uapi.store.UserFindByName(username)
	if err != nil {
		JSONNotFound("Error user not found. ", err, c)
		return
	}
	c.JSON(http.StatusOK, u.Copy())
}

// PostUser create new user.
func (uapi *UserAPI) PostUser(c *gin.Context) {
	in := new(model.User)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}
	if err := in.Validate(); err != nil {
		JSONBadRequest("Error inserting user. ", err, c)
		return
	}

	u := new(model.User)
	u.Username = in.Username
	pass, err := HashPassword(in.Password)
	if err != nil {
		JSONInternalServer("Error inserting user. ", err, c)
		return
	}
	u.Password = pass
	u.Email = in.Email
	u.Bio = in.Bio

	if err = uapi.store.UserCreate(u); err != nil {
		JSONInternalServer("Error inserting user. ", err, c)
		return
	}

	c.JSON(http.StatusOK, u.Copy())
}

// PatchUser updates user.
func (uapi *UserAPI) PatchUser(c *gin.Context) {
	in := new(model.User)
	if err := c.ShouldBindJSON(in); err != nil {
		JSONValidation(validation.Messages(err), c)
		return
	}

	u, err := uapi.store.UserFind(in.Id)
	if err != nil {
		JSONInternalServer("Error finding user. ", err, c)
	}

	u.Username = in.Username
	u.Email = in.Email
	// for now we have no validation for Bio.
	if in.Bio != "" {
		u.Bio = in.Bio
	}

	if _, err := uapi.store.UserUpdateExcludePassword(u); err != nil {
		JSONInternalServer("Error updating user. ", err, c)
		return
	}

	c.JSON(http.StatusOK, u.Copy())
}

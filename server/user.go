package server

import (
	"github.com/Alireza-Ta/goask/model"
	"github.com/Alireza-Ta/goask/validation"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

//UserStore manages encapsulated database access.
type UserStore interface {
	CreateUser(user *model.User) error
	FindUser(id int) (*model.User, error)
	FindUserByName(username string) (*model.User, error)
	UpdateUserExcludePassword(user *model.User) (int, error)
	FindUserByLoginCredentials(username, password string) (*model.User, error)
}

//UserAPI provides handlers for managing users.
type UserAPI struct {
	store UserStore
}

// GetUser responds user by username.
func (uapi *UserAPI) GetUser(c *gin.Context) {
	username := c.Param("username")
	u, err := uapi.store.FindUserByName(username)
	if err != nil {
		JSONNotFound("Error user not found. ", err, c)
		return
	}
	c.JSON(http.StatusOK, u.Copy())
}

// PostUser create new user.
func (uapi *UserAPI) PostUser(c *gin.Context) {
	in := new(model.User)
	err := c.ShouldBindJSON(in)
	validationErr := validator.ValidationErrors{}

	if ve, ok := err.(validator.ValidationErrors); ok != true {
		for k, v := range ve {
			validationErr[k] = v
		}
	}

	requireFields := map[string]string{
		"Username":        in.Username,
		"Email":           in.Email,
		"Password":        in.Password,
		"ConfirmPassword": in.ConfirmPassword,
	}

	for k, v := range requireFields {
		if v == "" {
			validationErr[k] = &validator.FieldError{
				FieldNamespace: k,
				Field:          k,
				Name:           k,
				NameNamespace:  k,
				Tag:            "required",
				ActualTag:      "required",
				Kind:           24,
			}
		}
	}

	if len(validationErr) != 0 {
		JSONValidation(validation.Messages(validationErr), c)
		return
	}

	// TODO : check it and delete
	// if err := in.Validate(); err != nil {
	// 	JSONBadRequest("Error inserting user. ", err, c)
	// 	return
	// }

	u := new(model.User)
	u.Username = in.Username
	pass, err := HashPassword(in.Password)
	if err != nil {
		JSONInternalServer(err, c)
		return
	}
	u.Password = pass
	u.Email = in.Email
	u.Bio = in.Bio

	if err = uapi.store.CreateUser(u); err != nil {
		JSONInternalServer(err, c)
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

	u, err := uapi.store.FindUser(in.Id)
	if err != nil {
		JSONInternalServer(err, c)
	}

	u.Username = in.Username
	u.Email = in.Email
	// for now we have no validation for Bio.
	if in.Bio != "" {
		u.Bio = in.Bio
	}

	//if _, err := uapi.store.UpdateUserExcludePassword(u); err != nil {
	//	JSONInternalServer("Error updating user. ", err, c)
	//	return
	//}
	rowsAffected, err := uapi.store.UpdateUserExcludePassword(u)
	if err != nil {
		JSONInternalServer(err, c)
		return
	}

	if rowsAffected < 0 {
		NoRowsAffected := errors.New("No Rows Affected")
		JSONInternalServer(NoRowsAffected, c)
	}

	c.JSON(http.StatusOK, u.Copy())
}

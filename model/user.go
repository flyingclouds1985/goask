package model

import (
	"errors"
	"regexp"
	"time"
)

var (
	errUsernameRegex    = errors.New("Invalid username.It must start with alphabet")
	errPasswordMisMatch = errors.New("Password and confirm password mismatch.")
	regexUsername       = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-_.]+$")
)

// User model
type User struct {
	Id              int       `json:"id"`
	Username        string    `json:"username" sql:",unique" binding:"min=5,max=32"`
	Email           string    `json:"email" sql:",unique" binding:"omitempty,email"`
	Password        string    `json:"password" binding:"omitempty,min=8,max=64,eqfield=ConfirmPassword"`
	ConfirmPassword string    `json:"confirmPassword" sql:"-"`
	Bio             string    `json:"bio"`
	CreatedAt       time.Time `json:"created_at" sql:"type:timestamptz,default:now()"`
	UpdatedAt       time.Time `json:"updated_at" sql:"type:timestamptz,default:now()"`
}

// Copy makes a copy of the user without password and confirmPassword.
func (u *User) Copy() *User {
	return &User{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Bio:       u.Bio,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ExcludeTimes makes a copy of the user without timestamps, password and confirmPassword.
func (u *User) ExcludeTimes() *User {
	return &User{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
	}
}

// Validate validates the credentials.
func (u *User) Validate() error {
	switch {
	case !regexUsername.MatchString(u.Username):
		return errUsernameRegex
	case u.Password != u.ConfirmPassword:
		return errPasswordMisMatch
	default:
		return nil
	}
}

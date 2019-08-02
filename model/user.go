package model

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/go-pg/pg/v9/orm"
)

var (
	errUsernameRegex    = errors.New("Invalid username.It must start with alphabet")
	errPasswordRequired = errors.New("Password field required")
	regexUsername       = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-_.]+$")
	_ orm.BeforeInsertHook = (*User)(nil)
	_ orm.BeforeInsertHook = (*User)(nil)
)

// User model
type User struct {
	Id              int       `json: "id"`
	Username        string    `json:"username" binding:"required,min=5,max=32"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"omitempty,min=8,max=64,eqfield=ConfirmPassword"`
	ConfirmPassword string    `json:"confirmPassword" sql:"-"`
	Bio             string    `json:"bio"`
	CreatedAt       time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt       time.Time `json:"updated_at" sql:"type:timestamptz"`
}

// BeforeInsert runs before every insert.(orm hook)
func (u *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	u.CreatedAt = UnixTime()
	u.UpdatedAt = UnixTime()
	return ctx, nil
}

// BeforeUpdate runs before every update.(orm hook)
func (u *User) BeforeUpdate(ctx context.Context) (context.Context, error) {
	u.UpdatedAt = UnixTime()
	if u.CreatedAt.IsZero() {
		data := new(User)
		data.Id = u.Id
		var db orm.DB
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding User created_at column.", err.Error())
		}
		u.CreatedAt = data.CreatedAt
	}
	return ctx, nil
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
	case u.Password == "":
		return errPasswordRequired
	default:
		return nil
	}
}

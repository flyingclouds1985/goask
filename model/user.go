package model

import (
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/go-pg/pg/orm"
)

var (
	errUsernameRegex = errors.New("Invalid username.It must start with alphabet")
	regexUsername    = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-_.]+$")
)

// User model
type User struct {
	Id        int       `json: "id"`
	Username  string    `json:"username" binding:"required,min=5,max=32"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=8,max=64"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

// BeforeInsert runs before every insert.(orm hook)
func (u *User) BeforeInsert(db orm.DB) error {
	u.CreatedAt = UnixTime()
	u.UpdatedAt = UnixTime()
	return nil
}

// BeforeUpdate runs before every update.(orm hook)
func (u *User) BeforeUpdate(db orm.DB) error {
	u.UpdatedAt = UnixTime()
	if u.CreatedAt.IsZero() {
		data := new(User)
		data.Id = u.Id
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding User created_at column.", err.Error())
		}
		u.CreatedAt = data.CreatedAt
	}
	return nil
}

// Copy makes a copy of the user without password.
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

// ExcludeTimes makes a copy of the user without timestamps and password.
func (u *User) ExcludeTimes() *User {
	return &User{
		Id:       u.Id,
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
	}
}

// ValidateUsername validates the username in regex format.
func (u *User) ValidateUsername() error {
	if !regexUsername.MatchString(u.Username) {
		return errUsernameRegex
	}
	return nil
}

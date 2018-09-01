package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	Id        int       `json: "id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (u *User) BeforeInsert(db orm.DB) error {
	u.CreatedAt = UnixTime()
	u.UpdatedAt = UnixTime()
	return nil
}

// BeforeUpdate user
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

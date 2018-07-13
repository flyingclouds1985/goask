package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type User struct {
	Id        int       `json: "id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (u *User) BeforeInsert(db orm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

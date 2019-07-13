package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type Comments []Comment

type Comment struct {
	Id            int       `json: "id"`
	Body          string    `json:"body" sql:"type:text" binding:"required,min=2"`
	AuthorID      int       `json:"author_id"`
	Author        *User     `json:"author"`
	Vote          int       `json:"vote" sql:"default:0"`
	TrackableId   int       `json:"trackable_id"`
	TrackableType string    `json:"trackable_type"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt     time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (c *Comment) BeforeInsert(db orm.DB) error {
	c.CreatedAt = UnixTime()
	c.UpdatedAt = UnixTime()
	return nil
}

func (c *Comment) BeforeUpdate(db orm.DB) error {
	c.UpdatedAt = UnixTime()
	if c.CreatedAt.IsZero() {
		data := new(Comment)
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding comment created_at column.", err.Error())
		}
		c.CreatedAt = data.CreatedAt
	}
	return nil
}

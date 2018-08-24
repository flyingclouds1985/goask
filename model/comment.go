package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type CommentStore interface {
	CreateComment(*Comment) error
}

type Comment struct {
	Id            int    `json: "id"`
	TrackableId   int    `json:"trackable_id"`
	TrackableType string `json:"trackable_type"`
	Post          `pg:"override"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt     time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (c *Comment) BeforeInsert(db orm.DB) error {
	c.UpdatedAt = UnixTime()
	return nil
}

func (c *Comment) BeforeUpdate(db orm.DB) error {
	c.UpdatedAt = time.Now()
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

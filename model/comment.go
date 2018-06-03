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
	Id        int `json: "id"`
	Post      `pg:"override"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

type CommentsQuestion struct {
	CommentId  int `json:"comment_id"`
	QuestionId int `json:"question_id"`
}

type CommentsReply struct {
	CommentId int `json:"comment_id"`
	ReplyId   int `json:"reply_id"`
}

func (c *Comment) BeforeInsert(db orm.DB) error {
	c.UpdatedAt = time.Now()
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

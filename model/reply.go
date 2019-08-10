package model

import (
	"time"
)

type Replies []Reply

type Reply struct {
	Id         int       `json:"id"`
	Body       string    `json:"body" sql:"type:text" binding:"required,min=10"`
	AuthorID   int       `json:"author_id"`
	Author     *User     `json:"author"`
	Vote       int       `json:"vote" sql:"default:0"`
	Approved   int       `json:"approved"`
	QuestionId int       `json:"question_id"`
	Comments   []Comment `json:"comments" pg:"polymorphic:trackable_"`
	CreatedAt  time.Time `json:"created_at" sql:"type:timestamptz,default:now()"`
	UpdatedAt  time.Time `json:"updated_at" sql:"type:timestamptz,default:now()"`
}
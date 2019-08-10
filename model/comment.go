package model

import (
	"time"
)

type Comments []Comment

type Comment struct {
	Id            int       `json:"id"`
	Body          string    `json:"body" sql:"type:text" binding:"required,min=2"`
	AuthorID      int       `json:"author_id"`
	Author        *User     `json:"author"`
	Vote          int       `json:"vote" sql:"default:0"`
	TrackableId   int       `json:"trackable_id"`
	TrackableType string    `json:"trackable_type"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamptz,default:now()"`
	UpdatedAt     time.Time `json:"updated_at" sql:"type:timestamptz,default:now()"`
}

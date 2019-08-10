package model

import (
	"time"
)

// Tag Model
type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	QuestionId int
	CreatedAt  time.Time `json:"created_at" sql:"type:timestamptz,default:now()"`
	UpdatedAt  time.Time `json:"updated_at" sql:"type:timestamptz,default:now()"`
}

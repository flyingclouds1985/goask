package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Reply struct {
	Id        int `json: "id"`
	Post      `pg:"override"`
	Approved  int       `json:"approved"`
	Comments  []Comment `json:"comments" pg:"many2many:comments_replies"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (r *Reply) BeforeInsert(db orm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

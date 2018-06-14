package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type Reply struct {
	Id         int `json: "id"`
	Post       `pg:"override"`
	Approved   int       `json:"approved"`
	QuestionId int       `json:"question_id"`
	Comments   []Comment `json:"comments" pg:"polymorphic:trackable_"`
	CreatedAt  time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt  time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (r *Reply) BeforeInsert(db orm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Reply) BeforeUpdate(db orm.DB) error {
	r.UpdatedAt = time.Now()
	if r.CreatedAt.IsZero() {
		data := new(Reply)
		data.Id = r.Id
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding reply created_at column.", err.Error())
		}
		r.CreatedAt = data.CreatedAt
	}
	return nil
}

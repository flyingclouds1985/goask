package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type Replies []Reply

type Reply struct {
	Id         int       `json: "id"`
	Body       string    `json:"body" sql:"type:text" binding:"required,min=10"`
	AuthorID   int       `json:"author_id"`
	Author     *User     `json:"author"`
	Vote       int       `json:"vote" sql:"default:0"`
	Approved   int       `json:"approved"`
	QuestionId int       `json:"question_id"`
	Comments   []Comment `json:"comments" pg:"polymorphic:trackable_"`
	CreatedAt  time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt  time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (r *Reply) BeforeInsert(db orm.DB) error {
	r.CreatedAt = UnixTime()
	r.UpdatedAt = UnixTime()
	return nil
}

func (r *Reply) BeforeUpdate(db orm.DB) error {
	r.UpdatedAt = UnixTime()
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

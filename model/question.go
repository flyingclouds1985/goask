package model

import (
	"log"
	"net/url"
	"time"

	"github.com/go-pg/pg/orm"
)

type QuestionStore interface {
	QuestionCreate(*Question) error
	QuestionsList(url.Values) (*[]Question, error)
	QuestionFind(int) (*Question, error)
	QuestionUpdate(*Question) error
	QuestionVoteUpdaete(*Question) error
}

// Question model
type Question struct {
	Id        int `json:"id"`
	Post      `pg:"override"`
	Title     string    `json:"title"`
	Answered  int       `json:"answered" sql:"default:0"`
	Replies   []Reply   `json:"replies"`
	Comments  []Comment `json:"comments" pg:"polymorphic:trackable_"`
	Tags      []*Tag    `json:"tags"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (q *Question) BeforeInsert(db orm.DB) error {
	q.UpdatedAt = time.Now()
	return nil
}

func (q *Question) BeforeUpdate(db orm.DB) error {
	q.UpdatedAt = time.Now()
	if q.CreatedAt.IsZero() {
		data := new(Question)
		data.Id = q.Id
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding question created_at column.", err.Error())
		}
		q.CreatedAt = data.CreatedAt
	}
	return nil
}

package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

// Question model
type Question struct {
	Id        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Author    *User     `json:"author"`
	Vote      int       `json:"vote" sql:"default:0"`
	Title     string    `json:"title" binding:"required,min=15"`
	Body      string    `json:"body" sql:"type:text" binding:"required,min=50"`
	Answered  int       `json:"answered" sql:"default:0"`
	Replies   []Reply   `json:"replies"`
	Comments  []Comment `json:"comments" pg:"polymorphic:trackable_"`
	Tags      []*Tag    `json:"tags"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (q *Question) BeforeInsert(db orm.DB) error {
	q.CreatedAt = UnixTime()
	q.UpdatedAt = UnixTime()
	return nil
}

func (q *Question) BeforeUpdate(db orm.DB) error {
	q.UpdatedAt = UnixTime()
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

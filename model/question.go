package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type QuestionStore interface {
	CreateQuestion(*Question) error
}

type Question struct {
	Post     `pg:"override"`
	Title    string    `json:"title"`
	Comments []Comment `json:"comments" pg:"many2many:comments_questions"`
}

func (q *Question) BeforeUpdate(db orm.DB) error {
	q.UpdatedAt = time.Now()

	if q.CreatedAt.IsZero() {
		data := new(Question)
		err := db.Model(data).Column("created_at").Where("id = ?", q.Id).Select()
		if err != nil {
			log.Fatal("Error in finding question created_at column.", err.Error())
		}
		q.CreatedAt = data.CreatedAt
	}
	return nil
}

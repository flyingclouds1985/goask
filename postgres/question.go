package postgres

import (
	"github.com/Alireza-Ta/GOASK/model"
)

func CreateQuestion(q *model.Question) error {
	return db.Insert(q)
}

package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

// Questions is an array of questions.
type Questions []model.Question

// CreateQuestion persist a question in db.
func CreateQuestion(q *model.Question) error {
	return db.Insert(q)
}

// QuestionList returns a list of questions.
func QuestionList(query url.Values) (Questions, error) {
	var questions Questions

	err := db.Model(&questions).
		Apply(orm.Pagination(query)).
		Select()

	return questions, err
}

func QuestionFind(id int) (*model.Question, error) {
	q := &model.Question{Post: model.Post{Id: id}}
	err := db.Select(q)

	return q, err
}

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

// QuestionsList returns a list of questions.
func QuestionsList(query url.Values) (Questions, error) {
	var questions Questions

	err := db.Model(&questions).
		Apply(orm.Pagination(query)).
		Select()

	return questions, err
}

func QuestionFind(id int) (*model.Question, error) {
	q := new(model.Question)
	err := db.Model(q).Where("id = ?", id).Select()

	return q, err
}

func QuestionUpdate(q *model.Question) error {
	return db.Update(q)
}

func QuestionVoteUpdate(vote int, id int) error {
	_, err := db.Model(&model.Question{}).Set("vote = ?vote").Where("id = ?id").Update()
	return err
}

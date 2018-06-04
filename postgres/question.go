package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

// Questions is an array of question.
type Questions []model.Question

// CreateQuestion persist a question in db.
func (s *Store) CreateQuestion(q *model.Question) error {
	return s.db.Insert(q)
}

// QuestionsList returns a list of questions.
func (s *Store) QuestionsList(query url.Values) (Questions, error) {
	var questions Questions

	err := s.db.Model(&questions).
		Apply(orm.Pagination(query)).
		Select()

	return questions, err
}

func (s *Store) QuestionFind(id int) (*model.Question, error) {
	q := new(model.Question)
	err := s.db.Model(q).Where("id = ?", id).Select()

	return q, err
}

func (s *Store) QuestionUpdate(q *model.Question) error {
	return s.db.Update(q)
}

func (s *Store) QuestionVoteUpdate(q *model.Question) error {
	_, err := s.db.Model(q).Column("vote", "updated_at").WherePK().Update()
	return err
}

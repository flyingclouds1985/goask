package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

// Questions is an array of question.
type Questions []model.Question

// CreateQuestion persists a question in db.
func (s *Store) QuestionCreate(q *model.Question) error {
	return s.DB.Insert(q)
}

// QuestionsList returns a list of questions.
func (s *Store) QuestionsList(query url.Values) (Questions, error) {
	var q Questions

	err := s.DB.Model(&q).
		Apply(orm.Pagination(query)).
		Relation("Replies").
		Relation("Comments").
		Relation("Tags").
		Select()

	return q, err
}

func (s *Store) QuestionWithRelations(id int) (*model.Question, error) {
	q := new(model.Question)

	err := s.DB.Model(q).Where("id = ?", id).Relation("Replies").Relation("Comments").Select()

	return q, err
}

func (s *Store) QuestionFind(id int) (*model.Question, error) {
	q := new(model.Question)

	err := s.DB.Model(q).Where("id = ?", id).Select()

	return q, err
}

func (s *Store) QuestionUpdate(q *model.Question) error {
	return s.DB.Update(q)
}

func (s *Store) QuestionVoteUpdate(q *model.Question) error {
	_, err := s.DB.Model(q).Column("vote", "updated_at").WherePK().Update()
	return err
}

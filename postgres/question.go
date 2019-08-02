package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/goask/model"
	"github.com/go-pg/pg/urlvalues"
)

// CreateQuestion persists a question in db.
func (s *Store) CreateQuestion(q *model.Question) error {
	return s.DB.Insert(q)
}

// ListQuestion returns a list of questions.
func (s *Store) ListQuestion(query url.Values) (model.Questions, error) {
	var q model.Questions
	err := s.DB.Model(&q).
		Apply(urlvalues.Pagination(urlvalues.Values(query))).
		Relation("Replies").
		Relation("Comments").
		Relation("Tags").
		Select()

	return q, err
}

// QuestionWithRelations returns question with its relations.
func (s *Store) QuestionWithRelations(id int) (*model.Question, error) {
	q := new(model.Question)

	err := s.DB.Model(q).Where("id = ?", id).Relation("Replies").Relation("Comments").Relation("Tags").Select()

	return q, err
}

// FindQuestion finds question by its id.
func (s *Store) FindQuestion(id int) (*model.Question, error) {
	q := new(model.Question)

	err := s.DB.Model(q).Where("id = ?", id).Select()

	return q, err
}

// UpdateQuestion updates the question.
func (s *Store) UpdateQuestion(q *model.Question) error {
	return s.DB.Update(q)
}

// UpdateVote updates question's vote.
func (s *Store) UpdateVote(q *model.Question) error {
	_, err := s.DB.Model(q).Column("vote", "updated_at").WherePK().Update()
	return err
}

package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

type Comments []model.Comment

func (s *Store) CommentList(query url.Values) (Comments, error) {
	var comments Comments

	err := s.db.Model(&comments).
		Apply(orm.Pagination(query)).
		Select()

	return comments, err
}

func (s *Store) CommentCreate(c *model.Comment) error {
	return s.db.Insert(c)
}

func (s *Store) CommentQuestionCreate(cq *model.CommentsQuestion) error {
	return s.db.Insert(cq)
}

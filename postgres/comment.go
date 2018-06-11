package postgres

import (
	"net/url"

	"github.com/go-pg/pg"

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

func (s *Store) CommentCreate(c *model.Comment, question_id int) error {
	return s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Insert(c)
		if err != nil {
			return err
		}

		cq := new(model.CommentsQuestion)
		cq.CommentId = c.Id
		cq.QuestionId = question_id

		err = tx.Insert(cq)

		return err
	})
}

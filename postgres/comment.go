package postgres

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-pg/pg"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

type Comments []model.Comment

func (s *Store) QuestionCommentList(query url.Values) (Comments, error) {
	var question model.Question
	qid, _ := strconv.Atoi(query.Get("question_id"))
	fmt.Println(query)
	p, err := s.pagination(&question, query, qid, "Comments")

	return p.(*model.Question).Comments, err
}

func (s *Store) ReplyCommentList(query url.Values) (Comments, error) {
	var reply model.Reply
	rid, _ := strconv.Atoi(query.Get("reply_id"))
	p, err := s.pagination(&reply, query, rid, "Comments")

	return p.(*model.Reply).Comments, err
}

func (s *Store) pagination(model interface{}, query url.Values, id int, relation string) (interface{}, error) {
	err := s.db.Model(model).
		Apply(orm.Pagination(query)).
		Where("id = ?", id).
		Relation(relation).
		Select()

	return model, err
}

func (s *Store) QuestionCommentCreate(c *model.Comment, question_id int) error {
	return s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Insert(c)
		if err != nil {
			return err
		}

		cq := new(model.CommentsQuestion)
		cq.CommentId = c.Id
		cq.QuestionId = question_id

		err = tx.Insert(cq)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *Store) ReplyCommentCreate(c *model.Comment, reply_id int) error {
	return s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Insert(c)
		if err != nil {
			return err
		}

		cr := new(model.CommentsReply)
		cr.CommentId = c.Id
		cr.ReplyId = reply_id

		err = tx.Insert(cr)
		if err != nil {
			return err
		}

		return nil
	})
}

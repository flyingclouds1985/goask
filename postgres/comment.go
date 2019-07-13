package postgres

import (
	"net/url"
	"strconv"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

func (s *Store) QuestionCommentList(query url.Values) (model.Comments, error) {
	var question model.Question

	qid, _ := strconv.Atoi(query.Get("question_id"))
	p, err := s.pagination(&question, query, qid, "Comments")

	return p.(*model.Question).Comments, err
}

func (s *Store) ReplyCommentList(query url.Values) (model.Comments, error) {
	var reply model.Reply

	rid, _ := strconv.Atoi(query.Get("reply_id"))
	p, err := s.pagination(&reply, query, rid, "Comments")

	return p.(*model.Reply).Comments, err
}

func (s *Store) pagination(model interface{}, query url.Values, id int, relation string) (interface{}, error) {
	err := s.DB.Model(model).
		Apply(orm.Pagination(query)).
		Where("id = ?", id).
		Relation(relation).
		Select()

	return model, err
}

func (s *Store) CommentCreate(c *model.Comment) error {
	return s.DB.Insert(c)
}

package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

type Replies []model.Reply

func (s *Store) ReplyList(query url.Values) (Replies, error) {
	var replies Replies

	err := s.db.Model(&replies).
		Apply(orm.Pagination(query)).
		Select()

	return replies, err
}

func (s *Store) CreateReply(r *model.Reply) error {
	return s.db.Insert(r)
}

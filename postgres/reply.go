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
		Relation("Comments").
		Select()

	return replies, err
}

func (s *Store) ReplyCreate(r *model.Reply) error {
	return s.db.Insert(r)
}

func (s *Store) ReplyUpdate(r *model.Reply) error {
	return s.db.Update(r)
}

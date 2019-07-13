package postgres

import (
	"net/url"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg/orm"
)

// ListReply returns a list of replies.
func (s *Store) ListReply(query url.Values) (model.Replies, error) {
	var replies model.Replies

	err := s.DB.Model(&replies).
		Apply(orm.Pagination(query)).
		Relation("Comments").
		Select()

	return replies, err
}

// CreateReply creates the reply.
func (s *Store) CreateReply(r *model.Reply) error {
	return s.DB.Insert(r)
}

// UpdateReply updates the reply.
func (s *Store) UpdateReply(r *model.Reply) error {
	return s.DB.Update(r)
}

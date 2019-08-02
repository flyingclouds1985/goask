package postgres

import (
	"github.com/go-pg/pg/urlvalues"
	"net/url"
	"time"

	"github.com/Alireza-Ta/goask/model"
)

// ListReply returns a list of replies.
func (s *Store) ListReply(query url.Values) (model.Replies, error) {
	var replies model.Replies

	err := s.DB.Model(&replies).
		Apply(urlvalues.Pagination(urlvalues.Values(query))).
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
	r.UpdatedAt = time.Now()
	return s.DB.Update(r)
}

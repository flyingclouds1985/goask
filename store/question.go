package store

import (
	"github.com/Alireza-Ta/GOASK/postgres"
)

func (db *Store) CreateQuestion(q *postgres.Question) error {
	return db.Insert(q)
}

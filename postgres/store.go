package postgres

import (
	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var models = []interface{}{
	&model.User{},
	&model.Comment{},
	&model.Question{},
	&model.Reply{},
	&model.Tag{},
}

type Store struct {
	DB *pg.DB
}

func New(username, password, dbname string) *Store {
	// Don't forget to fill password field.
	return &Store{
		DB: openDB(username, password, dbname),
	}
}

func openDB(username, password, dbname string) *pg.DB {
	return pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: dbname,
	})
}

// CreateSchema create tables.
func (s *Store) CreateSchema() error {
	for _, model := range models {
		err := s.DB.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

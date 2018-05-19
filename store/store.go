package store

import (
	"log"

	"github.com/Alireza-Ta/GOASK/postgres"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Store struct {
	*pg.DB
}

func New() *pg.DB {
	// Don't forget to fill password field.
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "g",
	})

	return db
}

// CreateSchema create tables.
func CreateSchema() error {
	models := []interface{}{
		&postgres.User{},
		&postgres.Comment{},
		&postgres.Question{},
		&postgres.Reply{},
		&postgres.CommentsQuestion{},
		&postgres.CommentsReply{},
	}

	db := New()
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		})

		if err != nil {
			return err
		}
	}

	log.Print("tables created !")

	return nil
}

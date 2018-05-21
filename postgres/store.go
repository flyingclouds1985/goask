package postgres

import (
	"log"

	"github.com/Alireza-Ta/GOASK/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var db *pg.DB
var models = []interface{}{
	&model.User{},
	&model.Comment{},
	&model.Question{},
	&model.Reply{},
	&model.CommentsQuestion{},
	&model.CommentsReply{},
}

func init() {
	// Don't forget to fill password field.
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "g",
	})
}

// CreateSchema create tables.
func CreateSchema() error {
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

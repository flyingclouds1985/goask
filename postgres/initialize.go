package postgres

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func connection() *pg.DB {
	// don't forget to specify password field.
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "",
		Database: "GOASK_TEST",
	})

	return db
}

func CreateTables() error {
	models := []interface{}{
		&Author{},
		&Question{},
		&Reply{},
	}
	for _, model := range models {
		err := connection().CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

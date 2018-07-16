package model

import (
	"log"
	"time"

	"github.com/go-pg/pg/orm"
)

type Tag struct {
	Id        int       `json: "id"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (t *Tag) BeforeInsert(db orm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Tag) BeforeUpdate(db orm.DB) error {
	t.UpdatedAt = time.Now()
	if t.CreatedAt.IsZero() {
		data := new(Tag)
		data.Id = t.Id
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding Tag created_at column.", err.Error())
		}
		t.CreatedAt = data.CreatedAt
	}
	return nil
}

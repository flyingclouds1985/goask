package model

import (
	"context"
	"log"
	"time"

	"github.com/go-pg/pg/v9/orm"
)

var (
	_ orm.BeforeInsertHook = (*Tag)(nil)
	_ orm.BeforeInsertHook = (*Tag)(nil)
)

// Tag Model
type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	QuestionId int
	CreatedAt  time.Time `json:"created_at" sql:"type:timestamptz"`
	UpdatedAt  time.Time `json:"updated_at" sql:"type:timestamptz"`
}

// BeforeInsert tag
func (t *Tag) BeforeInsert(ctx context.Context) (context.Context, error) {
	t.CreatedAt = UnixTime()
	t.UpdatedAt = UnixTime()
	return ctx, nil
}

// BeforeUpdate tag
func (t *Tag) BeforeUpdate(ctx context.Context) (context.Context, error) {
	t.UpdatedAt = UnixTime()
	if t.CreatedAt.IsZero() {
		data := new(Tag)
		data.Id = t.Id
		var db orm.DB
		err := db.Model(data).Column("created_at").WherePK().Select()
		if err != nil {
			log.Fatal("Error in finding Tag created_at column.", err.Error())
		}
		t.CreatedAt = data.CreatedAt
	}
	return ctx, nil
}

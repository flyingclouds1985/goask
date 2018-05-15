package postgres

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type Post struct {
	Id        int
	Body      string
	AuthorID  int
	Author    *Author
	Vote      int
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func (p *Post) BeforeInsert(db orm.DB) error {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	return nil
}

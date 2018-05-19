package postgres

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type BaseModel struct {
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

func (bm *BaseModel) BeforeInsert(db orm.DB) error {
	bm.UpdatedAt = time.Now()

	return nil
}

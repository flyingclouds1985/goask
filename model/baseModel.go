package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type BaseModel struct {
	CreatedAt time.Time `json:"created_at" sql:"default:now()"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (bm *BaseModel) BeforeInsert(db orm.DB) error {
	bm.UpdatedAt = time.Now()

	return nil
}

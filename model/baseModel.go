package model

import (
	"time"

	"github.com/go-pg/pg/orm"
)

type BaseModel struct {
	Id        int       `json: "id"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamptz, default:now()"`
	UpdatedAt time.Time `json:"updated_at" sql:"type:timestamptz"`
}

func (bm *BaseModel) BeforeInsert(db orm.DB) error {
	bm.UpdatedAt = time.Now()
	return nil
}

package model

import (
	"time"
)

type Comments []Comment

type Comment struct {
	Id            int       `json:"id"`
	Body          string    `json:"body" sql:"type:text" binding:"required,min=2"`
	AuthorID      int       `json:"author_id"`
	Author        *User     `json:"author"`
	Vote          int       `json:"vote" sql:"default:0"`
	TrackableId   int       `json:"trackable_id"`
	TrackableType string    `json:"trackable_type"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamptz,default:now()"`
	UpdatedAt     time.Time `json:"updated_at" sql:"type:timestamptz,default:now()"`
}

//func (c *Comment) BeforeInsert(ctx context.Context) (context.Context, error) {
//	c.CreatedAt = UnixTime()
//	c.UpdatedAt = UnixTime()
//	return ctx, nil
//}
//
//func (c *Comment) BeforeUpdate(ctx context.Context) (context.Context, error) {
//	c.UpdatedAt = UnixTime()
//	if c.CreatedAt.IsZero() {
//		data := new(Comment)
//		var db orm.DB
//		err := db.Model(data).Column("created_at").WherePK().Select()
//		if err != nil {
//			log.Fatal("Error in finding comment created_at column.", err.Error())
//		}
//		c.CreatedAt = data.CreatedAt
//	}
//	return ctx, nil
//}

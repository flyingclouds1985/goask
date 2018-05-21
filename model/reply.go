package model

type Reply struct {
	Post     `pg:"override"`
	Approved int       `json:"approved"`
	Comments []Comment `json:"comments" pg:"many2many:comments_replies"`
}

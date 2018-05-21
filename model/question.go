package model

type QuestionStore interface {
	CreateQuestion(*Question) error
}

type Question struct {
	Post     `pg:"override"`
	Title    string    `json:"title"`
	Comments []Comment `json:"comments" pg:"many2many:comments_questions"`
}

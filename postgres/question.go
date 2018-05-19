package postgres

type QuestionStore interface {
	QuestionCreate(*Question) error
}

type Question struct {
	Post     `pg:"override"`
	Title    string
	Comments []Comment `pg:"many2many:comments_questions"`
}

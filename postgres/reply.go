package postgres

type Reply struct {
	Post     `pg:"override"`
	Approved int
	Comments []Comment `pg:"many2many:comments_replies"`
}

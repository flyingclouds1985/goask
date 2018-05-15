package postgres

type Question struct {
	Post  `pg:"override"`
	Title string
}

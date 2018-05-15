package postgres

type Reply struct {
	Post     `pg:"override"`
	Approved int
}

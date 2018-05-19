package postgres

type Post struct {
	Id        int
	Body      string
	AuthorID  int
	Author    *User
	Vote      int
	BaseModel `pg:"override"`
}

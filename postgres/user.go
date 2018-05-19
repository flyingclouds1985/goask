package postgres

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Bio       string
	BaseModel `pg:"override"`
}

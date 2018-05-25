package model

type User struct {
	BaseModel `pg:"override"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
}

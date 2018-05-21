package model

type Post struct {
	Id        int    `json: "id"`
	Body      string `json:"body"`
	AuthorID  int    `json:"author_id"`
	Author    *User  `json:"author"`
	Vote      int    `json:"vote" sql:"default:0"`
	BaseModel `pg:"override"`
}

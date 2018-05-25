package model

type Post struct {
	BaseModel `pg:"override"`
	AuthorID  int    `json:"author_id"`
	Author    *User  `json:"author"`
	Vote      int    `json:"vote" sql:"default:0"`
	Body      string `json:"body" sql:"type:text"`
}

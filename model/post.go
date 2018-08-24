package model

// Post model
type Post struct {
	AuthorID int    `json:"author_id"`
	Author   *User  `json:"author"`
	Vote     int    `json:"vote" sql:"default:0"`
	Body     string `json:"body" sql:"type:text" binding:"required,min=50"`
}

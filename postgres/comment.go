package postgres

type Comment struct {
	Post `pg:"override"`
}

type CommentsQuestion struct {
	CommentId  int
	QuestionId int
}

type CommentsReply struct {
	CommentId  int
	QuestionId int
}

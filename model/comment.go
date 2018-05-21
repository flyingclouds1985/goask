package model

type Comment struct {
	Post `pg:"override"`
}

type CommentsQuestion struct {
	CommentId  int `json:"comment_id"`
	QuestionId int `json:"question_id"`
}

type CommentsReply struct {
	CommentId int `json:"comment_id"`
	ReplyId   int `json:"reply_id"`
}

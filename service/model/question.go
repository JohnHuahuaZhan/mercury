package model

//UserInfo 目前全部not null
type BasicQuestionInfo struct {
	QuestionId uint64 `form:"question_id" json:"question_id"`
	Caption    string `form:"caption" json:"caption"`
	AuthorId   uint64 `form:"author_id" json:"author_id"`
	CategoryId uint64 `form:"category_id" json:"category_id"`
	Status     int32  `form:"status" json:"status"`
}
type QuestionInfo struct {
	BasicQuestionInfo
	Content string `form:"content" json:"content"`
}

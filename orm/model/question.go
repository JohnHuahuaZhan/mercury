package model

import "time"

//Question 除了时间其他均为not null。QuestionId有唯一性约束
type Question struct {
	ID         uint64    `db:"id"`
	QuestionId uint64    `db:"question_id"`
	Caption    string    `db:"caption"`
	Content    string    `db:"content"`
	AuthorId   uint64    `db:"author_id"`
	CategoryId uint64    `db:"category_id"`
	Status     int32     `db:"status"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

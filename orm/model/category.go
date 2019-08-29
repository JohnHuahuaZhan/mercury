package model

import "time"

//category_id category_name 有唯一性约束不为空
type Category struct {
	ID           uint32    `db:"id"`
	CategoryId   uint64    `db:"category_id"`
	CategoryName string    `db:"category_name"`
	CreateTime   time.Time `db:"create_time"`
	UpdateTime   time.Time `db:"update_time"`
}

package model

import "time"

//UserInfo 目前全部not null
type UserInfo struct {
	ID         uint64    `db:"id"`
	UserId     uint64    `db:"user_id"`
	Nickname   string    `db:"nickname"`
	Sex        int       `db:"sex"`
	Username   string    `db:"username"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

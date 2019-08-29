package model

//UserInfo 目前全部not null
type BasicUserInfo struct {
	Nickname string `form:"nickname" json:"nickname"`
	Sex      int    `form:"sex" json:"sex"`
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
}
type UserInfoSecret struct {
	BasicUserInfo
	Password string `form:"password" json:"password"`
	UserId   uint64 `form:"user_id" json:"user_id"`
}

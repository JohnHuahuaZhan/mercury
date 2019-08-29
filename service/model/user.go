package model

//UserInfo 目前全部not null
type BasicUserInfo struct {
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type UserInfoSecret struct {
	BasicUserInfo
	Password string `json:"password"`
	UserId   uint64 `json:"user_id"`
}

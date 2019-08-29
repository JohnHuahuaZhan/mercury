package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/db"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
)

//Insert UserInfo都存在默认值。只要不违反唯一约束，就成插入
//需要username,  password, email, user_id, sex, nickname
func InsertUser(user *model.UserInfo) error {
	//-------------开始插入-----------------------------
	sqlstr := "insert into user(username,  password, email, user_id, sex, nickname)values(?,?,?,?,?,?)"
	_, err := db.GetDB().Exec(sqlstr, user.Username, user.Password, user.Email, user.UserId, user.Sex, user.Nickname)
	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//UserByName 不存在则返回nil
func UserByName(username string) (*model.UserInfo, error) {
	sqlstr := "select username,password, user_id from user where username=?"
	info := new(model.UserInfo)
	err := db.GetDB().Get(info, sqlstr, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //没有错误，也没数据
		} else {
			return nil, orm.ErrSql
		}
	} else {
		return info, nil
	}
}

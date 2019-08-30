package dao

import (
	"database/sql"
	"github.com/JohnHuahuaZhan/mercury/orm"
	"github.com/JohnHuahuaZhan/mercury/orm/model"
	"github.com/jmoiron/sqlx"
)

type UserDao struct {
	*Dao
}

func NewUserDao(db *sqlx.DB, tx *sqlx.Tx) *UserDao {
	return &UserDao{NewDao(db, tx)}
}

//Insert UserInfo都存在默认值。只要不违反唯一约束，就成插入
//需要username,  password, email, user_id, sex, nickname
func (usrDao *UserDao) InsertUser(user *model.UserInfo) error {
	//-------------开始插入-----------------------------
	sqlstr := "insert into user(username,  password, email, user_id, sex, nickname)values(?,?,?,?,?,?)"
	var err error
	if nil != usrDao.tx {
		_, err = usrDao.tx.Exec(sqlstr, user.Username, user.Password, user.Email, user.UserId, user.Sex, user.Nickname)
	} else {
		if nil == usrDao.db {
			return orm.ErrDBParameterNull
		} else {
			_, err = usrDao.db.Exec(sqlstr, user.Username, user.Password, user.Email, user.UserId, user.Sex, user.Nickname)
		}
	}

	if nil != err {
		return orm.ErrSql
	}
	return nil
}

//UserByName 不存在则返回nil
func (usrDao *UserDao) UserByName(username string) (*model.UserInfo, error) {
	sqlstr := "select username,password, user_id from user where username=?"
	info := new(model.UserInfo)
	var err error

	if nil != usrDao.tx {
		err = usrDao.tx.Get(info, sqlstr, username)
	} else {
		if nil == usrDao.db {
			return nil, orm.ErrDBParameterNull
		} else {
			err = usrDao.db.Get(info, sqlstr, username)
		}
	}
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

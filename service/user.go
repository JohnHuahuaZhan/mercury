package service

import (
	"github.com/pingguoxueyuan/gostudy/mercury/orm/dao"
	model_orm "github.com/pingguoxueyuan/gostudy/mercury/orm/model"
	"github.com/pingguoxueyuan/gostudy/mercury/service/model"
	"github.com/pingguoxueyuan/gostudy/mercury/unique"
	"github.com/pingguoxueyuan/gostudy/mercury/util"
)

const (
	PasswordSalt = "HBZciU2SiSDr4uPeJ1e7qlIlMbyusQ0v" //密码加盐使用
)

// ServiceCreateUser 注册
func Register(user *model.UserInfoSecret) (err error) {

	userModel := new(model_orm.UserInfo)
	info, err := dao.UserByName(user.Username)
	if nil != err {
		return ErrUnKnow
	}
	if nil != info {
		return ErrUserExist
	}

	userModel.Username = user.Username

	//加盐
	passwd := user.Password + PasswordSalt //密码加盐
	md5pwd := util.Md5([]byte(passwd))     //做md5
	userModel.Password = md5pwd

	//生成全局唯一USER ID
	userId, err := unique.GetId()
	userModel.UserId = userId

	if err != nil {
		return ErrUnKnow
	}

	userModel.Email = user.Email
	userModel.Sex = user.Sex
	userModel.Nickname = user.Nickname

	err = dao.Insert(userModel)
	if nil != err {
		return ErrUnKnow
	}
	return nil
}

// ServiceCreateUser 注册
//请使用者注意，会修改你传入的的UserInfo的userID。只需要传入username，password
func Login(user *model.UserInfoSecret) (*model.UserInfoSecret, error) {
	info, err := dao.UserByName(user.Username)
	if nil != err {
		return nil, ErrUnKnow
	}
	if nil == info {
		return nil, ErrUserPasswordWrong
	}

	//加盐
	passwd := user.Password + PasswordSalt //密码加盐
	md5pwd := util.Md5([]byte(passwd))     //做md5

	if info.Password != md5pwd {
		return nil, ErrUserPasswordWrong
	}
	user.UserId = info.UserId
	return user, nil
}

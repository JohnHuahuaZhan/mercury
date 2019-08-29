package service

import "errors"

var (
	ErrParameter         = errors.New("参数错误")
	ErrUserExist         = errors.New("用户已存在")
	ErrUserPasswordWrong = errors.New("用户名或密码错误")
	ErrCategoryNotExist  = errors.New("分类不存在")
	ErrUnKnow            = errors.New("未知错误")
)

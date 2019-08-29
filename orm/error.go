package orm

import "errors"

var (
	ErrParameter = errors.New("参数错误")
	ErrSql       = errors.New("sql错误")
	ErrUnKnow    = errors.New("未知错误")
)

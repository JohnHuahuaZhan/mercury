package orm

import "errors"

var (
	ErrParameter       = errors.New("参数错误")
	ErrDBParameterNull = errors.New("db 和 tx同时为空")
	ErrSql             = errors.New("sql错误")
	ErrUnKnow          = errors.New("未知错误")
)

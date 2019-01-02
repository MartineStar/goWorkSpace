package model

import (
	"errors"
)
//根据业务逻辑,自定义错误类型
var (
	ERROR_USER_NOTEXISTS	= errors.New("用户不存在...")
	ERROR_USER_EXIST 		= errors.New("该用户已存在...")
	ERROR_USER_PWD			= errors.New("密码不正确")
)
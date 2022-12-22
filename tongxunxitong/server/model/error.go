package model

import "errors"

// 自定义错误
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在，请注册用户")
	ERROR_USER_EXITS     = errors.New("用户id已被注册，请重新注册")
	ERROR_USER_PWD       = errors.New("您的密码输入有误，请重新输入")
)

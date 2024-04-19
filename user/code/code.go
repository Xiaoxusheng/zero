package code

import "zero/pkg/code"

var (
	SuccessCode        = "success"
	UserExistError     = code.New(10001, "该手机号已经被注册！")
	UserNameExistError = code.New(10002, "用户名已经存在！")
	UserPwdError       = code.New(10003, "密码错误！")
	UserNotFoundError  = code.New(10002, "用户不存在！")
)

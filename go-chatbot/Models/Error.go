package Models

import "errors"

var (
	ErrorUserExit        = errors.New("用户名已存在")
	ErrorEmailExit       = errors.New("邮箱已存在")
	ErrorWrongPassword   = errors.New("密码错误")
	ErrorUserNotExit     = errors.New("用户名不存在")
	ErrorUserNotLogin    = errors.New("用户未登录")
	ErrorAuthHeaderValid = errors.New("无效Token")
	ErrorInvalidParm     = errors.New("参数错误")
)

func ErrorIs(err1 error, err2 error) bool {
	return err1.Error() == err2.Error()
}

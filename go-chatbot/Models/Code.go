package Models

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParm
	CodeUserExist
	CodeUserNotExist
	CodeWrongPassword
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthHeaderIsNil
	CodeAuthHeaderValid
	CodeSendVerifiErr
	CodeVerifiNotFund
	CodeVerifiErr
)

var CodeMesMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParm:     "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeWrongPassword:   "密码错误",
	CodeInvalidPassword: "密码格式错误",
	CodeServerBusy:      "服务正忙",
	CodeAuthHeaderIsNil: "需要登录后进行操作",
	CodeAuthHeaderValid: "无效的Token",
	CodeSendVerifiErr: "发送验证码失败",
	CodeVerifiNotFund: "验证码过期或未发送",
	CodeVerifiErr: "验证码错误",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMesMap[c]
	if !ok {
		msg = CodeMesMap[CodeServerBusy]
	}
	return msg
}

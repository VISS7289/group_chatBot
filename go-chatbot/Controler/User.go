package Controler

import (
	"go-chatbot/Logic"
	"go-chatbot/Models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	//参数校验
	var p Models.ParmRegister
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Register Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	if len(p.Username) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "用户名不能为空")
		return
	}
	if len(p.Password) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码不能为空")
		return
	}
	if len(p.Username) < 4 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "用户名过短")
		return
	}
	if len(p.Password) < 8 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码过短")
		return
	}
	if len(p.RePassword) < 8 || p.RePassword != p.Password {
		Models.ResponseErrorWithMsg(c, Models.CodeWrongPassword, "请重新确认密码")
		return
	}
	//业务处理
	atoken, rtoken, err := Logic.Register(&p)
	if err != nil {
		if Models.ErrorIs(err, Models.ErrorUserExit) {
			Models.ResponseError(c, Models.CodeUserExist)
			return
		}
		zap.L().Error("SomeOne Have Unknow Error When Regist:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, []string{atoken, rtoken})
}

func LoginHandler(c *gin.Context) {
	//参数校验
	var p Models.ParmLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	if len(p.Username) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "用户名不能为空")
		return
	}
	if len(p.Password) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码不能为空")
		return
	}
	//业务处理
	atoken, rtoken, err := Logic.Login(&p)
	if err != nil {
		if Models.ErrorIs(err, Models.ErrorWrongPassword) || Models.ErrorIs(err, Models.ErrorUserNotExit) {
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err.Error())
			return
		}
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, []string{atoken, rtoken})
}

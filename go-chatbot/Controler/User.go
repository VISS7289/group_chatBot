package Controler

import (
	"fmt"
	"go-chatbot/Logic"
	"go-chatbot/Models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 处理用户注册
// @Summary 处理用户注册接口
// @Description 处理用户注册接口
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmRegister false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /register [post]
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
	//验证验证码
	if code := Logic.VerifiExam(&p); code != Models.CodeSuccess {
		Models.ResponseErrorWithMsg(c, code, "验证码错误")
		return
	}
	//业务处理
	info, err, code := Logic.Register(&p)
	if err != nil {
		if Models.ErrorIs(err, Models.ErrorUserExit) {
			Models.ResponseError(c, Models.CodeUserExist)
			return
		}
		zap.L().Error("SomeOne Have Unknow Error When Regist:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	if code != Models.CodeSuccess {
		Models.ResponseErrorWithMsg(c, code, "验证码错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)
}

func ExistName(c *gin.Context) {
	name := c.Query("username")
	fmt.Println(name)
	if err, code := Logic.RepeatUsername(name); err != nil {
		Models.ResponseErrorWithMsg(c, code, "用户名重复")
		return
	}
	Models.ResponseSuccess(c, []string{})
}

// LoginHandler 处理用户登录
// @Summary 处理用户登录接口
// @Description 处理用户登录接口
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmLogin false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /login [post]
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
	info, err := Logic.Login(&p)
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

	Models.ResponseSuccess(c, info)
}

package Controler

import (
	"fmt"
	"go-chatbot/Logic"
	"go-chatbot/Models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 检验邮箱是否重复
// @Summary 检验邮箱是否重复
// @Description 需要用户名、密码、确认密码、邮箱与验证码
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmRegister false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /existEmail [post]
func RegisterHandler(c *gin.Context) {
	// 参数校验，将请求体中的json数据绑定到ParmRegister结构体中
	var p Models.ParmRegister
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Register Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 对用户名和密码进行长度和格式校验
	if len(p.Username) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "用户名不能为空")
		return
	}
	if len(p.Password) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码不能为空")
		return
	}
	if len(p.Username) < 4 {
		println(p.Username)
		println(len(p.Username))
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
		// 如果验证码错误，则返回相应的错误码和错误信息
		Models.ResponseErrorWithMsg(c, code, "验证码错误")
		return
	}
	//业务处理，处理用户注册请求
	info, err, code := Logic.Register(&p)
	if err != nil {
		// 如果发生用户已存在错误，则返回CodeUserExist错误
		if Models.ErrorIs(err, Models.ErrorUserExit) {
			Models.ResponseError(c, Models.CodeUserExist)
			return
		}
		// 否则，记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Regist:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	if code != Models.CodeSuccess {
		// 如果验证码错误，则返回相应的错误码和错误信息
		Models.ResponseErrorWithMsg(c, code, "验证码错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)
}

// ExistName 检验名称是否重复
// @Summary 检验名称是否重复
// @Description 需要发送者用户名
// @Tags 注册登录相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmGetNewMsgOne false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /existName [get]
func ExistName(c *gin.Context) {
	// 获取查询参数中的用户名
	name := c.Query("username")
	fmt.Println(name)
	// 检查用户名是否重复
	if err, code := Logic.RepeatUsername(name); err != nil {
		// 如果用户名重复，则返回相应的错误码和错误信息
		Models.ResponseErrorWithMsg(c, code, "用户名重复")
		return
	}
	// 返回成功响应
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
	// 参数校验，将请求体中的json数据绑定到ParmLogin结构体中
	var p Models.ParmLogin
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 对用户名和密码进行非空校验
	if len(p.Username) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "用户名不能为空")
		return
	}
	if len(p.Password) == 0 {
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "密码不能为空")
		return
	}
	//业务处理，处理用户登录请求
	info, err := Logic.Login(&p)
	if err != nil {
		// 如果发生用户名或密码错误，则返回相应的错误码和错误信息
		if Models.ErrorIs(err, Models.ErrorWrongPassword) || Models.ErrorIs(err, Models.ErrorUserNotExit) {
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err.Error())
			return
		}
		// 否则，记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)
}

// GetUserDetail 获取用户详情路由
// @Summary 获取用户详情路由
// @Description 需要token验证，通过用户ID查找用户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmQueryUserDetial false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /user/detial [post]
func GetUserDetail(c *gin.Context) {
	// 参数校验，将请求体中的json数据绑定到ParmQueryUserDetial结构体中
	var p Models.ParmQueryUserDetial
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理，查询用户详细信息
	info, err := Logic.QueryDetialByID(&p)
	if err != nil {
		// 如果查询发生错误，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)

}

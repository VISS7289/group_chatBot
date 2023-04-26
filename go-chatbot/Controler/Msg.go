package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

// GetNewMsgOne 获取最新一条消息
// @Summary 获取最新一条消息
// @Description 需要发送者id与接收者id
// @Tags 消息相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmGetNewMsgOne false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /msg/newone [post]
func GetNewMsgOne(c *gin.Context) {
	//参数校验
	var p Models.ParmGetNewMsgOne
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	info, err := Logic.GetNewMsgOne(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)
}

// GetUnReadMsgNum 获取未读消息数
// @Summary 获取未读消息数
// @Description 需要发送者id与接收者id
// @Tags 消息相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmGetNewMsgOne false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /msg/unread [post]
func GetUnReadMsgNum(c *gin.Context) {
	//参数校验
	var p Models.ParmGetNewMsgOne
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	num, err := Logic.GetUnReadMsgNum(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, num)
}

// GetOldChatF 获取用户聊天消息
// @Summary 获取用户聊天消息
// @Description 需要发送者id与接收者id以及当前page位置与一页最多获取多少消息
// @Tags 消息相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmGetOldChat false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /msg/getOldChatF [post]
func GetOldChatF(c *gin.Context) {
	//参数校验
	var p Models.ParmGetOldChat
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	info, err := Logic.GetOldMsgF(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, info)
}

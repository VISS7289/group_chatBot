package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

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

package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

// SerchUsername 根据用户名模糊搜索
// @Summary 根据用户名模糊搜索
// @Description 需要token验证，通过用户名称进行模糊搜索
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchUsername false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /serch/username [post]
func SerchUsername(c *gin.Context) {
	//参数校验，将请求体中的json数据绑定到ParmSerchUsername结构体中
	var p Models.ParmSerchUsername
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理，搜索用户名
	userRes, err := Logic.SerchUser(&p)
	if err != nil {
		// 如果发生错误，则根据错误类型返回相应的错误码和错误信息
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
	Models.ResponseSuccess(c, userRes)
}

// SerchIsFriend 判断两个id是否为好友
// @Summary 判断两个id是否为好友
// @Description 需要token验证，根据两用户的id，判断两用户是否为好友
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchIsFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /serch/isfriend [post]
func SerchIsFriend(c *gin.Context) {
	//参数校验
	var p Models.ParmSerchIsFriend
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	// state 0为好友 1为申请中 2为非好友
	friendRes, err := Logic.SerchIsFriend(&p)
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
	Models.ResponseSuccess(c, friendRes)
}

package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

func FriendRequest(c *gin.Context) {
	//参数校验
	var p Models.ParmFriendRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	q := Models.ParmSerchIsFriend{
		UserId:   p.UserId,
		FriendId: p.FriendId,
	}
	friendRes, err := Logic.SerchIsFriend(&q)
	if err != nil {
		if Models.ErrorIs(err, Models.ErrorWrongPassword) || Models.ErrorIs(err, Models.ErrorUserNotExit) {
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err.Error())
			return
		}
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	if friendRes.State != 2 && friendRes.State != 1{
		zap.L().Error("SomeOne Illegality Send Friend Request:", zap.Error(err))
		Models.ResponseError(c, Models.CodeIllegalityFriendReq)
		return
	}
	//业务处理
	//state 0为好友 1为申请中 2为非好友
	if friendRes.State ==2{
		if err := Logic.FriendRequest(&p); err != nil {
			zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//返回响应
		Models.ResponseSuccess(c, "success")
	}

	if friendRes.State ==1{
		if err := Logic.FriendRequestN(&p); err != nil {
			zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//返回响应
		Models.ResponseSuccess(c, "success")
	}

}

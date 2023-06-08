package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

// FriendRequest 好友申请
// @Summary 提交好友请求
// @Description 需要token验证，发起者id为用户id，请求者id为好友id，消息内容为msg
// @Tags 好友相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchIsFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /friend/request [post]
func FriendRequest(c *gin.Context) {
	//参数校验，将请求体中的json数据绑定到ParmFriendRequest结构体中
	var p Models.ParmFriendRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//检查好友关系
	q := Models.ParmSerchIsFriend{
		UserId:   p.UserId,
		FriendId: p.FriendId,
	}
	friendRes, err := Logic.SerchIsFriend(&q)
	if err != nil {
		// 如果返回的错误是密码错误或用户不存在，则返回CodeInvalidParm错误，并返回错误信息
		if Models.ErrorIs(err, Models.ErrorWrongPassword) || Models.ErrorIs(err, Models.ErrorUserNotExit) {
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err.Error())
			return
		}
		// 否则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	// 如果好友关系不是2或1，则记录错误日志，并返回CodeIllegalityFriendReq错误
	if friendRes.State != 2 && friendRes.State != 1 {
		zap.L().Error("SomeOne Illegality Send Friend Request:", zap.Error(err))
		Models.ResponseError(c, Models.CodeIllegalityFriendReq)
		return
	}
	//业务处理
	//state 0为好友 1为申请中 2为非好友
	if friendRes.State == 2 {
		//如果好友状态是2，则发送好友请求
		if err := Logic.FriendRequest(&p); err != nil {
			// 发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
			zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//发送好友请求成功，则返回响应
		Models.ResponseSuccess(c, "success")
	}

	if friendRes.State == 1 {
		//如果好友状态是1，则更新好友请求
		if err := Logic.FriendRequestN(&p); err != nil {
			// 发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
			zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//更新好友请求成功，则返回响应
		Models.ResponseSuccess(c, "success")
	}

}

// FriendDeleate 好友删除
// @Summary 删除某个好友
// @Description 需要token验证，发起者id为用户id，被删除者id为好友id
// @Tags 好友相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchIsFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /friend/deleate [post]
func FriendDeleate(c *gin.Context) {
	//参数校验，将请求体中的json数据绑定到ParmSerchIsFriend结构体中
	var p Models.ParmSerchIsFriend
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理，删除好友
	err := Logic.DelFriend(&p)
	if err != nil {
		// 如果发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}

	//返回响应
	Models.ResponseSuccess(c, "success")
}

// FriendReject 好友请求拒绝
// @Summary 拒绝某人发起的好友请求
// @Description 需要token验证，发起者id为用户id，被拒绝者id为好友id
// @Tags 好友相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchIsFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /friend/reject [post]
func FriendReject(c *gin.Context)  {
	//参数校验，将请求体中的json数据绑定到ParmSerchIsFriend结构体中
	var p Models.ParmSerchIsFriend
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理，拒绝好友请求
	err := Logic.RejectFriend(&p)
	if err != nil {
		// 如果发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}

	//返回响应
	Models.ResponseSuccess(c, "success")
}

// FriendAccept 好友接受
// @Summary 接受好友请求
// @Description 需要token验证，发起者id为用户id，接受者id为好友id
// @Tags 好友相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmSerchIsFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /friend/accept [post]
func FriendAccept(c *gin.Context) {
	//参数校验，将请求体中的json数据绑定到ParmSerchIsFriend结构体中
	var p Models.ParmSerchIsFriend
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理，接受好友请求
	err := Logic.AcceptFriend(&p)
	if err != nil {
		// 如果发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("SomeOne Have Unknow Error When Delete Friend:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")
}

// GetMyFriend 获取好友
// @Summary 获取到好友列表或好友请求
// @Description 需要token验证，需要提供用户id与结果方式，0 代表好友列表，1代表好友申请列表
// @Tags 好友相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmGetFriend false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /friend/myFriend [post]
func GetMyFriend(c *gin.Context) {
	//参数校验，将请求体中的json数据绑定到ParmGetFriend结构体中
	var p Models.ParmGetFriend
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Login Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	if p.State ==0{
		//获取我的好友列表
		FriendRes, err := Logic.GetMyFriend(&p)
		if err != nil {
			// 如果发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
			zap.L().Error("SomeOne Have Unknow Error When Get His Friend:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//返回响应
		Models.ResponseSuccess(c, FriendRes)
		return
	}else if p.State ==1{
		//获取我的好友请求列表
		FriendRes, err := Logic.GetMyFriendReq(&p)
		if err != nil {
			// 如果发生未知错误，则记录错误日志，并返回CodeInvalidParm错误
			zap.L().Error("SomeOne Have Unknow Error When Get His Friend:", zap.Error(err))
			Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
			return
		}
		//返回响应
		Models.ResponseSuccess(c, FriendRes)
		return
	}

	// 如果状态码不是0或1，则返回参数错误
	Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "参数错误")


}

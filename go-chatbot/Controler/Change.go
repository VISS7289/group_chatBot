package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Logic"
	"go-chatbot/Models"
	"go.uber.org/zap"
)

// ChangeUserImg 修改用户图片
// @Summary 修改用户头像
// @Description 修改用户头像，需要用户id与base64格式图片
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChangeImg false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /change/img [post]
func ChangeUserImg(c *gin.Context) {
	// 数据绑定，将请求体中的json数据绑定到ParmChangeImg结构体中
	var p Models.ParmChangeImg
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	err := Logic.ChangeUserImg(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

// ChangeUser 修改用户简介
// @Summary 修改用户简介
// @Description 修改用户头像，需要用户id、修改后数据、修改的数据的类型与密码，对于特定类型数据，需要密码确认才可以修改
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChange false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /change/update [post]
func ChangeUser(c *gin.Context) {
	// 数据绑定
	var p Models.ParmChange
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	//业务处理
	err := Logic.ChangeUser(&p)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

// ChangeNick 修改好友备注
// @Summary 修改好友备注即昵称
// @Description 需要用户id，好友id与修改后昵称
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmFriendRequest false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /change/nick [post]
func ChangeNick(c *gin.Context) {
	// 数据绑定
	var p Models.ParmFriendRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 对绑定的数据进行节选
	q := Models.ParmSerchIsFriend{
		UserId:   p.UserId,
		FriendId: p.FriendId,
	}
	// 业务逻辑
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
	// 非法发送好友请求
	if friendRes.State != 0 {
		zap.L().Error("SomeOne Illegality Send Friend Request:", zap.Error(err))
		Models.ResponseError(c, Models.CodeIllegalityFriendReq)
		return
	}
	//业务处理
	//state 0为好友 1为申请中 2为非好友

	if err := Logic.ChangeNick(&p); err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Login:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, "未知错误")
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

// ChangeEmail 修改邮箱
// @Summary 修改邮箱
// @Description 需要用户id，新的邮箱与验证码
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChangeEmail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /change/email [post]
func ChangeEmail(c *gin.Context) {
	// 数据绑定
	var p Models.ParmChangeEmail
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 节选部分数据
	q := Models.VerifiExam{
		VerifiCode: p.VerifiCode,
		Email:      p.NewEmail,
	}
	// Redis处理验证码
	val, err := Redis.RedisFind(q.Email)
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val != q.VerifiCode {
		Models.ResponseError(c, Models.CodeVerifiErr)
		return
	}
	//业务逻辑
	s := Models.ParmChange{
		UserId: p.UserId,
		Data:   p.NewEmail,
		Type:   "email",
		Psw:    "",
	}
	err = Logic.ChangeUser(&s)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

// ChangePsw 修改密码
// @Summary 修改密码
// @Description 需要用户id，邮箱，新的密码，确认密码以及验证码
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChangePsw false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /change/psw [post]
func ChangePsw(c *gin.Context) {
	// 数据绑定
	var p Models.ParmChangePsw
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Get User Detail Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 节选部分数据
	q := Models.VerifiExam{
		VerifiCode: p.VerifiCode,
		Email:      p.Email,
	}
	// Redis验证验证码
	val, err := Redis.RedisFind(q.Email)
	// 错误情况处理
	if err != nil {
		Models.ResponseError(c, Models.CodeVerifiNotFund)
		return
	}
	if val != q.VerifiCode {
		Models.ResponseError(c, Models.CodeVerifiErr)
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
	s := Models.ParmChange{
		UserId: p.UserId,
		Data:   p.Password,
		Type:   "password",
		Psw:    "",
	}
	err = Logic.ChangeUser(&s)
	if err != nil {
		zap.L().Error("SomeOne Have Unknow Error When Change Img:", zap.Error(err))
		Models.ResponseErrorWithMsg(c, Models.CodeInvalidParm, err)
		return
	}
	//返回响应
	Models.ResponseSuccess(c, "success")

}

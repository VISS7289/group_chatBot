package Routs

import (
	"fmt"
	"go-chatbot/Controler"
	"go-chatbot/Logger"
	"go-chatbot/Middlewares"
	"go-chatbot/Models"
	"go-chatbot/Settings"

	_ "go-chatbot/docs"

	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(Middlewares.Cors())
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	//接口文档路由
	r.GET("/swagger/*any", gs.WrapHandler(swaggerfiles.Handler))
	//debug路由
	r.POST("/ping", Middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		Models.ResponseSuccess(c, "helloWord")
	})
	//刷新token路由
	r.POST("/refersh", Controler.RefershHandler)
	//获取用户详情路由
	r.POST("/user/detial", Middlewares.JWTAuthMiddleware(), Controler.GetUserDetail)

	//根据用户名模糊搜索
	r.POST("/serch/username", Middlewares.JWTAuthMiddleware(), Controler.SerchUsername)
	//判断两个id是否为好友
	r.POST("/serch/isfriend", Middlewares.JWTAuthMiddleware(), Controler.SerchIsFriend)

	//好友申请
	r.POST("/friend/request", Middlewares.JWTAuthMiddleware(), Controler.FriendRequest)
	//好友接受
	r.POST("/friend/accept", Middlewares.JWTAuthMiddleware(), Controler.FriendAccept)
	//好友删除
	r.POST("/friend/deleate", Middlewares.JWTAuthMiddleware(), Controler.FriendDeleate)
	//好友请求拒绝
	r.POST("/friend/reject", Middlewares.JWTAuthMiddleware(), Controler.FriendReject)
	//获取好友
	r.POST("/friend/myFriend", Middlewares.JWTAuthMiddleware(), Controler.GetMyFriend)

	//修改用户图片
	r.POST("/change/img", Middlewares.JWTAuthMiddleware(), Controler.ChangeUserImg)
	//修改用户简介
	r.POST("/change/update", Middlewares.JWTAuthMiddleware(), Controler.ChangeUser)
	//修改好友备注
	r.POST("/change/nick", Middlewares.JWTAuthMiddleware(), Controler.ChangeNick)
	//修改邮箱
	r.POST("/change/email", Middlewares.JWTAuthMiddleware(), Controler.ChangeEmail)
	//修改密码
	r.POST("/change/psw", Middlewares.JWTAuthMiddleware(), Controler.ChangePsw)

	//获取最新一条消息
	r.POST("/msg/newone", Middlewares.JWTAuthMiddleware(), Controler.GetNewMsgOne)
	//获取未读消息数
	r.POST("/msg/unread", Middlewares.JWTAuthMiddleware(), Controler.GetUnReadMsgNum)
	//获取用户聊天消息
	r.POST("/msg/getOldChatF", Middlewares.JWTAuthMiddleware(), Controler.GetOldChatF)


	//注册功能路由
	r.POST("/register", Controler.RegisterHandler)
	r.GET("/existName", Controler.ExistName)
	r.POST("/existEmail", Controler.RegisterHandler)
	//登录功能路由
	r.POST("/login", Controler.LoginHandler)
	//聊天功能路由
	r.POST("/chat", Controler.ChatHandler)
	r.POST("/chat_rand", Controler.RandChatHandler)
	//邮箱验证路由
	r.POST("/verificationCode", Controler.VerifiHandler)
	//检查验证码是否正确
	r.POST("/verifiExam", Controler.VerifiExam)
	_ = r.Run(fmt.Sprintf(":%d", Settings.Conf.Port))

	return r
}

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

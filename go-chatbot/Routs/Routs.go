package Routs

import (
	"fmt"
	"go-chatbot/Controler"
	"go-chatbot/Logger"
	"go-chatbot/Middlewares"
	"go-chatbot/Models"
	"go-chatbot/Settings"

	"github.com/gin-gonic/gin"
)

func Init(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	//debug路由
	r.GET("/ping", Middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		Models.ResponseSuccess(c, "helloWord")
	})
	//注册功能路由
	r.POST("/register", Controler.RegisterHandler)
	//登录功能路由
	r.POST("/login", Controler.LoginHandler)
	//聊天功能路由
	r.POST("/chat", Controler.ChatHandler)
	r.POST("/chat_rand", Controler.RandChatHandler)
	_ = r.Run(fmt.Sprintf(":%d", Settings.Conf.Port))
	return r
}

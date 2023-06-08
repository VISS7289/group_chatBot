package Controler

import (
	"go-chatbot/Middlewares"
	"go-chatbot/Models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (int64, error) {
	// 从gin上下文中获取用户ID
	uid, ok := c.Get(Middlewares.ContexUserID)
	// 如果获取不到用户ID，则返回错误UserNotLogin
	if !ok {
		return 0, Models.ErrorUserNotLogin
	}
	// 将uid转换为int64类型的用户ID
	userid, ok := uid.(int64)
	// 如果类型转换失败，则返回错误UserNotLogin
	if !ok {
		return 0, Models.ErrorUserNotLogin
	}
	// 返回用户ID和nil错误
	return userid, nil
}

package Controler

import (
	"go-chatbot/Middlewares"
	"go-chatbot/Models"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (int64, error) {
	uid, ok := c.Get(Middlewares.ContexUserID)
	if !ok {
		return 0, Models.ErrorUserNotLogin
	}
	userid, ok := uid.(int64)
	if !ok {
		return 0, Models.ErrorUserNotLogin
	}
	return userid, nil
}

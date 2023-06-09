package Middlewares

import (
	"go-chatbot/Models"
	"go-chatbot/Pkg/JWT"
	"strings"

	"github.com/gin-gonic/gin"
)

const ContexUserID = "userID"

//此函数返回一个JWT鉴权的中间件函数
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//从请求头中提取Authorization字段
		authHeader := c.Request.Header.Get("Authorization")
		//验证格式
		if authHeader == "" {
			Models.ResponseError(c, Models.CodeAuthHeaderIsNil)
			c.Abort()
			return
		}
		//解析Authorization
		parts := strings.SplitN(authHeader, " ", 2)
		//格式不匹配
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			Models.ResponseError(c, Models.CodeAuthHeaderValid)
			c.Abort()
			return
		}
		//解析JWT token
		mc, err := JWT.ParseToken(parts[1])
		//v, _ := err.(*jwt.ValidationError)
		//if v.Errors == jwt.ValidationErrorExpired {
		//	Models.ResponseError(c, Models.CodeAuthHeaderIsNil)
		//	c.Abort()
		//	return
		//}
		//验证失败
		if err != nil {
			Models.ResponseError(c, Models.CodeAuthHeaderValid)
			c.Abort()
			return
		}
		//将用户ID设置到请求上下文中
		c.Set(ContexUserID, mc.UserID)
		//放行请求
		c.Next()
	}
}

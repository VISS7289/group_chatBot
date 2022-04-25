package Middlewares

import (
	"go-chatbot/Models"
	"go-chatbot/Pkg/JWT"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

const ContexUserID = "userID"

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			Models.ResponseError(c, Models.CodeAuthHeaderIsNil)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			Models.ResponseError(c, Models.CodeAuthHeaderValid)
			c.Abort()
			return
		}
		mc, err := JWT.ParseToken(parts[1])
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			Models.ResponseError(c, Models.CodeAuthHeaderIsNil)
			c.Abort()
			return
		}
		if err != nil {
			Models.ResponseError(c, Models.CodeAuthHeaderValid)
			c.Abort()
			return
		}
		c.Set(ContexUserID, mc.UserID)
		c.Next()
	}
}

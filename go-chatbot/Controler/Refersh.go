package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Models"
	"go-chatbot/Pkg/JWT"
	"go.uber.org/zap"
)

func RefershHandler(c *gin.Context) {
	var p Models.ParmToken
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Token Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	newAToken, newRToken, err := JWT.RefreshToken(p.AToken,p.RToken)
	if err != nil {
		Models.ResponseError(c, Models.CodeAuthHeaderValid)
		return
	}
	Models.ResponseSuccess(c, []string{newAToken, newRToken})
}

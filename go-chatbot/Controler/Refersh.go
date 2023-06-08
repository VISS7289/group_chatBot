package Controler

import (
	"github.com/gin-gonic/gin"
	"go-chatbot/Models"
	"go-chatbot/Pkg/JWT"
	"go.uber.org/zap"
)

// RefershHandler 刷新用户token
// @Summary 刷新用户token接口
// @Description 通过access token与refersh token进行token刷新
// @Tags token相关接口
// @Accept application/json
// @Produce application/json
// @Param object query Models.ParmToken false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /refersh [post]
func RefershHandler(c *gin.Context) {
	var p Models.ParmToken
	// 刷新令牌
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Token Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 返回新的令牌
	newAToken, newRToken, err := JWT.RefreshToken(p.AToken,p.RToken)
	if err != nil {
		Models.ResponseError(c, Models.CodeAuthHeaderValid)
		return
	}
	Models.ResponseSuccess(c, []string{newAToken, newRToken})
}

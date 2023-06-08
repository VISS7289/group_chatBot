package Controler

import (
	"bytes"
	"go-chatbot/Models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ChatHandler 聊天功能路由
// @Summary 聊天功能路由
// @Description 需要发送的聊天消息
// @Tags 聊天相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChat false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /chat [post]
func ChatHandler(c *gin.Context) {
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 发送 POST 请求到 Python 服务器
	url := "http://localhost:8087/chat"
	post := "{\"chat_content\":\"" + p.ChatContent +
		"\"}"
	var jsonStr = []byte(post)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Python Server Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}

	// 读取响应体并返回响应
	body, _ := ioutil.ReadAll(resp.Body)
	Models.ResponseSuccess(c, string(body))
	return
}

// RandChatHandler 随机聊天
// @Summary 随机聊天
// @Description 需要发送的聊天消息
// @Tags 聊天相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChat false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /chat_rand [post]
func RandChatHandler(c *gin.Context) {
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 发送 POST 请求到 Python 服务器
	url := "http://localhost:8087/chat_rand"
	post := "{\"chat_content\":\"" + p.ChatContent +
		"\"}"
	var jsonStr = []byte(post)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Python Server Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}
	// 读取响应体并返回响应
	body, _ := ioutil.ReadAll(resp.Body)
	Models.ResponseSuccess(c, string(body))
	return
}

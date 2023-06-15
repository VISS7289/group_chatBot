package Controler

import (
	"bytes"
	"go-chatbot/Models"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var History [][]string

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
	// 数据绑定，将请求体中的json数据绑定到ParmChat结构体中
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 发送 POST 请求到 Python 服务器
	// 构造POST请求的URL和请求体
	//url := "http://localhost:8087/chat"
	//post := "{\"chat_content\":\"" + p.ChatContent +
	//	"\"}"

	url := "http://localhost:7002"
	post := "{\"prompt\":" + p.ChatContent + ",\"history\": []}"
	var jsonStr = []byte(post)
	// 构造POST请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	// 发送POST请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		// 如果发送POST请求出现错误，则记录错误日志，并返回CodeServerBusy错误
		zap.L().Error("Python Server Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}

	// 读取响应体并返回响应
	body, _ := ioutil.ReadAll(resp.Body)
	var ParmServer Models.ParmChatServer
	if err := json.Unmarshal(body, &ParmServer); err != nil {
		zap.L().Error("Python Server Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}
	Models.ResponseSuccess(c, ParmServer.Response)
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
	// 数据绑定，将请求体中的json数据绑定到ParmChat结构体中
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		// 如果绑定数据失败，则记录错误日志，并返回CodeInvalidParm错误
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 发送 POST 请求到 Python 服务器
	// 构造POST请求的URL和请求体
	url := "http://localhost:8087/chat_rand"
	post := "{\"chat_content\":\"" + p.ChatContent +
		"\"}"
	var jsonStr = []byte(post)
	// 构造POST请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	// 发送POST请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		// 如果发送POST请求出现错误，则记录错误日志，并返回CodeServerBusy错误
		zap.L().Error("Python Server Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeServerBusy)
		return
	}
	// 读取响应体并返回响应
	body, _ := ioutil.ReadAll(resp.Body)
	Models.ResponseSuccess(c, string(body))
	return
}

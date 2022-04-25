package Controler

import (
	"bytes"
	"go-chatbot/Models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ChatHandler(c *gin.Context) {
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
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
	body, _ := ioutil.ReadAll(resp.Body)
	Models.ResponseSuccess(c, string(body))
	return
}

func RandChatHandler(c *gin.Context) {
	var p Models.ParmChat
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Request Parm Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
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
	body, _ := ioutil.ReadAll(resp.Body)
	Models.ResponseSuccess(c, string(body))
	return
}

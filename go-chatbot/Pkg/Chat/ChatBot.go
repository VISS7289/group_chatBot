package Chat

import (
	"bytes"
	"go-chatbot/Models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

// 该函数用于用POST请求向Python服务发送聊天内容
func GetChat(ChatContent string) (string, error) {
	// 定义请求url和数据
	//url := "http://localhost:8087/chat"
	//post := "{\"chat_content\":\"" + ChatContent +
	//	"\"}"

	url := "http://localhost:7002"
	post := "{\"prompt\":\"" + ChatContent + "。回答限制在50字以内。\",\"history\": [],\"max_length\":50}"
	// 构造JSON数据
	var jsonStr = []byte(post)
	// 构造POST请求
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// 设置Header
	req.Header.Set("Content-Type", "application/json")
	// 初始化Client
	client := &http.Client{}
	// 发送POST请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Python Server Error", zap.Error(err))
		return "",err
	}
	// 读取响应体
	body, _ := ioutil.ReadAll(resp.Body)
	var ParmServer Models.ParmChatServer
	if err := json.Unmarshal(body, &ParmServer); err!=nil{
		zap.L().Error("Python Server Error", zap.Error(err))
		return "error",err
	}
	return ParmServer.Response, nil
}

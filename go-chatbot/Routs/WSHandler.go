package Routs

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-chatbot/Models"
	"go.uber.org/zap"
	"net/http"
)

const destory = 60 * 60 * 24 * 30 //过期时间 30天


func CreateId(uid, toid string) string {
	return uid + "->" + toid
}

func Handler(c *gin.Context) {
	uid := c.Query("send_id")
	touid := c.Query("accept_id")
	// 升级为ws协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		zap.L().Error("Websocket Unknow Error", zap.Error(err))
		Models.ResponseError(c, Models.CodeInvalidParm)
		return
	}
	// 创建一个用户实例
	client := &Client{
		ID:CreateId(uid,touid),
		SendId:CreateId(touid,uid),
		Socket:conn,
		Send: make(chan []byte),
	}
	// 用户注册到用户管理
	Manager.Register <- client

	go client.Read()
	go client.Write()
}


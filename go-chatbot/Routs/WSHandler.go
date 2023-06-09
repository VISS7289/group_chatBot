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

// Handler websocket聊天
// @Summary 建立websocket长连接
// @Description 发送者id，接收者id
// @Tags 聊天相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query Models.ParmChat false "查询参数"
// @Security ApiKeyAuth
// @Router /ws [post]
func Handler(c *gin.Context) {
	// 获取发送者 ID
	uid := c.Query("send_id")
	// 获取接收者 ID
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
		ID:     CreateId(uid, touid), // 创建客户端 ID
		SendId: CreateId(touid, uid), // 创建发送者 ID
		Socket: conn,                 // WebSocket 连接对象
		Send:   make(chan *JsonMsg),  // 发送消息的通道
	}
	// 用户注册到用户管理
	Manager.Register <- client
	// 在新的 Go 协程中异步地执行客户端的读写操作
	go client.Read()  // 读取客户端发送的消息
	go client.Write() // 向客户端发送消息
}

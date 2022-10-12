package Routs

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Dao/Redis"
	"go-chatbot/Models"
	"go-chatbot/Pkg/Chat"
	"go.uber.org/zap"
	"math/rand"
	"strings"
	"time"
)

// 发送消息
type SendMsg struct {
	Type int    `json:"type"`
	Msg  string `json:"message"`
}

// 回复消息
type ReplyMsg struct {
	SendId string         `json:"send_id"`
	Code   Models.ResCode `json:"code"`
	Msg    string         `json:"message"`
}

type JsonMsg struct {
	SendId string `json:"send_id"`
	Msg    []byte `json:"message"`
}

// 用户结构体
type Client struct {
	ID     string
	SendId string
	Socket *websocket.Conn
	Send   chan *JsonMsg
}

// 广播类（包括广播内容和源用户）
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

// 用户管理
type CliendManager struct {
	Client     map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

// 信息转json
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Msg       string `json:"msg,omitempty"`
}

var Manager = CliendManager{
	Client:     make(map[string]*Client),
	Broadcast:  make(chan *Broadcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendmsg := new(SendMsg)
		if err := c.Socket.ReadJSON(&sendmsg); err != nil {
			zap.L().Error("Websocket Unknow Error", zap.Error(err))
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendmsg.Type == 1 {
			r1 := Redis.RedisVal(c.ID)
			r2 := Redis.RedisVal(c.SendId)
			if r1 > "3" && r2 == "" {
				replymsg := ReplyMsg{
					Code: Models.CodeExceedMsg,
					Msg:  Models.CodeExceedMsg.Msg(),
				}
				msg, _ := json.Marshal(replymsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				_ = Redis.RedisAdd(c.ID)
			}
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendmsg.Msg),
			}
		}
		// 需要机器人主动发消息
		if sendmsg.Type == 0 {
			msgs, _ := Chat.GetChat(sendmsg.Msg)
			splitId := strings.Split(c.ID, "->")
			replymsg := ReplyMsg{
				Code: Models.CodeSuccess,
				Msg:  string(msgs),
			}
			msg, _ := json.Marshal(replymsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			InsertMsg := Models.NewMsg{
				SendId:   splitId[1],
				AcceptId: splitId[0],
				State:    false,
				Msg:      string(msgs),
				Type:     0,
			}
			if err := Mysql.InsertMsg(InsertMsg); err != nil {
				zap.L().Error("Verifi Parm Error", zap.Error(err))
			}
		}

	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case mssage, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replymsg := ReplyMsg{
				SendId: mssage.SendId,
				Code:   Models.CodeSuccess,
				Msg:    fmt.Sprintf("%s", string(mssage.Msg)),
			}
			msg, _ := json.Marshal(replymsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			continue
		}
	}
}

func (manager *CliendManager) Start() {
	for {
		fmt.Println("-----------监听管道通信-------------")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新连接：%v\n", conn.ID)
			Manager.Client[conn.ID] = conn
			replymsg := ReplyMsg{
				Code: Models.CodeWSSuccess,
				Msg:  Models.CodeWSSuccess.Msg(),
			}
			msg, _ := json.Marshal(replymsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Printf("连接失败：%v\n", conn.ID)
			if _, ok := Manager.Client[conn.ID]; ok {
				replymsg := ReplyMsg{
					Code: Models.CodeWebSocketEnd,
					Msg:  Models.CodeWebSocketEnd.Msg(),
				}
				msg, _ := json.Marshal(replymsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Client, conn.ID)
			}
		case broadcast := <-Manager.Broadcast: //1->2
			message := broadcast.Message
			sendid := broadcast.Client.SendId //2->1
			id := broadcast.Client.ID         //1->2
			splitId := strings.Split(id, "->")
			// 编辑发送的消息
			smsg := &JsonMsg{
				SendId: splitId[0],
				Msg:    message,
			}

			// 对机器人特殊处理
			// 判断是否为机器人
			if splitId[1] == "00000000000000000" {
				fmt.Println("对方是机器人")
				msgs, _ := Chat.GetChat(string(message))
				InsertMsg := Models.NewMsg{
					SendId:   splitId[0],
					AcceptId: splitId[1],
					State:    false,
					Msg:      string(message),
					Type:     0,
				}
				if err := Mysql.InsertMsg(InsertMsg); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				replymsg := ReplyMsg{
					SendId: splitId[1],
					Code:   Models.CodeSuccess,
					Msg:    string(msgs),
				}
				time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
				msg, _ := json.Marshal(replymsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				InsertMsg = Models.NewMsg{
					SendId:   splitId[1],
					AcceptId: splitId[0],
					State:    false,
					Msg:      string(msgs),
					Type:     0,
				}
				if err := Mysql.InsertMsg(InsertMsg); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				if err := Redis.RedisDelete(id); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				if err := Redis.RedisDelete(splitId[1] + "->" + splitId[0]); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				continue
			}

			flag := false // 默认对方不在线
			for id, conn := range Manager.Client {
				if id != sendid {
					continue
				}
				select {
				case conn.Send <- smsg:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Client, conn.ID)
				}
			}
			if flag {
				fmt.Println("对方在线")
				replymsg := ReplyMsg{
					SendId: splitId[0],
					Code:   Models.CodeWebSocketReply,
					Msg:    Models.CodeWebSocketReply.Msg(),
				}
				msg, _ := json.Marshal(replymsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				InsertMsg := Models.NewMsg{
					SendId:   splitId[0],
					AcceptId: splitId[1],
					State:    false,
					Msg:      string(message),
					Type:     0,
				}
				if err := Mysql.InsertMsg(InsertMsg); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				if err := Redis.RedisDelete(id); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
				if err := Redis.RedisDelete(splitId[1] + "->" + splitId[0]); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}
			} else {
				fmt.Println("对方不在线")
				// 如果对方不在线，查看是否在app内
				flag2 := false                    // 默认对方不在app
				idWithApp := splitId[1] + "->app" //2->app
				for id, conn := range Manager.Client {
					if id != idWithApp {
						continue
					}
					select {
					case conn.Send <- smsg:
						flag2 = true
					default:
						close(conn.Send)
						delete(Manager.Client, conn.ID)
					}
				}
				if flag2 {
					fmt.Println("但在app")
				} else {
					fmt.Println("且不在app")
				}

				replymsg := ReplyMsg{
					SendId: splitId[0],
					Code:   Models.CodeWebSocketReplyNoOne,
					Msg:    Models.CodeWebSocketReplyNoOne.Msg(),
				}
				fmt.Println(replymsg)
				msg, _ := json.Marshal(replymsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				InsertMsg := Models.NewMsg{
					SendId:   splitId[0],
					AcceptId: splitId[1],
					State:    true,
					Msg:      string(message),
					Type:     0,
				}
				if err := Mysql.InsertMsg(InsertMsg); err != nil {
					zap.L().Error("Verifi Parm Error", zap.Error(err))
				}

			}

		}
	}
}

package Controler

import "go-chatbot/Models"

type _ResponsePostList struct {
	Code    Models.ResCode `json:"code"`    // 响应状态码
	Message string         `json:"message"` // 提示信息
	Data    []interface{}  `json:"data"`    // 数据
}

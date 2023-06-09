package Models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
{
	"code":1000,
	"msg":xx
	"data":{}
}
*/

// 回应码
type ResponseData struct {
	Code ResCode
	Msg  interface{}
	Data interface{}
}

// 返回Error
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// 返回自定义MSG的Error
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

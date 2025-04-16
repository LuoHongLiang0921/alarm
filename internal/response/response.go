package response

import (
	resultcode "alarm/internal/result_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义了一个基础的响应结构体，包含状态码、消息和数据字段。
type Response struct {
	Code    resultcode.Code `json:"code"`    // HTTP 状态码
	Message string          `json:"message"` // 响应消息
	Data    interface{}     `json:"data"`    // 响应数据
}

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    resultcode.OK,
		Message: resultcode.OK.String(),
	})
}

func Resp(ctx *gin.Context, resp *Response) {
	ctx.JSON(http.StatusOK, resp)
}

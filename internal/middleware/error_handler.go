package middleware

import (
	"alarm/internal/response"
	"github.com/gin-gonic/gin"
)

// ErrorHandler 是一个全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Resp(c, &response.Response{
					Code:    500,
					Message: "Internal Server Error",
					Data:    nil,
				})
			}
		}()
		c.Next()
	}
}
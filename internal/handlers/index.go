package handlers

import (
	"alarm/internal/response"
	"alarm/internal/service"
	"alarm/internal/transformer"
	"github.com/gin-gonic/gin"
)

type IIndex interface {
	Get(ctx *gin.Context)
}

type index struct{}

// Get 处理 GET / 请求
func (i *index) Get(c *gin.Context) {
	// 调用服务层逻辑
	resp := service.IndexService.GetIndexResponse()

	// 使用 Transformer 转换响应
	indexResp := transformer.Index.ToIndex(resp)

	// 返回响应
	response.Resp(c, indexResp)
}

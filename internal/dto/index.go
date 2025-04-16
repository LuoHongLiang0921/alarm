package dto

// IndexResponse 定义了 Index 接口的响应结构体
type IndexResponse struct {
	ID        int64       `json:"id"`        // ID
	Message   string      `json:"message"`   // 响应消息
	Data      interface{} `json:"data"`      // 响应数据
	Status    int         `json:"status"`    // 响应状态码
	Timestamp int64       `json:"timestamp"` // 时间戳
}

package transformer

import (
	"alarm/internal/dto"
	"alarm/internal/response"
	resultcode "alarm/internal/result_code"
)

type IIndex interface {
	ToIndex(response *dto.IndexResponse) *response.Response
}

type index struct{}

// ToIndex 实现 IIndex 接口的 ToIndex 方法
func (i *index) ToIndex(indexData *dto.IndexResponse) *response.Response {
	return &response.Response{
		Code:    resultcode.OK,
		Message: indexData.Message,
		Data:    indexData.Data,
	}
}

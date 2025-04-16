package service

import (
	"alarm/internal/dao"
	"alarm/internal/dto"
	"alarm/internal/logger"
	"go.uber.org/zap"
)

// IIndexService 定义了 Index 服务的接口
type IIndexService interface {
	GetIndexResponse() *dto.IndexResponse
}

// indexService 是 IndexService 的具体实现
type indexService struct{}

// NewIndexService 创建一个新的 IndexService 实例
func NewIndexService() IIndexService {
	return &indexService{}
}

// GetIndexResponse 返回 Index 的响应数据
func (s *indexService) GetIndexResponse() *dto.IndexResponse {
	logger.Logger.Info("GetIndexResponse")
	index, err := dao.IndexDAO.GetIndex()
	if err != nil {
		logger.Logger.Error("Failed to get index", zap.Error(err))
		return &dto.IndexResponse{
			Message: "Failed to fetch index",
		}
	}
	return &dto.IndexResponse{
		ID:      index.ID,
		Message: index.Name,
	}
}

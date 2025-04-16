package main

import (
	"alarm/internal/dao"
	"alarm/internal/handlers"
	"alarm/internal/infra"
	"alarm/internal/logger"
	"alarm/internal/net/http"
	"alarm/internal/service"
	"alarm/internal/transformer"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 初始化日志
	logger.Init()

	provider, cleanupFunc := infra.InitProvider()
	defer cleanupFunc()

	// 初始化服务层
	service.Init()

	// 初始化 DAO 层
	dao.Init(provider.DB)

	// 初始化转换器层
	transformer.Init()

	// 初始化处理器层
	handlers.Init()

	// 注册路由
	http.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":2025"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

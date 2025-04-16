package http

import (
	"alarm/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes 注册路由
func SetupRoutes(r *gin.Engine) {
	//v1 := r.Group("/v1")
	// 注册 index 路由
	r.GET("/", handlers.Index.Get)
}

//func (s *Server) Routes() {
//	//v1 := s.Group("/v1")
//	s.GET("/", handlers.Index.Get)
//}

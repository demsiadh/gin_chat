package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

// Router gin路由
func Router() *gin.Engine {
	server := gin.Default()
	server.GET("/index", service.GetIndex)
	server.GET("/user/getUserBasicList", service.GetUserBasicList)
	return server
}

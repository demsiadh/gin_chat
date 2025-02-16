package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	server := gin.Default()
	server.GET("/index", service.GetIndex)
	return server
}

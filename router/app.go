package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router gin路由
func Router() *gin.Engine {
	server := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/index", service.GetIndex)
	server.GET("/user/getUserBasicList", service.GetUserBasicList)
	server.GET("/user/createUser", service.CreateUser)
	server.GET("/user/deleteUser", service.DeleteUser)
	server.POST("/user/updateUser", service.UpdateUser)
	server.POST("/user/loginUser", service.LoginUser)
	return server
}

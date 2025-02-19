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

	// 静态资源
	server.Static("/asset", "asset/")
	server.LoadHTMLGlob("views\\**\\*")
	server.GET("/", service.GetIndex)
	server.GET("/index", service.GetIndex)

	// 注册
	server.GET("/register", service.Register)
	// 用户模块
	server.GET("/user/getUserBasicList", service.GetUserBasicList)
	server.POST("/user/createUser", service.CreateUser)
	server.GET("/user/deleteUser", service.DeleteUser)
	server.POST("/user/updateUser", service.UpdateUser)
	server.POST("/user/loginUser", service.LoginUser)

	server.GET("/message/sendMsg", service.SendMsg)
	server.GET("/message/sendUserMsg", service.SendUserMsg)
	return server
}

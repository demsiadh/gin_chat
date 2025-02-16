package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
	server := router.Router()
	// 服务器端口
	err := server.Run(":8088")
	if err != nil {
		panic(err.Error())
	}
}

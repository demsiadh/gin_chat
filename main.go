package main

import "ginchat/router"

func main() {
	server := router.Router()
	// 服务器端口
	err := server.Run(":8088")
	if err != nil {
		panic(err.Error())
	}
}

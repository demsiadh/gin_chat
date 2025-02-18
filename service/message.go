package service

import (
	"fmt"
	"ginchat/common"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 定义一个全局的WebSocket升级器，用于将HTTP连接升级到WebSocket连接
// CheckOrigin函数始终返回true，表示接受任何来源的WebSocket连接请求
var (
	upGrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// SendMsg 函数用于处理WebSocket连接，并向客户端发送消息
// 参数context: 包含了HTTP请求和响应的上下文信息
func SendMsg(context *gin.Context) {
	// 使用升级器将HTTP连接升级到WebSocket连接
	ws, err := upGrade.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		// 如果升级失败，返回500错误码和错误信息
		context.JSON(http.StatusInternalServerError, common.NewErrorResponse("升级协议失败"))
		return
	}
	// 在函数退出时关闭WebSocket连接
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println("关闭连接失败: ", ws, err)
			return
		}
	}(ws)
	// 调用msgHandler函数处理WebSocket连接
	msgHandler(ws, context)
}

// msgHandler 函数用于接收和发送消息
// 参数ws: 代表WebSocket连接
// 参数context: 包含了HTTP请求和响应的上下文信息
func msgHandler(ws *websocket.Conn, context *gin.Context) {
	for {
		// 从Redis订阅频道获取消息
		msg, err := utils.Subscribe(context, utils.RedisPre)
		if err != nil {
			fmt.Println("订阅失败: ", err)
			return
		}
		// 获取当前时间并格式化
		now := time.Now().Format("2006-01-02 15:04:05")
		// 将时间和消息格式化
		m := fmt.Sprintf("[ws][%s]:%s", now, msg)
		// 发送消息到WebSocket连接
		if err := ws.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			fmt.Println("消息发送失败: ", err)
			return
		}
	}
}

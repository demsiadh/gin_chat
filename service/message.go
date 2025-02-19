package service

import (
	"encoding/json"
	"fmt"
	"ginchat/common"
	"ginchat/models"
	"ginchat/utils"
	"github.com/fatih/set"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"sync"
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
		// TODO 这里暂时接不到消息，因为没有消息发送到Redis
		if err := ws.WriteMessage(websocket.TextMessage, []byte(m)); err != nil {
			fmt.Println("消息发送失败: ", err)
			return
		}
	}
}

type Node struct {
	Conn      *websocket.Conn //连接
	DataQueue chan []byte     //消息
	GroupSets set.Interface   //好友 / 群
}

// 映射关系
var clientMap = make(map[int]*Node)

// 读写锁
var rwLocker sync.RWMutex

var udpSendChan = make(chan []byte, 1024)

func SendUserMsg(context *gin.Context) {
	// 1.获取参数
	query := context.Request.URL.Query()
	tempUserId := query.Get("userId")
	userId, _ := strconv.Atoi(tempUserId)
	//tempTargetId := query.Get("targetId")
	//targetId, _ := strconv.Atoi(tempTargetId)
	//content := query.Get("content")
	//msgType := query.Get("type")
	//if tempUserId == "" || tempTargetId == "" || content == "" || msgType == "" {
	//	context.JSON(http.StatusBadRequest, common.NewErrorResponse("参数错误"))
	//	return
	//}

	// TODO 2.校验Token
	isvalida := true
	// 升级协议
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, common.NewErrorResponse("升级协议失败"))
		return
	}

	// 3.获取连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	// 4.用户关系

	// 5.userid和node绑定并且加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// 6.完成发送
	go sendProc(node)
	// 7.完成接收
	go recvProc(node)
	sendMsg(userId, []byte("欢迎来到聊天室"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("发送消息失败: ", err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息失败: ", err)
			return
		}
		udpSendChan <- data

		fmt.Println("收到消息: ", string(data))
	}
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			fmt.Println("udpSendProc  data :", string(data))
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成udp数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	var buf [1024]byte
	for {
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc  data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := models.Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		sendMsg(msg.TargetId, data)
	}
}

func sendMsg(userId int, msg []byte) {
	rwLocker.RLock()
	defer rwLocker.RUnlock()
	node, ok := clientMap[userId]
	if ok {
		node.DataQueue <- msg
	}
}

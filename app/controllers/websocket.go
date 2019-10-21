//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: websocket.go
//@Time: 2019/10/21 下午2:41

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 定义向webSocket客户端传输的广播消息结构体类型
type BroadcastMessageStructure struct {
	Message     string `json:"message"`     // 消息
	TipsMessage string `json:"tipsMessage"` // 提示消息
}

var (
	Clients   = make(map[*websocket.Conn]bool)       // 链接客户端字典
	Broadcast = make(chan BroadcastMessageStructure) // 广播channel
)

var upgrade = websocket.Upgrader{ // Configure the upgrade
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	go handleSendMessages() // 单独起一个goroutine 负责向所有WebSocket客户端广播消息
}

// 处理webSocket客户端链接请求
func HandleConnections(c *gin.Context) {
	// 将初始GET请求升级到websocket
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("get upgrade webSocket err:", err)
		return
	}
	// 确保在函数返回时关闭连接o
	defer func() { _ = ws.Close() }()
	// 注册新的webSocket客户端
	Clients[ws] = true
	log.Printf("webSocket remoteAddr:%s connected!", ws.RemoteAddr())
	Broadcast <- BroadcastMessageStructure{TipsMessage: fmt.Sprintf("%s 上线！", ws.RemoteAddr())}
	//接收webSocket客户端发来的数据
	for {
		n, buffer, err := ws.ReadMessage()
		if err != nil {
			log.Printf("webSocket remoteAddr:%s is close!", ws.RemoteAddr())
			Broadcast <- BroadcastMessageStructure{TipsMessage: fmt.Sprintf("%s 下线！", ws.RemoteAddr())}
			delete(Clients, ws)
			break
		}
		// 打印WebSocket发送的Msg
		msg := fmt.Sprintf("webSocket IP:%s len:%d msg:%s\n", ws.RemoteAddr(), n, string(buffer))
		Broadcast <- BroadcastMessageStructure{Message: msg}

		//log.Println(msg)
	}
}

// 发送广播至所有WebSocket客户端页面
func handleSendMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				_ = client.Close()
				delete(Clients, client)
			}
		}
	}
}

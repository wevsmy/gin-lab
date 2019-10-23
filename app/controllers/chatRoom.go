//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: chatRoom.go
//@Time: 2019/10/23 上午11:27

package controllers

import (
	"fmt"
	"github.com/dustin/go-broadcast"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
)

// 聊天室

var roomChannels = make(map[string]broadcast.Broadcaster)

func openListener(roomid string) chan interface{} {
	listener := make(chan interface{})
	room(roomid).Register(listener)
	return listener
}

func closeListener(roomid string, listener chan interface{}) {
	room(roomid).Unregister(listener)
	close(listener)
}

func deleteBroadcast(roomid string) {
	b, ok := roomChannels[roomid]
	if ok {
		b.Close()
		delete(roomChannels, roomid)
	}
}

func room(roomid string) broadcast.Broadcaster {
	b, ok := roomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		roomChannels[roomid] = b
	}
	return b
}

// ----------------------

func RoomStream(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	defer closeListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

func RoomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())

	c.HTML(http.StatusOK, "chat_room.tmpl.html", gin.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func RoomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	room(roomid).Submit(userid + ": " + message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func RoomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	deleteBroadcast(roomid)
}

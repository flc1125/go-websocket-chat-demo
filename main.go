package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var control = NewControl()

func handler(response http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		conn.Close()
		return
	}

	client := NewClient(conn) // 创建设备信息
	control.Subscribe(client) // 设备连接中控台

	for {
		// 接受消息
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			conn.Close()
			control.UnSubscribe(client) // 设备断开中控台
			return
		}

		// 发送内容
		p = []byte(fmt.Sprintf(
			"【%s】在 %s 说：%s",
			client.Id,
			time.Now().Format("2006-01-02 15:04:05"),
			string(p)),
		)

		message := NewMessage(client, p) // 创建消息体

		control.Send(message) // 广播消息
	}
}

func handlerIndex(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/ws", handler)

	http.ListenAndServe("0.0.0.0:8080", nil)
}

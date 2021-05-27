package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// 中控台
type Control struct {
	Clients  map[string]*Client
	Messages chan *Message
}

// 创建一个中控
func NewControl() *Control {
	c := &Control{
		Clients:  make(map[string]*Client),
		Messages: make(chan *Message, 1000),
	}

	c.Broadcast() // 广播中心

	return c
}

// 设备号订阅
func (c *Control) Subscribe(client *Client) *Control {
	c.Clients[client.Id] = client

	// 夹道欢迎
	msg := fmt.Sprintf("【%s】进入房间……当前在线人数【%d】", client.Id, len(c.Clients))
	control.Send(NewMessage(client, []byte(msg)))

	return c
}

// 设备号取消订阅
func (c *Control) UnSubscribe(client *Client) *Control {
	delete(c.Clients, client.Id)

	// 奔走相告
	msg := fmt.Sprintf("【%s】离开房间……当前在线人数【%d】", client.Id, len(c.Clients))
	control.Send(NewMessage(client, []byte(msg)))

	return c
}

// 消息广播（全体）
func (c *Control) Broadcast() {
	go func() {
		for {
			message := <-c.Messages

			// time.Sleep(time.Second * 2) // 模拟当设备多时，消息发送延迟时间

			for _, client := range c.Clients {
				client.Conn.WriteMessage(websocket.TextMessage, message.Content)
			}
		}
	}()
}

// 发送消息
func (c *Control) Send(message *Message) {
	c.Messages <- message
}

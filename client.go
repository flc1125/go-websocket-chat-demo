package main

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

type Client struct {
	Conn *websocket.Conn
	Id   string
}

// 生成唯一标识（设备号）
func generateId() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn: conn,
		Id:   generateId(),
	}
}

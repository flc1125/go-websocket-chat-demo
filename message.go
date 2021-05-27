package main

type Message struct {
	Client  *Client
	Content []byte
}

func NewMessage(client *Client, content []byte) *Message {
	return &Message{
		Client:  client,
		Content: content,
	}
}

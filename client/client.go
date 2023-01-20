package main

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	Connect_Server net.Conn
	ListenerPort   string
}

func InitClient(ListenerPort string) *Client {
	return &Client{
		ListenerPort: ListenerPort,
	}
}

func (c *Client) ConnectClient() error {
	connect_server, err := net.Dial("tcp", c.ListenerPort)

	if err != nil {
		log.Println(err.Error())
	}

	c.Connect_Server = connect_server

	defer c.Connect_Server.Close()

	log.Println("Connect Success Server_____")

	c.SendMsg()

	return nil
}

func (c *Client) SendMsg() {
	for {
		var send_msg string

		fmt.Scanln(&send_msg)

		c.Connect_Server.Write([]byte(send_msg))

		go c.RecvComment()
	}
}

func (c *Client) RecvComment() {
	recvBuf := make([]byte, 2048)

	for {
		recv, err := c.Connect_Server.Read(recvBuf)

		if err != nil {
			log.Println(err.Error())
		}

		msg := recvBuf[:recv]

		log.Println(string(msg))
	}
}

func main() {
	client := InitClient(":8000")

	client.ConnectClient()
}

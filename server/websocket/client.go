package websocket

import (
	"fmt"
	"runtime/debug"

	"github.com/gorilla/websocket"
)

const heartbeatExpireTime = 6 * 60



type Client struct {
	Addr string // 客户端地址
	Socket *websocket.Conn // 用户连接
	Send chan []byte // 待发送的数据
	Receive chan []byte // 待接受的数据
	AppId uint32 // 登录的平台Id app/web/ios
	UserId string // 用户Id，用户登录以后才有
	FirstTime uint64 // 首次连接事件
	HeartbeatTime uint64 // 用户上次心跳时间
	LoginTime uint64 // 登录时间 登录以后才有
}

type Message struct {

	FromUserId string
	ToUserId
	Body *MessageBody
}

//消息类型 简单定义三种
const (
	MessageTypeText = 1
	MessageTypeImage = 2
	MessageTypeVideo = 3
)

//消息体
type MessageBody struct {
	Type uint8
	Content string
}

type login struct {
	AppId uint32
	UserId string
	Client *Client
}

//初始化客户端连接
func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client := &Client{
		Addr: addr,
		Socket: socket,
		Send: make(chan []byte, 100),
		Receive: make(chan []byte, 100),
		FirstTime: firstTime,
		HeartbeatTime: firstTime,
	}
	return
}


func deferError() {
	if r := recover(); r != nil {
		fmt.Println("数据读取停止", string(debug.Stack()), r)
	}
}

//从客户端读数据
func (c *Client) read()  {
	//捕获错误并打印
	defer deferError()

	//读取channel关闭
	defer func ()  {
		fmt.Println("关闭读取客户端数据通道send", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()

		if err != nil {
			fmt.Println("读取客户端数据错误", c.Addr, err)
			return
		}

		fmt.Println("读取客户端数据处理", string(message))

		//todo processData
	}
}


//向客户端写消息
func (c *Client) write() {

	defer deferError()

	defer func ()  {
		ClientManager.Unregister <- c
		c.Socket.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Receive
			if !ok {
				// 发送数据错误 关闭连接
				fmt.Println("Client发送数据 关闭连接", c.Addr, "ok", ok)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
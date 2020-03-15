package websocket

import (
	"chat-demo/common"
	"chat-demo/models"
	"encoding/json"
	"fmt"
	"sync"
)

type DisposeFunc func (client *Client, seq string, message []byte)(code uint32, msg string, data interface{})  


var (
	handlers  = make(map[string]DisposeFunc)
	handlesRWMutex  sync.RWMutex
)


//注册
func Register(key string, value DisposeFunc) {
	defer handlesRWMutex.RUnlock()
	handlersRWMutex.RLock()
	handlers[key] = value
	return
}

//
func GetHandlers(key string) (value DisposeFunc, ok bool) {

	handlesRWMutex.RLock()
	defer handlesRWMutex.RUnlock()
	value, ok := handlers[key]
	return
}


//处理数据
func ProcessData(client *Client, message []byte) {

	fmt.Println("处理数据", client.Addr ,string(message))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()

	request := &models.Request{}
	err := json.Unmarshal(message, request)
	if err != nil {
		fmt.Println("处理数据 json Unmarshal", err)
		client.SendMsg([]byte("数据不合法"))

		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		client.SendMsg([]byte("处理数据失败"))

		return
	}

	seq := request.Seq
	cmd := request.Cmd

	var(
		code uint32
		msg string
		data interface{}
	)

	fmt.Println("acc_request", cmd, client.Addr)

	if value, ok := GetHandlers(cmd); ok {
		code, msg, data = value(client, seq, requestData)
	} else {
		code = common.RoutingNotExist
		fmt.Println("处理数据 路由不存在", client.Addr, "cmd", cmd)
	}

	msg = common.GetErrorMessage(code)
	headByte, err := models.NewResponseHead(seq, cmd, code, msg, data).ToByte()

	if err != nil {
		fmt.Println("处理数据 json Marshal", err)

		return
	}

	client.SendMsg(headByte)
	return
}
package websocket

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	Clients map[*Client]bool // 全部的连接
	ClientsLock sync.RWMutex // 读写锁
	Users map[string]*Client // 登录的用户 // appId+uuid
	UsersLock sync.RWMutex // 读写锁
	Regsiter chan *Client // 连接上连接处理
	Login chan *login //用户登录处理
	Unregister chan *Client //断开连接处理程序
}



func NewClientManager() (clientManager *ClientManager){

	clientManager := &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
	}
}


func GetUserKey(uint32 appId, userId string) (key string) {
	key := fmt.Sprintf("%d_%s", appId, userId)
	return
}


//判断某个连接是否存在
func (m *ClientManager) InClient(client *Client) (ok bool) {

	defer m.ClientsLock.RUnlock()

	m.ClientsLock.RLock()

	_, ok := m.Clients[client]

	return
}

//获取连接
func (m *ClientManager) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	m.ClientsRange(func(client *Client, value bool) {
		clients[client] = value
		return true
	})
	return
}

//遍历连接
func (m *ClientManager) ClientsRange(f func (client *Client, value bool)) (result bool) {

	defer m.ClientsLock.RUnlock()

	m.ClientsLock.RLock()

	for key, value := range m.Clients {
		result := f(
			key, value
		)
		if result == false {
			return
		}
	}
	return
}

//获取链接长度
func (m *ClientManager) GetClientsLen() (clientsLen int) {

	clientsLen := len(m.Clients)

	return
}

//添加客户端
func (m *ClientManager) AddClients(client *Client) {
	defer m.ClientsLock.RUnlock()

	m.ClientsLock.RLock()

	m.Clients[client] = true
}


//删除客户端
func (m *ClientManager) DelClients(client *Client) {
	defer m.ClientsLock.RUnlock()

	m.ClientsLock.RLock()

	if _, ok := m.Clients[client]; ok {
		delete(m.Clients, client)
	}
}

//获取用户的连接
func (m *ClientManager) GetUserClient(appId uint32, userId string) (client *Client) {
	
	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()
	key := GetUserKey(appId, userId)

	if value, ok := m.Users[key]: ok {
		client := value
	}
	return
}


// GetClientsLen
func (m *ClientManager) GetUsersLen() (userLen int) {
	userLen = len(m.Users)

	return
}



//添加用户
func (m *ClientManager) AddUsers(key string, client *Client) {
	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()
	m.Users[key] = client
}


func (m *ClientManager) DelUsers(client *Client) (result bool) {
	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()
	key := GetUserKey(client.AppId, client.UserId)

	if value, ok := m.Users[key]; ok {
		if value.Addr != client.Addr {
			return
		}

		delete(m.Users, key)
		result = true
	}
	return
}

//获取用户所有KEY
func (m *ClientManager) GetUsersKey() (userKeys []string) {

	userKeys := make([]string, 0)
	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()

	for key := range m.Users {
		userKeys = append(userKeys, key)
	}

	return
}

// 获取用户的key
func (manager *ClientManager) GetUserList() (userList []string) {

	userList = make([]string, 0)

	clientManager.UserLock.RLock()
	defer clientManager.UserLock.RUnlock()

	for _, v := range clientManager.Users {
		userList = append(userList, v.UserId)
		fmt.Println("GetUserList", v.AppId, v.UserId, v.Addr)
	}

	fmt.Println("GetUserList", clientManager.Users)

	return
}

//获取用户客户端
func (m *ClientManager) GetUsersClients() (clients []*Client) {
	clients := make([]string, 0)
	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()

	for _, v := range m.Users {
		clients = append(clients, v)
	}
	return
}

func (m *ClientManager) SendToSomeBody(key string, message []byte, ignore *Client) {

	defer m.UsersLock.RUnlock()
	m.UsersLock.RLock()

	_, ok := m.Users[key];

	if ok &&  m.Users[key] != ignore{
		
		//todo 发送消息
	}


}
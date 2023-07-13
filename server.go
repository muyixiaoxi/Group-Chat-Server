package main

import (
	"fmt"
	"net"
	"time"
)

var clientsMap = make(map[string]net.Conn) //储存当前群聊中所有用户连接信息

// 监听
func listenClient() {
	listener, err := net.Listen("tcp", ":8000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.Listen err = ", err)
	}
	for {
		conn, err := listener.Accept() //阻塞等待用户连接
		if err != nil {
			fmt.Println("listener.Accpet err = ", err)
		}
		clientsMap[conn.RemoteAddr().String()] = conn

		//获取用户连接，将用户添加至map集合
		conn.RemoteAddr()

		go addReceiver(conn)
	}

}

// 向连接添加接收器
func addReceiver(newConnect net.Conn) {
	for {
		byteMsg := make([]byte, 2048)
		n, err := newConnect.Read(byteMsg)
		if err != nil {
			newConnect.Close()
		}
		fmt.Println(string(byteMsg[:n]))
		msgBroadcast(byteMsg[:n], newConnect.RemoteAddr().String())
	}
}

// 广播所有client
func msgBroadcast(byteMsg []byte, key string) {
	for k, con := range clientsMap {
		if k != key {
			con.Write(byteMsg)
		}
	}
}

func main() {
	initGroupChatServer()
}

// 初始化
func initGroupChatServer() {
	fmt.Println("服务器已启动...")
	time.Sleep(time.Second)
	fmt.Println("等待客户端请求连接...")
	listenClient()
	select {}
}

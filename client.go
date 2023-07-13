package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 用户名
var loginName string

// 本机连接
var selfConnect net.Conn

// 读取行文字
var reader = bufio.NewReader(os.Stdin)

// 初始化
func initGroupChatClient() {
	fmt.Println("请问您怎么称呼？")
	bName, _, _ := reader.ReadLine()
	loginName = string(bName)
	connect(":8000")
	select {}
}

//建立连接
func connect(addr string) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)
	fmt.Println("tcpAddr:", tcpAddr)
	conn, err := net.Dial("tcp", addr)
	selfConnect = conn
	if err != nil {
		fmt.Println("连接服务器失败")
		os.Exit(1)
	}

	go msgSender()
	go msgReceiver()
}

//消息接收器
func msgReceiver() {
	buf := make([]byte, 2048)
	for {
		n, _ := selfConnect.Read(buf) //从建立连接的缓冲区读消息
		fmt.Println(string(buf[:n]))
	}
}

//消息发送器
func msgSender() {
	for {
		bMsg, _, _ := reader.ReadLine()
		bMsg = []byte(loginName + ":" + string(bMsg))
		selfConnect.Write(bMsg) //发消息
	}
}

func main() {
	initGroupChatClient()
}

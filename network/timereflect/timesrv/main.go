package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr("tcp4", ":8080")
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpServer)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}
	for {
		// conn 就是已经建立连接的socket对象
		conn, err := listener.Accept()
		if err != nil {
			// 记录错误日志
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// 一般是服务端做主动关闭，防止客户端恶意保持长连接
	// 服务端关闭连接的时候，也会清空socket的发送缓冲区
	defer func() { conn.Close() }()

	socketSendCh := make(chan string, 1)
	defer close(socketSendCh)

	done1 := make(chan struct{})
	done2 := make(chan struct{})
	defer close(done1)
	defer close(done2)

	defer func() {
		if err := recover(); err != nil {
			// 记录异常日志
		}
	}()

	go generateTime(socketSendCh, done1)
	go sendBuff(conn, socketSendCh, done2)

	<-done1
	<-done2
}

func generateTime(socketSendCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	// 生成时间，然后向socketSendCh发送数据
	timeTemplate := "2006-01-02 15:04:05" //常规类型

	t := time.Now().Format(timeTemplate)

	fmt.Println(t)

	socketSendCh <- t
}

func sendBuff(conn net.Conn, socketSendCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	// 从socketSendCh去除数据，发送数据到客户端
	buff := <-socketSendCh

	// 发送数据到客户端
	conn.Write([]byte(buff))
}

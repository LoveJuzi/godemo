package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	client()
}

func client() {
	server := "127.0.0.1:8080"
	addr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	//建立tcp连接
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	handle(conn) // 无需启动一个协程，因为客户端只需要处理一个连接
}

func handle(conn net.Conn) {
	standardOutputCh := make(chan string, 1)
	defer close(standardOutputCh)
	done1 := make(chan struct{})
	defer close(done1)
	done2 := make(chan struct{})
	defer close(done2)

	go receiveBuff(conn, standardOutputCh, done1)
	go printTime(standardOutputCh, done2)

	<-done1
	<-done2
}

func receiveBuff(conn net.Conn, standardOutputCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	//接收响应
	response, err := ioutil.ReadAll(conn)
	if err != nil {
		// 记录错误日志
		// 此处不退出程序，因为可能是服务端提前关闭导致的一个错误
	}
	standardOutputCh <- string(response)
}

func printTime(standardOutputCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()
	t := <-standardOutputCh
	fmt.Println(t)
}

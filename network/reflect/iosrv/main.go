package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	serve()
}

func serve() {
	// 申请IP地址空间 绑定端口
	tcpServer, err := net.ResolveTCPAddr("tcp4", ":8080")
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	// 开启监听
	listner, err := net.ListenTCP("tcp", tcpServer)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	// 等待连接
	for {
		conn, err := listner.Accept()
		if err != nil {
			// 记录错误日志
			continue
		}

		go handle(conn)
	}
}

// CHEOF 通道结束符
var CHEOF = string([]byte{0x01})

// 各种资源的管理都在handle中处理
func handle(conn net.Conn) {
	defer func() { conn.Close() }() // 服务端关闭连接

	done := [3]chan struct{}{}
	for idx := range done {
		done[idx] = make(chan struct{}, 1)
	}

	sktSendCh := make(chan string, 1)
	sdOutCh := make(chan string, 1)

	defer func() {
		if err := recover(); err != nil {
			// 记录异常日志
			fmt.Println(err)
		}
	}()

	go recieveBuff(conn, sktSendCh, sdOutCh, done[0])
	go sendBuff(conn, sktSendCh, done[1])
	go printBuff(sdOutCh, done[2])

	<-done[0]

	// 结束sendBuff和printBuff
	sktSendCh <- CHEOF
	sdOutCh <- CHEOF

	<-done[1]
	<-done[2]
}

func recieveBuff(conn net.Conn, sktSendCh chan string, sdOutCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	reader := bufio.NewReader(conn)
	var ba [1 << 10]byte
	for {
		n, err := reader.Read(ba[:]) // 这里可能阻塞
		if err != nil {
			if err == io.EOF {
				fmt.Println("接收到EOF")
				break
			}
			// 记录错误日志
			panic("read error")
		}
		sktSendCh <- string(ba[:n])
		sdOutCh <- string(ba[:n])
	}
}

func sendBuff(conn net.Conn, sktSendCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	for {
		buff := <-sktSendCh

		if buff == CHEOF {
			break
		}

		var err error
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(buff)
		if err != nil {
			// 记录错误日志
			panic("write error")
		}

		err = writer.Flush()
		if err != nil {
			// 记录错误日志
			panic("write flush error")
		}
	}
}

func printBuff(sdOutCh chan string, done chan struct{}) {
	defer func() { done <- struct{}{} }()

	for {
		buff := <-sdOutCh

		if buff == CHEOF {
			break
		}

		fmt.Print(buff)
	}
}

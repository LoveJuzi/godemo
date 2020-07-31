package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	serve()
}

func bgtask(f func(), ech chan<- struct{}, rfs ...func()) {
	go func() {
		defer recover()                      // 捕获一切异常
		defer func() { ech <- struct{}{} }() // 发送退出信号
		defer func() {                       // 捕获任务panic
			if err := recover(); err != nil {
				// 记录异常日志
				fmt.Println(err)

				for _, v := range rfs {
					v()
				}
			}
		}()

		f()
	}()
}

func synctask(f func(), rfs ...func()) <-chan struct{} {
	d := make(chan struct{})
	go func() {
		defer recover()                    // 捕获一切异常
		defer func() { d <- struct{}{} }() // 发送退出信号
		ch := make(chan struct{})          // 后台任务退出信号
		bgtask(f, ch, rfs...)              // 后台处理任务
		<-ch                               // 等待后台任务结束
	}()
	return d
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

		// 这其实是一个后台任务，但是由于我们不关注退出信号，
		// 使用同步任务进行模拟一下
		synctask(taskHandle(conn))
	}
}

// CHEOF 通道结束符
var CHEOF = string([]byte{0x01})

// Handle 各种资源的管理都在handle中处理
func Handle(conn net.Conn) {
	defer func() { conn.Close() }() // 服务端关闭连接

	ch1 := make(chan string)
	ch2 := make(chan string)

	T1 := synctask(taskRecieveBuff(conn, ch1, ch2), func() { close(ch1) }, func() { close(ch2) })
	T2 := synctask(taskSendBuff(conn, ch1), func() { close(ch1) })
	T3 := synctask(taskPrintBuff(ch2), func() { close(ch2) })

	<-T1       // 等待T1任务结束
	close(ch1) // 通知T2任务结束
	close(ch2) // 通知T3任务结束
	<-T2       // 等待T2任务结束
	<-T3       // 等待T3任务结束
}

func taskHandle(conn net.Conn) func() {
	return func() {
		Handle(conn)
	}
}

// RecieveBuff 从socket的接收缓冲区读取数据
func RecieveBuff(conn net.Conn, ch1 chan string, ch2 chan string) {
	var ba [1 << 10]byte
	for {
		n, err := conn.Read(ba[:])
		if err != nil {
			if err == io.EOF {
				fmt.Println("接收到EOF")
				break
			}
			// 记录错误日志
			break
		}
		ch1 <- string(ba[:n])
		ch2 <- string(ba[:n])
	}
}

func taskRecieveBuff(conn net.Conn, sktSendCh chan string, sdOutCh chan string) func() {
	return func() {
		RecieveBuff(conn, sktSendCh, sdOutCh)
	}
}

// SendBuff 向socket的发送缓冲区发送数据
func SendBuff(conn net.Conn, ch <-chan string) {
	for v := range ch {
		_, err := conn.Write([]byte(v))
		if err != nil {
			// 记录错误日志
			break
		}
		fmt.Println(v)
	}
}

func taskSendBuff(conn net.Conn, sktSendCh chan string) func() {
	return func() {
		SendBuff(conn, sktSendCh)
	}
}

// PrintBuff 标准输出输出内容
func PrintBuff(ch <-chan string) {
	for v := range ch {
		fmt.Print(v)
	}
}

func taskPrintBuff(ch <-chan string) func() {
	return func() {
		PrintBuff(ch)
	}
}

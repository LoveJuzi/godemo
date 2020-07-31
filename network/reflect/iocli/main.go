package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	client()
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

func client() {
	// 连接的地址
	server := "127.0.0.1:8080"
	// 连接
	addr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		// 记录错误日志
		os.Exit(1)
	}

	handle(conn)
}

func handle(conn net.Conn) {
	// 通信channel
	sktSendCh := make(chan string, 1)
	sdOutCh := make(chan string, 1)

	// 任务分发
	// 针对不同的任务类型，我们仅仅需要重新编写不同的任务模板，也就是task的实现方式
	// golang语言使用这种方式，可以很好的进行任务并发
	T1 := synctask(taskScanBuff(sktSendCh), func() { close(sktSendCh) })
	T2 := synctask(taskSendBuff(conn, sktSendCh), func() { close(sktSendCh) })
	T3 := synctask(taskRecieveBuff(conn, sdOutCh), func() { close(sdOutCh) })
	T4 := synctask(taskPrintBuff(sdOutCh), func() { close(sdOutCh) })

	// 任务执行顺序控制流
	// 任务控制流最好在主程中进行描述，这样便于后期维护
	// 当然，为了优雅，也可以写到具体的任务中，但是这样会导致一个问题
	// 当任务流变的复杂的时候，具体任务的控制流就变得复杂，不利于维护
	<-T1             // 等待T1任务结束
	close(sktSendCh) // 通知T2结束任务
	<-T3             // 等待T3任务结束
	close(sdOutCh)   // 通知T4结束任务
	<-T2             // 等待T2结束任务
	<-T4             // 等待T3结束任务
}

// ScanBuff 从标准输入读入数据
func ScanBuff(ch chan<- string) {
	// 读文件
	// 读文件就有EOF标识符
	var b string = ""
	for {
		fmt.Scanln(&b)
		fmt.Println(b)
		ch <- b
	}
}

// SendBuff 向socket的发送缓冲区发送数据
func SendBuff(conn net.Conn, ch <-chan string) {
	for b := range ch {
		var err error
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(b)
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

// RecieveBuff 从socket的接收缓冲区接收数据
func RecieveBuff(conn net.Conn, ch chan<- string) {
	b, err := ioutil.ReadAll(conn)
	if err != nil {
		// 记录错误日志
		panic("client recieve error")
	}

	ch <- string(b)
}

// PrintBuff 打印数据到标准输出
func PrintBuff(ch <-chan string) {
	for b := range ch {
		fmt.Printf(b)
	}
}

func taskScanBuff(ch chan<- string) func() {
	return func() {
		ScanBuff(ch)
	}
}

func taskSendBuff(conn net.Conn, ch <-chan string) func() {
	return func() {
		SendBuff(conn, ch)
	}
}

func taskRecieveBuff(conn net.Conn, ch chan<- string) func() {
	return func() {
		RecieveBuff(conn, ch)
	}
}

func taskPrintBuff(ch <-chan string) func() {
	return func() {
		PrintBuff(ch)
	}
}

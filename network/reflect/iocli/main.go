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
	T1 := task(scanBuff(sktSendCh))
	T2 := task(sendBuff(conn, sktSendCh))
	T3 := task(recieveBuff(conn, sdOutCh))
	T4 := task(printBuff(sdOutCh))

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

type taskFun func()

// task 任务整体的管理函数
// f 代表实际的业务函数
// 采用这种方式的一个最大的好处就是，业务代码完全独立出去
// 不需要关心任务的调度，也不用关心任务的异常处理情况
// task 这种编程方式，也给出了一个简单的模型：
//      如果多个函数含有相同的入参，这些入参的功能是相同的，
//      那么我们就可以使用这种函数模型进行统一的包装
func task(f taskFun) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer func() { done <- struct{}{} }() // 任务完成
		defer func() {                        // 异常处理
			if err := recover(); err != nil {
				// 记录异常日志
				fmt.Println(err)
			}
		}()

		f() // 执行具体的任务函数
	}()

	return done
}

// 实际任务函数，但是这里为了演示，直接使用了函数生成式，没有采用包装的形式
func scanBuff(ch chan<- string) taskFun {
	return func() {
		// 读文件
		// 读文件就有EOF标识符
		var b string = "abcddd"
		fmt.Println(b)
		ch <- b
	}
}

func sendBuff(conn net.Conn, ch <-chan string) taskFun {
	return func() {
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
}

func recieveBuff(conn net.Conn, ch chan<- string) taskFun {
	return func() {
		b, err := ioutil.ReadAll(conn)
		if err != nil {
			// 记录错误日志
			panic("client recieve error")
		}

		ch <- string(b)
	}
}

func printBuff(ch <-chan string) taskFun {
	return func() {
		for b := range ch {
			fmt.Printf(b)
		}
	}
}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// task 就是任务模板
// 当我们启动一个新的goroutine的时候，我们相当于启动了一个新的任务
// 每个任务都有可能出现不同程度的异常
// task 任务模板的作用就是描述如何处理这些不同程度的异常
// 以及如何恢复自身任务和其他任务的关系

// 这里写了两个不同的任务模板，一个是“bgtask”，一个是“task”
// bgtask 是一个后台任务模板
// 入参有三种类型，一个是任务函数，一个是退出信号通道
// 另一个是异常恢复函数集
// bgtask的功能是：
//   1. 执行任务函数
//   2. 如果任务函数执行异常，使用异常恢复函数集进行恢复
//   3. 发送退出信号到退出信号通道
//   4. 捕获一切此此函数中的异常
// task 是一个同步任务模板
// 入参有两种类型，一个是任务函数，一个是异常恢复函数集
// task的功能是：
//   1. 执行任务函数（此处使用bgtask进行执行）
//   2. 发送退出信号
//   3. 捕获一切此函数中的异常

// 一系列同质的后台任务，也是一个脱离管理的异步任务，只能通过发送退出信号表示退出
// ech 用来存储任务退出的信号
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

func taskRWMain(ch <-chan int8) func() {
	return func() {
		RWMain(ch)
	}
}

func taskGenerateReader(ch <-chan int8, rs chan<- struct{}) func() {
	return func() {
		GenerateReader(ch, rs)
	}
}

func taskGenerateWriter(ch <-chan int8, ws <-chan struct{}) func() {
	return func() {
		GenerateWriter(ch, ws)
	}
}

func taskReader(r int8) func() {
	return func() {
		Reader(r)
	}
}

func taskWriter(w int8) func() {
	return func() {
		Writer(w)
	}
}

// SIZE 队列大小
const SIZE = 1 << 10

// RWMain 读者写者主程序
// rs 管道看似多余，其实，这是一种通信管理方式
// 通过这种方式，我们将计数的操作都限定在主管理函数
func RWMain(ch <-chan int8) {
	rch := make(chan int8, SIZE)              // 读者队列
	wch := make(chan int8, SIZE)              // 写者队列
	ws := make(chan struct{}, 1)              // 写程序排斥锁
	rs := make(chan struct{}, math.MaxUint32) // 读者退出信号队列
	wg := &sync.WaitGroup{}                   // 正在执行读任务的读者计数器

	// 这个其实是一个后台任务，但是我们并不关心它的退出信号，使用同步任务模拟了一下
	synctask(func() {
		for range rs { // 读者读任务完毕
			wg.Done()
		}
	})
	defer close(rs) // 结束上面的任务，防止goroutine资源不被释放

	waitAllReaderExit := func() { // 检查是否有读者
		wg.Wait()
	}
	waitWriterExit := func() { // 检查是否有写者
		ws <- struct{}{}
		<-ws
	}

	T1 := synctask(taskGenerateReader(rch, rs), func() { close(rch) }, func() { close(rs) })
	T2 := synctask(taskGenerateWriter(wch, ws), func() { close(wch) }, func() { close(ws) })

	for rw := range ch { // 分发任务
		waitWriterExit()
		if rw == 0 {
			wg.Add(1)
			rch <- rw
		} else if rw == 1 {
			ws <- struct{}{}
			waitAllReaderExit() // 等待所有的读者退出
			wch <- rw
		}
	}

	waitAllReaderExit()
	waitWriterExit()

	// 同步任务控制流
	close(rch) // 通知T1任务结束
	close(wch) // 通知T2任务结束
	<-T1       // 等待T1任务结束
	<-T2       // 等待T2任务结束
}

// GenerateReader 生成读者
func GenerateReader(ch <-chan int8, rs chan<- struct{}) {
	for r := range ch {
		fmt.Println("读者队列长度：", len(ch))
		bgtask(taskReader(r), rs) // 启动一个后台读任务
	}
}

// GenerateWriter 生成写者
func GenerateWriter(ch <-chan int8, ws <-chan struct{}) {
	for w := range ch {
		fmt.Println("写者队列长度：", len(ch)+1)
		<-synctask(taskWriter(w)) // 启动一个同步写任务
		<-ws                      // 释放写者锁
	}
}

// Reader 读者
func Reader(r int8) {
	fmt.Println("读者正在读...", r)
	t := time.Second * time.Duration(rand.Intn(5)+2)
	fmt.Println("读时长：", t)
	time.Sleep(t)
}

// Writer 写者
func Writer(w int8) {
	fmt.Println("写者正在写...", w)
	t := time.Second * time.Duration(rand.Intn(5)+2)
	fmt.Println("写时长：", t)
	time.Sleep(t)
}

func main() {
	ch := make(chan int8, SIZE) // 读写主程序的接收队列

	T := synctask(taskRWMain(ch), func() { close(ch) })

	task := []int8{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	for _, v := range task {
		ch <- v
	}

	close(ch)
	<-T

	time.Sleep(time.Second * 10)
}

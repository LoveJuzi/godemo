package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SIZE 队列大小
const SIZE = 1 << 10

func task(f func(), rfs ...func()) <-chan struct{} {
	d := make(chan struct{})
	go func() {
		defer func() { d <- struct{}{} }() // 退出信号
		defer recover()                    // 捕获二次panic
		defer func() {                     // 捕获任务panic
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
	return d
}

func taskRWMain(ch <-chan int8) func() {
	return func() {
		RWMain(ch)
	}
}

// RWMain 读者写者主程序
// rs 管道看似多余，其实，这是一种通信管理方式
// 通过这种方式，我们将计数的操作都限定在主管理函数
func RWMain(ch <-chan int8) {
	rch := make(chan int8, SIZE)     // 读者队列
	wch := make(chan int8, SIZE)     // 写者队列
	ws := make(chan struct{}, 1)     // 写程序排斥
	rs := make(chan struct{}, 1<<31) // 读者完成队列
	wg := &sync.WaitGroup{}          // 正在读的读者计数器

	task(func() { // 读者读取完毕，需要通知计数减少
		for range rs {
			wg.Done()
		}
	})
	defer close(rs) // 结束上面的task，防止goroutine资源不被释放

	waitAllReaderExit := func() {
		wg.Wait()
	}
	waitWriterExit := func() {
		ws <- struct{}{} // 检查是否有写者
		<-ws
	}

	T1 := task(taskGenerateReader(rch, rs), func() { close(rch) })
	T2 := task(taskGenerateWriter(wch, ws), func() { close(wch) })

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

	close(rch)
	close(wch)
	<-T1
	<-T2
}

func taskGenerateReader(ch <-chan int8, rs chan<- struct{}) func() {
	return func() {
		GenerateReader(ch, rs)
	}
}

// GenerateReader 生成读者
func GenerateReader(ch <-chan int8, rs chan<- struct{}) {
	for r := range ch {
		fmt.Println("读者队列长度：", len(ch))
		go func(r int8) { // 这种写法很危险！！！
			<-task(taskReader(r))
			rs <- struct{}{} // 通知读者退出了
		}(r)
	}
}

func taskGenerateWriter(ch <-chan int8, ws <-chan struct{}) func() {
	return func() {
		GenerateWriter(ch, ws)
	}
}

// GenerateWriter 生成写者
func GenerateWriter(ch <-chan int8, ws <-chan struct{}) {
	for w := range ch {
		fmt.Println("写者队列长度：", len(ch)+1)
		go func(w int8) { // 这种写法非常的危险！！！
			<-task(taskWriter(w))
			<-ws // 通知写者退出了
		}(w)
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

	T := task(taskRWMain(ch), func() { close(ch) })

	task := []int8{0}
	// task := []int8{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	for _, v := range task {
		ch <- v
	}

	close(ch)
	<-T

	time.Sleep(time.Second * 10)
}

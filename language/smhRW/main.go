package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SIZE 队列大小
const SIZE = 1 << 10

// RWMain 读者写者主程序
// readerFinished 管道看似多余，其实，这是一种通信管理方式
// 通过这种方式，我们将计数的操作都限定在主管理函数
func RWMain(ch chan int8, done chan struct{}) {
	rch := make(chan int8, SIZE)                 // 读者队列
	wch := make(chan int8, SIZE)                 // 写者队列
	ws := make(chan struct{}, 1)                 // 写程序排斥
	wg := &sync.WaitGroup{}                      // 正在读的读者计数器
	readerFinished := make(chan struct{}, 1<<31) // 读者完成队列
	defer func() {
		// 处理异常
		close(readerFinished)
		close(ws)
		close(wch)
		close(rch)
	}()

	go func() { // 读者读取完毕，需要通知计数减少
		for {
			select {
			case <-readerFinished:
				wg.Done()
			}
		}
	}()
	waitAllReaderExit := func() {
		wg.Wait()
	}
	waitWriterExit := func() {
		ws <- struct{}{} // 检查是否有写者
		<-ws
	}

	go generateReader(rch, readerFinished)
	go generateWriter(wch, ws)

	var rw int8
	for { // 分发任务
		select {
		case rw = <-ch:
			waitWriterExit()
			if rw == 0 {
				wg.Add(1)
				rch <- rw
			} else if rw == 1 {
				ws <- struct{}{}
				waitAllReaderExit() // 等待所有的读者退出
				wch <- rw
			} else {
				goto DONE
			}
		}
	}
DONE:
	waitAllReaderExit()
	done <- struct{}{}
}

// 生成读者
func generateReader(ch chan int8, wch chan struct{}) {
	var r int8
	for {
		select {
		case r = <-ch:
			fmt.Println("读者队列长度：", len(ch))
			go Reader(r, wch)
		}
	}
}

// 生成写者
func generateWriter(ch chan int8, ws chan struct{}) {
	var w int8
	for {
		select {
		case w = <-ch:
			fmt.Println("写者队列长度：", len(ch)+1)
			go Writer(w, ws)
		}
	}
}

// Reader 读者
func Reader(r int8, wch chan struct{}) {
	fmt.Println("读者正在读...", r)
	t := time.Second * time.Duration(rand.Intn(5)+2)
	fmt.Println("读时长：", t)
	time.Sleep(t)
	wch <- struct{}{}
}

// Writer 写者
func Writer(w int8, ws chan struct{}) {
	fmt.Println("写者正在写...", w)
	t := time.Second * time.Duration(rand.Intn(5)+2)
	fmt.Println("写时长：", t)
	time.Sleep(t)
	<-ws
}

func main() {
	ch := make(chan int8, SIZE) // 读写主程序的接收队列
	done := make(chan struct{})
	defer close(ch)
	defer close(done)

	go RWMain(ch, done)

	task := []int8{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1}
	for _, v := range task {
		ch <- v
	}

	<-done
}

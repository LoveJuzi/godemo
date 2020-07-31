package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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

func taskMPCMain(m, n int) func() {
	return func() {
		MPCMain(m, n)
	}
}

func taskProducerFactory(ch1 <-chan string, ch chan<- string, m int) func() {
	return func() {
		ProducerFactory(ch1, ch, m)
	}
}

func taskProducerMachine(ch1 <-chan string, ch chan<- string) func() {
	return func() {
		ProducerMachine(ch1, ch)
	}
}

func taskConsumerFactory(ch <-chan string, n int) func() {
	return func() {
		ConsumerFactory(ch, n)
	}
}

func taskConsumerMachine(ch <-chan string) func() {
	return func() {
		ConsumerMachine(ch)
	}
}

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 10

// MPCMain 主程序
func MPCMain(m, n int) {
	ch1 := make(chan string, SIZE)
	ch := make(chan string, SIZE)

	// 异常恢复代码，先下游后上游
	T1 := synctask(taskProducerFactory(ch1, ch, m), func() { close(ch) }, func() { close(ch1) })
	T2 := synctask(taskConsumerFactory(ch, n), func() { close(ch) })

	for i := 0; i < 100000; i++ { // 分发任务
		ch1 <- strconv.Itoa(i)
	}

	// 任务控制流
	close(ch1) // 通知T1任务结束
	<-T1       // 等待T1任务结束
	close(ch)  // 通知T2任务结束
	<-T2       // 等待T2任务结束
}

// ProducerFactory 用来启动生产者
func ProducerFactory(ch1 <-chan string, ch chan<- string, m int) {
	ch2 := make(chan string, m)
	T := make([]<-chan struct{}, m)
	for i := range T {
		T[i] = synctask(taskProducerMachine(ch2, ch), func() { close(ch) }, func() { close(ch2) })
	}
	for v := range ch1 { // 需要传递一下，因为close无法处理<-chan
		ch2 <- v
	}

	// 任务控制流
	close(ch2)         // 通知T任务结束
	for i := range T { // 等待T任务结束
		<-T[i]
	}
}

// ProducerMachine 每个Machine都会处理一次请求
func ProducerMachine(ch1 <-chan string, ch chan<- string) {
	for v := range ch1 {
		ch <- Producer(v)
	}
}

// ConsumerFactory 用来启动消费者
func ConsumerFactory(ch <-chan string, n int) {
	ch2 := make(chan string, n)
	T := make([]<-chan struct{}, n)
	for i := range T {
		T[i] = synctask(taskConsumerMachine(ch2), func() { close(ch2) })
	}
	for v := range ch { // 需要传递一下，因为close无法处理<-chan
		ch2 <- v
	}

	// 任务控制流
	close(ch2)         // 通知T任务结束
	for i := range T { // 等待T任务结束
		<-T[i]
	}
}

// ConsumerMachine 每个Machine都会处理一次请求
func ConsumerMachine(ch <-chan string) {
	for v := range ch {
		Consumer(v)
	}
}

// Producer 生产者加工
func Producer(src string) (dst string) {
	dst = src + "HHHHHHHH"
	return
}

// Consumer 消费者消费
func Consumer(s string) {
	fmt.Println(s)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)+1))
}

func main() {
	<-synctask(taskMPCMain(3, 10))
}

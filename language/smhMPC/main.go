package main

import (
	"fmt"
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

func main() {
	<-synctask(taskMPCMain(3, 10))
}

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 1

func taskMPCMain(m, n int) func() {
	return func() {
		MPCMain2(m, n)
	}
}

// MPCMain2 主程序
func MPCMain2(m, n int) {
	ch1 := make(chan string, SIZE)
	ch := make(chan string, SIZE)

	T1 := synctask(taskProducerFactory(ch1, ch, m), func() { close(ch1) }, func() { close(ch) })
	T2 := synctask(taskConsumerFactory(ch, n), func() { close(ch) })

	for i := 0; i < 1000; i++ { // 任务分发
		ch1 <- "aaa"
	}
	close(ch1)

	// 任务控制流
	<-T1      // 等待T1任务结束
	close(ch) // 通知T2任务结束
	<-T2      // 等待T2任务结束
}

func taskProducerFactory(ch1 <-chan string, ch chan<- string, m int) func() {
	return func() {
		ProducerFactory(ch1, ch, m)
	}
}

// ProducerFactory 用来启动生产者
func ProducerFactory(ch1 <-chan string, ch chan<- string, m int) {
	mch := make(chan struct{}, m) // 生产者个数
	for i := 0; i < m; i++ {
		mch <- struct{}{}
	}
	for v := range ch1 {
		if _, ok := <-mch; !ok {
			break
		}
		bgtask(taskProducer2(v, ch), mch, func() { close(ch) })
	}
}

func taskProducer2(src string, ch chan<- string) func() {
	return func() {
		ch <- Producer2(src)
	}
}

// Producer2 生产者加工
func Producer2(src string) (dst string) {
	dst = src + "HHHHHHHH"
	return
}

func taskConsumerFactory(ch <-chan string, n int) func() {
	return func() {
		ConsumerFactory(ch, n)
	}
}

// ConsumerFactory 用来启动消费者
func ConsumerFactory(ch <-chan string, n int) {
	nch := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		nch <- struct{}{}
	}
	for v := range ch {
		if _, ok := <-nch; !ok {
			break
		}
		bgtask(taskConsumer2(v), nch)
	}
}

func taskConsumer2(s string) func() {
	return func() {
		Consumer2(s)
	}
}

// Consumer2 消费者消费
func Consumer2(s string) {
	fmt.Println(s)
}

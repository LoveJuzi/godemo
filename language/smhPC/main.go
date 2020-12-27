package main

import "fmt"

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 4

func main() {
	ch := make(chan string, SIZE) // 管道通信

	T1 := task(taskProducer(ch)) // 启动生产任务
	T2 := task(taskConsumer(ch)) // 启动消费任务

	// 任务控制流
	<-T1      // 等待T1任务结束
	close(ch) // 通知T2任务结束
	<-T2      // 等待T2任务结束
}

type taskFun func()

func task(f taskFun) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }() // 退出信号
		defer func() {                        // 异常捕获
			if err := recover(); err != nil {
				// 记录异常日志
				fmt.Println(err)
			}
		}()

		f()
	}()
	return done
}

// 包装生产者任务
func taskProducer(ch chan<- string) taskFun {
	return func() {
		Producer(ch)
	}
}

// 包装消费者任务
func taskConsumer(ch <-chan string) taskFun {
	return func() {
		Consumer(ch)
	}
}

// Producer 生产者
func Producer(ch chan<- string) {
	fmt.Println("生产者开始生产...")
	for i := 0; i < 50; i++ {
		product := fmt.Sprintf("产品%d", i+1)
		fmt.Println("生产======>", product)
		ch <- product
	}
	fmt.Println("生产者结束生产...")
}

// Consumer 消费者
func Consumer(ch <-chan string) {
	fmt.Println("消费者开始消费...")
	for product := range ch { // ch关闭后，循环会退出
		fmt.Println("消费======>", product)
	}
	fmt.Println("消费者结束消费...")
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	MPCMain(3, 10) // 3个生产者，10个消费者
}

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 10

// MPCMain 多生产者消费者主程序
// m 表示生产者的个数，n表示消费者的个数
func MPCMain(m, n int) {
	ch := make(chan string, SIZE)

	T1 := []<-chan struct{}{}
	T2 := []<-chan struct{}{}
	for i := 0; i < m; i++ {
		T1 = append(T1, task(taskProducer(ch, i+1)))
	}
	for i := 0; i < n; i++ {
		T2 = append(T2, task(taskConsumer(ch, i+1)))
	}

	// 任务控制流
	for _, v := range T1 { // 等待T1任务结束
		<-v
	}
	close(ch)              // 通知T2任务结束
	for _, v := range T2 { // 等待T2任务结束
		<-v
	}
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

func taskProducer(ch chan<- string, id int) taskFun {
	return func() {
		Producer(ch, id)
	}
}

func taskConsumer(ch <-chan string, id int) taskFun {
	return func() {
		Consumer(ch, id)
	}
}

// Producer 生产者
func Producer(ch chan<- string, id int) {
	fmt.Printf("生产者%d开始生产...\n", id)
	for i := 0; i < 50; i++ {
		product := fmt.Sprintf("产品%d", i+1)
		fmt.Printf("生产%d======>%s\n", id, product)
		ch <- product
	}
	fmt.Printf("生产者%d结束生产...\n", id)
}

// Consumer 消费者
func Consumer(ch <-chan string, id int) {
	fmt.Printf("消费者%d开始消费...\n", id)
	for product := range ch {
		fmt.Printf("消费%d======>%s\n", id, product)
	}
	fmt.Printf("消费者%d结束消费...\n", id)
}

func t1(chs []<-chan string, id int) {
	ch1 := make(chan string)
	T1 := task(tasktt1(chs, ch1))
	T2 := task(taskConsumer(ch1, id))

	<-T1
	close(ch1)
	<-T2
}

func tasktt1(chs []<-chan string, ch1 chan<- string) taskFun {
	return func() {
		tt1(chs, ch1)
	}
}

// 将多channel数据放入ch1中
func tt1(chs []<-chan string, ch1 chan<- string) {
	T := []<-chan struct{}{}
	for _, ch := range chs {
		T = append(T, task(taskttt1(ch, ch1)))
	}
	for _, v := range T {
		<-v
	}
}

func taskttt1(ch <-chan string, ch1 chan<- string) taskFun {
	return func() {
		ttt1(ch, ch1)
	}
}

func ttt1(ch <-chan string, ch1 chan<- string) {
	for v := range ch {
		ch1 <- v
	}
}

func t2(chs []chan<- string, id int) {
	ch1 := make(chan string)

	wg := &sync.WaitGroup{}
	go func() {
		for _, ch := range chs {
			wg.Add(1)
			go func(ch chan<- string) {
				for v := range ch1 {
					ch <- v
				}
				wg.Done()
			}(ch)
		}
	}()

	Producer(ch1, id)
	close(ch1)
	wg.Wait() // 等待所有的生产数据写入到消息通道
}

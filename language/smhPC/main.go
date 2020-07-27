package main

import "fmt"

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 4

func main() {
	ch := make(chan string, SIZE)
	pchs := make(chan struct{})
	cchs := make(chan struct{})

	defer close(ch)   // 这个会最后一个释放
	defer close(pchs) // 这个会第二个释放
	defer close(cchs) // 这个会优先释放

	go Producer(ch, pchs) // 启动生产者
	go Consumer(ch, cchs) // 启动消费者

	<-pchs // 等待生产者生产完毕

	ch <- "" // 关闭消费者，需要等待生产者生产完毕

	<-cchs // 等待消费者消费完毕
}

// Producer 生产者
func Producer(ch chan string, chs chan struct{}) {
	fmt.Println("生产者开始生产...")
	for i := 0; i < 50; i++ {
		product := fmt.Sprintf("产品%d", i+1)
		fmt.Println("生产======>", product)
		ch <- product
	}
	fmt.Println("生产者结束生产...")
	chs <- struct{}{}
}

// Consumer 消费者
func Consumer(ch chan string, chs chan struct{}) {
	fmt.Println("消费者开始消费...")
	var product string
	for {
		select {
		case product = <-ch:
			if product == "" {
				goto DONE
			}
			fmt.Println("消费======>", product)
		}
	}
DONE:
	fmt.Println("消费者结束消费...")
	chs <- struct{}{}
}

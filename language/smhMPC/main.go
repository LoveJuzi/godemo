package main

import "fmt"

func test(A []struct{}) {
	fmt.Println(A[0])
}

func main() {
	A := [1 << 32]struct{}{}
	test(A[:])
	// fmt.Println(A)
	// MPCMain(3, 10) // 3个生产者，10个消费者
}

// SIZE 生产者消费者通道的大小
const SIZE int = 1 << 10

// MPCMain 多生产者消费者主程序
// m 表示生产者的个数，n表示消费者的个数
func MPCMain(m, n int) {
	ch := make(chan string, SIZE)
	pchs := make([]chan struct{}, m)
	cchs := make([]chan struct{}, n)

	for i := 0; i < m; i++ {
		pchs[i] = make(chan struct{})
	}
	for i := 0; i < n; i++ {
		cchs[i] = make(chan struct{})
	}

	defer close(ch)
	defer func() {
		for _, v := range pchs {
			close(v)
		}
	}()
	defer func() {
		for _, v := range cchs {
			close(v)
		}
	}()

	for idx, v := range pchs {
		go Producer(ch, v, idx) // 启动生产者
	}
	for idx, v := range cchs {
		go Consumer(ch, v, idx) // 启动消费者
	}

	for i := 0; i < m; i++ {
		<-pchs[i]
	}
	for i := 0; i < n; i++ {
		ch <- "" // 关闭消费者，需要等待生产者生产完毕
	}
	for i := 0; i < n; i++ {
		<-cchs[i]
	}
}

// Producer 生产者
func Producer(ch chan string, chs chan struct{}, id int) {
	defer func() {
		// 处理异常
		chs <- struct{}{}
	}()
	fmt.Printf("生产者%d开始生产...\n", id)
	for i := 0; i < 50; i++ {
		product := fmt.Sprintf("产品%d", i+1)
		fmt.Printf("生产%d======>%s\n", id, product)
		ch <- product
	}
	fmt.Printf("生产者%d结束生产...\n", id)
}

// Consumer 消费者
func Consumer(ch chan string, chs chan struct{}, id int) {
	defer func() {
		// 处理异常
		chs <- struct{}{}
	}()
	fmt.Printf("消费者%d开始消费...\n", id)
	var product string
	for {
		select {
		case product = <-ch:
			if product == "" {
				goto DONE
			}
			fmt.Printf("消费%d======>%s\n", id, product)
		}
	}
DONE:
	fmt.Printf("消费者%d结束消费...\n", id)
}

package main

import "fmt"

func main() {
	c := make(chan int)

	done := make(chan struct{})
	done2 := make(chan struct{})

	go func() {
		for v := range c {
			fmt.Println(v)
		}
		// 	for {
		// 		select {
		// 		case v, ok := <-c:
		// 			if !ok {
		// 				goto DONE
		// 			}
		// 			fmt.Println(v)
		// 		}
		// 	}
		// DONE:
		done <- struct{}{}
	}()

	go func() {
		for v := range c {
			fmt.Println(v)
		}
		// for {
		// 	select {
		// 	case v, ok := <-c:
		// 		if !ok {
		// 			goto DONE
		// 		}
		// 		fmt.Println(v)
		// 	}
		// }
		// DONE:
		done2 <- struct{}{}
	}()

	c <- 1
	c <- 2
	c <- 3

	close(c) // 发送信号，已经没有写数据了

	<-done
	<-done2
}

package main

import (
	"fmt"
)

func hello(ch chan int) {
	fmt.Println("Hello world")
	ch <- 0
}

func main() {
	ch := make(chan int)
	defer close(ch)
	go hello(ch)
	<-ch
}

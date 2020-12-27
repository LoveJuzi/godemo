package main

import (
	"fmt"
)

func safeClose(ch chan interface{}) {
	defer func() { recover() }()
	close(ch)
}

func main() {
	ch := make(chan string)
	T := task(func() {
		a := 2
		b := 2
		fmt.Println(0 / (a - b))
	}, ch)
	<-T
	T = task(func() {
		// time.Sleep(10 * time.Second)
	}, ch)
	<-T
}

func task(f func(), chs ...chan string) <-chan struct{} {
	d := make(chan struct{})
	go func() {
		defer func() { d <- struct{}{} }()
		defer recover() // 防止二次产生panic
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				// a := 2
				// b := 2
				// fmt.Println(0 / (a - b))
				for _, v := range chs {
					close(v)
					close(v)
				}
				// for _, v := range chs {
				// 	close(v)
				// }
			}

		}()
		f()
	}()
	return d
}

package main

import (
	"fmt"
)

// TQueue 线程安全队列
type TQueue struct {
	Data     []int
	Len      int
	beginIdx int
	endIdx   int

	popmutex  chan struct{}
	pushmutex chan struct{}
	fullch    chan struct{} //用于检查队列的满
	emptych   chan struct{} //用于检查队列的空
}

// NewTQueue 生成一个队列，需要指定队列的大小
func NewTQueue(l int) *TQueue {
	return &TQueue{
		Data:      make([]int, l+1),
		Len:       l + 1,
		beginIdx:  0,
		endIdx:    0,
		fullch:    make(chan struct{}, l),
		emptych:   make(chan struct{}, l),
		popmutex:  make(chan struct{}, 1),
		pushmutex: make(chan struct{}, 1),
	}
}

// RleaseTQueue 释放不必要的资源
func RleaseTQueue(tq *TQueue) {
	close(tq.fullch)
	close(tq.emptych)
	close(tq.popmutex)
	close(tq.pushmutex)
}

func idxplusplus(idx int, l int) int {
	return (idx + 1) % l
}

// IsEmpty 队列是否为空
func IsEmpty(tq *TQueue) bool {
	return tq.beginIdx == tq.endIdx
}

// IsFull 队列是否为满
func IsFull(tq *TQueue) bool {
	return tq.beginIdx == idxplusplus(tq.endIdx, tq.Len)
}

// Push 向队列中添加一个元素
// 当队列满的时候需要条件等待
func Push(tq *TQueue, e int) {
	tq.fullch <- struct{}{}
	defer func() {
		tq.emptych <- struct{}{}
	}()

	tq.pushmutex <- struct{}{}
	defer func() {
		<-tq.pushmutex
	}()

	tq.Data[tq.endIdx] = e
	tq.endIdx = idxplusplus(tq.endIdx, tq.Len)
}

// Pop 从队列中提取一个元素
// 当队列空时需要条件等待
func Pop(tq *TQueue) int {
	<-tq.emptych
	defer func() {
		<-tq.fullch
	}()

	tq.popmutex <- struct{}{}
	defer func() {
		<-tq.popmutex
	}()

	e := tq.Data[tq.beginIdx]
	tq.beginIdx = idxplusplus(tq.beginIdx, tq.Len)

	return e
}

func main() {
	done1 := make(chan struct{})
	done2 := make(chan struct{})
	defer close(done1)
	defer close(done2)
	tq := NewTQueue(5)
	defer RleaseTQueue(tq)

	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("======>", i)

			Push(tq, i)
		}
		done1 <- struct{}{}
	}()
	go func() {
		for {
			e := Pop(tq)
			fmt.Println(e)

			if e == 49 {
				break
			}
		}
		done2 <- struct{}{}
	}()

	<-done1
	<-done2
}

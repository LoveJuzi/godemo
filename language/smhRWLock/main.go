package main

import "sync"

// RWLock 读写锁
type RWLock struct {
	writer   bool
	writerch chan struct{}
	wg       *sync.WaitGroup
}

// NewRWLock 申请读写锁
func NewRWLock() *RWLock {
	return &RWLock{
		writer:   false,                  // 防止写操作饥饿
		writerch: make(chan struct{}, 1), // 写者互斥
		wg:       &sync.WaitGroup{},      // 读者计数器
	}
}

// RLock 读锁定
func RLock(rwlock *RWLock) {
	if rwlock.writer {
		rwlock.writerch <- struct{}{}
		<-rwlock.writerch
	}
	rwlock.wg.Add(1)
}

// RUnLock 读解锁
func RUnLock(rwlock *RWLock) {
	rwlock.wg.Done()
}

// Lock 写锁定
func Lock(rwlock *RWLock) {
	rwlock.writerch <- struct{}{}
	rwlock.writer = true
	rwlock.wg.Wait()
}

// UnLock 写解锁
func UnLock(rwlock *RWLock) {
	rwlock.writer = false
	<-rwlock.writerch
}

func main() {

}

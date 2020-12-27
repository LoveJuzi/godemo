package task

import "fmt"

// BgTask 后台任务
func BgTask(f func(), ech chan<- struct{}, rfs ...func()) {
	go func() {
		defer func() { recover() }()         // 捕获一切异常
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

// SyncTask 同步任务
func SyncTask(f func(), rfs ...func()) <-chan struct{} {
	d := make(chan struct{})
	go func() {
		defer func() { recover() }()       // 捕获一切异常
		defer func() { d <- struct{}{} }() // 发送退出信号
		ch := make(chan struct{})          // 后台任务退出信号
		BgTask(f, ch, rfs...)              // 后台处理任务
		<-ch                               // 等待后台任务结束
	}()
	return d
}

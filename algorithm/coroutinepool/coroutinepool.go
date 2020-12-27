package coroutinepool

/*
协程池：
功能：
	1. 需要提供一个申请函数，用于申请一个可运行的协程
	2. 需要提供一个释放函数，用户归还一个已经运行完毕的协程

注意事项：
	1. 需要考虑申请函数和释放函数的资源争抢问题
	2. 需要定义一个函数类型
*/

type job func()

// CoroutinePoolParam 池子的参数
type CoroutinePoolParam struct {
	Size int // 池子的大小
}

// CoroutinePool 协程池
type CoroutinePool struct {
	pool chan job
	size int
}

// InitPool 初始化池子
func InitCoroutinePool(param CoroutinePoolParam) *CoroutinePool {
	cp := &CoroutinePool{
		pool: make(chan job, param.Size),
		size: param.Size,
	}

	for i := 0; i < cp.size; i++ {
		go func() {
			for {
				f := <-cp.pool
				if f == nil {
					break
				}
				f()
			}
		}()
	}

	return cp
}

// Exec 在协程内部执行一个任务
func (cp *CoroutinePool) Exec(e job) {
	cp.pool <- e
}

// Release 释放所有的协程
func (cp *CoroutinePool) Release() {
	for i := 0; i < cp.size; i++ {
		cp.pool <- nil
	}
}

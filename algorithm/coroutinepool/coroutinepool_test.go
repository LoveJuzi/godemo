package coroutinepool

import (
	"fmt"
	"testing"
)

func TestPool(t *testing.T) {
	param := CoroutinePoolParam{Size: 10}
	pool := InitCoroutinePool(param)
	for i := 0; i < 20; i++ {
		a := i
		pool.Exec(func() {
			fmt.Println(a)
		})
	}
	pool.Release()
}

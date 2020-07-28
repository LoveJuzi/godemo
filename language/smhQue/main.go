package main

import (
	"fmt"
	"unsafe"
)

// BUFF 读者写着共享的区域
// 这个就是竞争条件
// 读者可以随便访问，当目前没有写着进入的时候
// 写者只能惟一进入，并且需要等到所有读者退出后，才能进行写入
var BUFF []string

// Reader 读者
func Reader() {

}

// Writer 写者
func Writer() {

}

func main() {
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2占中的字节数是 %d", n2, unsafe.Sizeof(n2))
}

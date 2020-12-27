package memory

// Node 一个有效的地址
type Node struct {
	a int
	b int
	c **int
	d [1 << 31]int
}

// ModifyA 修改 a 的值
func ModifyA(n Node) {
	n.a++
}

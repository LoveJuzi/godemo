package huffman

import (
	"container/heap"
)

// Interface 哈夫曼树的接口
type Interface interface {
	heap.Interface
}

// buildNode 使用原来的两个基础节点生成一个新的节点
// 若原来的节点是 x 和 y
// 生成的新的节点是 z
// 要求是 z 的做孩子是 x, z 的右孩子是 y
// z 的频率是 x 和 y 的频率之和
// Node 的一个可行的定义如下:
// 		type Node struct {
// 			ch byte         // 关键字
// 			f  float32      // 频率
// 			L  *Node        // 左孩子
// 			R  *Node        // 右孩子
// 		}
type buildNode func(interface{}, interface{}) interface{}

// Huffman 哈夫曼编码的算法过程
func Huffman(C Interface, xyz buildNode) interface{} {
	N := C.Len()

	heap.Init(C)

	for i := 1; i < N; i++ {
		heap.Push(C, xyz(heap.Pop(C), heap.Pop(C)))
	}

	return heap.Pop(C)
}

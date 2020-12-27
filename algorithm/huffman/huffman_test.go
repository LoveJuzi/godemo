package huffman

import (
	"fmt"
	"testing"
)

type Node struct {
	ch byte
	f  float32
	L  *Node
	R  *Node
}

func (p *Node) String() string {
	return fmt.Sprintf("{ch:%c, f:%f} ", p.ch, p.f)
}

type nodeHeap []*Node

func (h *nodeHeap) Len() int {
	return len(*h)
}

func (h *nodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *nodeHeap) Less(i, j int) bool {
	return (*h)[i].f < (*h)[j].f
}

func (h *nodeHeap) Push(v interface{}) {
	*h = append(*h, v.(*Node))
}

func (h *nodeHeap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *nodeHeap) String() string {
	rt := ""

	for _, v := range *h {
		rt += fmt.Sprintf("{ch:%c, f:%.2f} ", v.ch, v.f)
	}

	return rt
}

func buildZ(x, y interface{}) interface{} {
	return &Node{
		f: x.(*Node).f + y.(*Node).f,
		L: x.(*Node),
		R: y.(*Node),
	}
}

type HuffmanNode Node
type HuffmanTree *HuffmanNode

func checkHuffmanTreeEqual(a *HuffmanNode, b *HuffmanNode) bool {
	return true
}

func TestHuffman(t *testing.T) {
	C := []*Node{{ch: 'a', f: 0.50},
		{ch: 'c', f: 0.12},
		{ch: 'b', f: 0.20},
		{ch: 'e', f: 0.08},
		{ch: 'd', f: 0.10},
	}
	var NC nodeHeap = C
	huffmanTree := Huffman(&NC, buildZ).(*Node)
	checkHuffmanTreeEqual((*HuffmanNode)(huffmanTree), (*HuffmanNode)(huffmanTree))
}

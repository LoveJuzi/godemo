package leetcode590

// Node 定义
type Node struct {
    Val int
    Children []*Node
}

// Node2 遍历node的定义
type Node2 struct {
	p *Node
	l int
}

func postorder(root *Node) []int {
	T := []Node2{}
	if root !=nil {
		T = append(T, Node2{root, 0})
	}
	rt := []int{}

	for {
		if len(T) == 0 {
			break
		}
		node2 := T[len(T) - 1]
		if node2.l == len(node2.p.Children) {
			rt = append(rt, node2.p.Val)
			T = T[:len(T)-1]
		} else {
			childNode := node2.p.Children[node2.l]
			node2.l++
			T[len(T)-1] = node2
			T = append(T, Node2{childNode, 0})
		}
	}

	return rt
}

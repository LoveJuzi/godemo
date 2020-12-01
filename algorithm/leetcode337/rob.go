package leetcode337

// TreeNode 结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return f(root, buildHT())
}

func buildHT() map[*TreeNode]int {
	return make(map[*TreeNode]int)
}

func f(node *TreeNode, HT map[*TreeNode]int) int {
	if node == nil {
		return 0
	}

	if _, ok := HT[node]; ok {
		return HT[node]
	}

	var v1, v2, v3, v4 int

	var v5, v6 int

	v5 = f(node.Left, HT)
	v6 = f(node.Right, HT)

	if node.Left != nil {
		v1 = f(node.Left.Left, HT)
		v2 = f(node.Left.Right, HT)
	}

	if node.Right != nil {
		v3 = f(node.Right.Left, HT)
		v4 = f(node.Right.Right, HT)
	}

	if node.Val+v1+v2+v3+v4 > v5+v6 {
		HT[node] = node.Val + v1 + v2 + v3 + v4
	} else {
		HT[node] = v5 + v6
	}

	return HT[node]
}

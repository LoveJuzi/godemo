package leetcode872

// TreeNode 结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLeaf(node *TreeNode, T []int) []int {
	if node.Left == nil && node.Right == nil {
		T = append(T, node.Val)
		return T
	}

	if node.Left != nil {
		T = getLeaf(node.Left, T)
	}

	if node.Right != nil {
		T = getLeaf(node.Right, T)
	}

	return T
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	T1 := getLeaf(root1, []int{})
	T2 := getLeaf(root2, []int{})

	if len(T1) != len(T2) {
		return false
	}

	for i := range T1 {
		if T1[i] != T2[i] {
			return false
		}
	}

	return true
}

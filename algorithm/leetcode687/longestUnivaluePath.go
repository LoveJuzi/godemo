package leetcode687

// TreeNode 结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dc(node *TreeNode, T []int, k int) int {
	if node == nil {
		return 0
	}
	if k == node.Val {
		lpath := dc(node.Left, T, k)
		rpath := dc(node.Right, T, k)

		if T[0] < lpath+rpath {
			T[0] = lpath + rpath
		}

		if lpath > rpath {
			return lpath + 1
		}
		return rpath + 1
	}

	lpath := dc(node.Left, T, node.Val)
	rpath := dc(node.Right, T, node.Val)

	if T[0] < lpath+rpath {
		T[0] = lpath + rpath
	}

	return 0
}

func longestUnivaluePath(root *TreeNode) int {
	T := make([]int, 1)

	dc(root, T, 1<<31)

	return T[0]
}

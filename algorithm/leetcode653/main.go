package leetcode653

// TreeNode 定义Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return findTarget2(root, k, root)
}

func findTarget2(node *TreeNode, k int, root *TreeNode) bool {
	if node == nil {
		return false
	}
	if k-node.Val != node.Val && findTargetValue(root, k-node.Val) {
		return true
	}

	if findTarget2(node.Left, k, root) {
		return true
	}

	if findTarget2(node.Right, k, root) {
		return true
	}

	return false
}

func findTargetValue(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if root.Val == k {
		return true
	}
	if root.Val > k {
		return findTargetValue(root.Left, k)
	}
	return findTargetValue(root.Right, k)
}

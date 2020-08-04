package leetcode669

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	parent := &TreeNode{Left: root}

	db(parent, root, L, R)

	return parent.Left
}

func db(parent *TreeNode, node *TreeNode, L, R int) {
	if node == nil {
		return
	}
	if node.Val >= L && node.Val <= R {
		db(node, node.Left, L, R)
		db(node, node.Right, L, R)
		return
	}

	db(parent, trim(parent, node, L), L, R)

}

func trim(parent *TreeNode, node *TreeNode, L int) *TreeNode {
	LORR := parent.Left == node
	if node.Left == nil && node.Right == nil {
		if LORR {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return nil
	}

	if node.Left == nil || node.Val < L {
		if LORR {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		return node.Right
	}

	if LORR {
		parent.Left = node.Left
	} else {
		parent.Right = node.Left
	}
	return node.Left
}

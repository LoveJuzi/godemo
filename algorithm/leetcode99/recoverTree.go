package leetcode99

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode 结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	_, T := f(nil, root, buildT())
	if len(T) == 2 {
		T[0].Val, T[1].Val = T[1].Val, T[0].Val
	}

	if len(T) == 4 {
		T[0].Val, T[3].Val = T[3].Val, T[0].Val
	}
}

func buildT() []*TreeNode {
	return make([]*TreeNode, 0)
}

func f(preNode *TreeNode, node *TreeNode, T []*TreeNode) (*TreeNode, []*TreeNode) {
	if node == nil {
		return preNode, T
	}

	preNode, T = f(preNode, node.Left, T)
	if preNode != nil && preNode.Val > node.Val {
		T = append(T, preNode)
		T = append(T, node)
	}
	preNode = node
	preNode, T = f(preNode, node.Right, T)

	return preNode, T
}

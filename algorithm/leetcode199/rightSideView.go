package leetcode199

// TreeNode 结点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// 按层遍历就可以了
	Q := []*TreeNode{}

	res := []int{}

	if root != nil {
		Q = append(Q, root)
	}

	for {
		if len(Q) == 0 {
			break
		}
		res = append(res, Q[len(Q)-1].Val)
		QT := Q
		Q = []*TreeNode{}
		for i := 0; i < len(QT); i++ {
			if QT[i].Left != nil {
				Q = append(Q, QT[i].Left)
			}
			if QT[i].Right != nil {
				Q = append(Q, QT[i].Right)
			}
		}
	}

	return res
}

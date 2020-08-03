package leetcode988

// TreeNode 定义Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	s := f(root, []byte{})
	return string(s)
}

func f(node *TreeNode, s []byte) []byte {
	if node == nil {
		return s
	}

	s3 := []byte{byte('a') + byte(node.Val)}
	s3 = append(s3, s...)

	if node.Left == nil {
		return f(node.Right, s3)
	}
	if node.Right == nil {
		return f(node.Left, s3)
	}

	return compare(f(node.Left, s3), f(node.Right, s3))
}

func compare(s1 []byte, s2 []byte) []byte {
	i := 0
	j := 0

	for {
		if i >= len(s1) {
			return s1
		}
		if j >= len(s2) {
			return s2
		}
		if s1[i] < s2[j] {
			return s1
		}
		if s1[i] > s2[j] {
			return s2
		}
		i++
		j++
	}
}

package L20201201

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	return f(head, 1<<(getN(head)-1))
}

func getN(head *ListNode) int {
	var rt int
	for p := head; p != nil; p = p.Next {
		rt++
	}
	return rt
}

func f(node *ListNode, m int) int {
	if node == nil {
		return 0
	}

	return node.Val*m + f(node.Next, m>>1)
}

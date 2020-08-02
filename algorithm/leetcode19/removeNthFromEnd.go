package leetcode19

// ListNode node定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 { // 处理特殊情况
		return head
	}

	ch1 := make(chan *ListNode, 1)
	ch2 := make(chan *ListNode, 2)
	ch3 := make(chan *ListNode, 3)
	ch4 := make(chan *ListNode, 1)

	go func() {
		h := <-ch1
		p := h
		for i := 0; i < n; i++ {
			p = p.Next
		}
		ch2 <- p
		ch2 <- h
	}()

	go func() {
		p := <-ch2
		h := <-ch2
		var r *ListNode = nil
		q := h

		for ; p != nil; p = p.Next {
			r = q
			q = q.Next
		}

		ch3 <- q
		ch3 <- r
		ch3 <- h
	}()

	go func() {
		q := <-ch3
		r := <-ch3
		h := <-ch3

		if r == nil {
			ch4 <- q.Next
		} else {
			r.Next = q.Next
			ch4 <- h
		}
	}()

	ch1 <- head
	return <-ch4
}

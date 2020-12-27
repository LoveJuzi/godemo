package leetcode51

// 输入: n 是整数
// 输出: 所有的解，凡是空白部分，用“.”表示，凡是有皇后的部分用“Q”表示
// 条件: 不能有两个皇后在棋盘中的位置能够相遇，分别是横向，纵向，斜向
//
// 暴力解法：判定每种情况是否都满足条件，这里有个问题，如果计算每种情况，那么总共需要 n * n 次计算，才能完成任务
//
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	return allQueue(n, cal(0, n, buildT(n)))
}

// 需要一个函数能够返回所有的子结构的解，然后和当前的位置进行组合
// 这里我需要一个棋盘表，棋盘表空间也是一个问题，否则运算的时候就需要
// 返回所有的可能组合
func cal(m int, n int, T [][]int) [][]int {
	var rt [][]int

	rt = [][]int{}

	// 表示棋盘所有的可能性都校验了
	if m == n {
		return rt
	}

	for i := 0; i < n; i++ {
		if T[m][i] == 1 {
			continue
		}

		T = setT(T, m, i, 1)
		childRt := cal(m+1, n, T)
		T = setT(T, m, i, 0)

		rt = merge(rt, childRt, i)
	}

	return rt
}

//
func buildT(n int) [][]int {
	var rt [][]int

	rt = make([][]int, n)
	for i := 0; i < n; i++ {
		rt[i] = make([]int, n)
	}

	return rt
}

//
func setT(T [][]int, i, j int, flag int) [][]int {
	for ii := i + 1; ii < len(T); ii++ {
		T[ii][j] = flag
	}

	for ii, jj := i, j; ii < len(T) && jj >= 0; ii, jj = ii+1, jj-1 {
		T[ii][jj] = flag
	}

	for ii, jj := i, j; ii < len(T) && jj < len(T); ii, jj = ii+1, jj+1 {
		T[ii][jj] = flag
	}

	return T
}

func merge(a [][]int, b [][]int, m int) [][]int {
	if len(b) == 0 {
		return append(a, []int{m})
	}

	for i := 0; i < len(b); i++ {
		a = append(a, append(b[i], m))
	}
	return a
}

func dot() string {
	return "."
}

func queue() string {
	return "Q"
}

// 一行的放置
// 其中：n 表示棋盘长度，m 表示在第 m 个位置放置一个皇后，m 的起始标签为0
func lineQueue(n int, m int) string {
	var rt string

	for i := 0; i < n; i++ {
		if i == m {
			rt += queue()
		} else {
			rt += dot()
		}
	}

	return rt
}

// 棋盘排列方式
func sigleQueue(n int, T []int) []string {
	var rt []string

	rt = []string{}

	for i := 0; i < n; i++ {
		rt = append(rt, lineQueue(n, i))
	}

	return rt
}

//
func allQueue(n int, T [][]int) [][]string {
	var rt [][]string

	rt = [][]string{}

	for i := 0; i < len(T); i++ {
		rt = append(rt, sigleQueue(n, T[i]))
	}

	return rt
}

package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	if k > 9 {
		return [][]int{}
	}

	rt := [][]int{}

	T := make([]int, k)
	ok := true
	j := 0
	ok = updateT(T, 0, 1)
	for {
		s := sumT(T)
		if ok && s < n {
			j = k - 1
			ok = updateT(T, j, T[j]+1)
			continue
		}
		if ok && s == n {
			// 一个可行的解
			T1 := make([]int, k)
			copy(T1, T)
			rt = append(rt, T1)
			ok = updateT(T, j, T[j]+1)
			continue
		}

		if j == 0 {
			break
		}

		j = j - 1
		ok = updateT(T, j, T[j]+1)
	}

	return rt
}

func updateT(T []int, idx int, n int) bool {
	for i := idx; i < len(T); i++ {
		T[i] = n
		if T[i] > 9 {
			return false
		}
		n++
	}
	return true
}

func sumT(T []int) int {
	var rt int
	for _, v := range T {
		rt += v
	}
	return rt
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}

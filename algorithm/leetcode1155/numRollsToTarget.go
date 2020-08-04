package main

import "fmt"

func main() {
	fmt.Println(numRollsToTarget(30, 30, 500))
}

func numRollsToTarget(d int, f int, target int) int {
	T := make([]map[int]int, d+1)
	for idx := range T {
		T[idx] = map[int]int{}
	}
	for i := 1; i <= f; i++ {
		T[1][i] = 1
	}

	rt := db(T, d, f, target)

	return rt
}

func db(T []map[int]int, d int, f int, target int) int {
	var mod int = 1000000007
	if v, ok := T[d][target]; ok { // direct return if the child problem has been resolved
		return v
	}

	if d == 1 { // empty child problem
		return 0
	}

	rt := 0
	for i := 1; i <= f; i++ { // find all child problem
		rt += db(T, d-1, f, target-i) // resovle child problem
	}

	T[d][target] = rt % mod

	return T[d][target]
}

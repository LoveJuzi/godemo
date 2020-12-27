package leetcode1346

func checkIfExist(arr []int) bool {
	N := len(arr)
	T := map[int]int{}
	for i := 0; i < N; i++ {
		if _, ok := T[arr[i]]; ok {
			return true
		}
		T[arr[i]<<1] = 1
		if arr[i]%2 == 0 {
			T[arr[i]>>1] = 1
		}
	}

	return false
}

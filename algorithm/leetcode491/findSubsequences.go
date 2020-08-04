package leetcode491

func findSubsequences(nums []int) [][]int {
	return f([]int{}, nums)
}

func f(h []int, o []int) [][]int {
	rt := [][]int{}

	seen := map[int]struct{}{}
	for i := 0; i < len(o); i++ { // find all child problem
		if _, ok := seen[o[i]]; ok { // empty child problem (duplicates)
			continue
		}
		seen[o[i]] = struct{}{}

		/* begin: resovle child problem */
		if len(h) != 0 && h[len(h)-1] > o[i] {
			continue
		}
		nh := []int{}
		nh = append(nh, h...)
		nh = append(nh, o[i])
		if len(h) != 0 {
			rt = append(rt, nh)
		}

		rt = append(rt, f(nh, o[i+1:])...)
		/*end: resolve child problem*/
	}

	return rt
}

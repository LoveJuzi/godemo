package leetcode1128

func numEquivDominoPairs(dominoes [][]int) int {
	rt := 0
	seen := make([]int, len(dominoes))
	for i := 0; i < len(dominoes); i++ { // find all child problem

		if seen[i] == 1 {
			continue
		}
		// resolve child problem
		seen[i] = 1
		cnt := 0
		deta := 1
		for j := i + 1; j < len(dominoes); j++ {
			if compare(dominoes, i, j) {
				cnt += deta
				deta++
				seen[j] = 1
			}
		}
		rt += cnt
	}

	return rt
}

func compare(d [][]int, i, j int) bool {
	if d[i][0] == d[j][0] && d[i][1] == d[j][1] {
		return true
	}
	if d[i][0] == d[j][1] && d[i][1] == d[j][0] {
		return true
	}
	return false
}

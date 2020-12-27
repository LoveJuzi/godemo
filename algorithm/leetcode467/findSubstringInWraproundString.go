package leetcode467

func findSubstringInWraproundString(p string) int {
	T := map[byte]int{}
	CHN := 26

	for i := 0; i < CHN; i++ {
		T[byte(i)+'a'] = 0
	}

	N := len(p)

	rt := 0
	d := 1
	nextCh := byte(0)
	i := 0
	for i < N {
		if nextCh == byte(0) || nextCh == p[i] {
			rt += d
			if T[p[i]] >= d {
				rt -= d
			} else {
				rt -= T[p[i]]
				T[p[i]] = d // 更新距离
			}
			nextCh = byte((int(p[i]-'a')+1)%CHN) + 'a'
			d++
			i++
		} else {
			nextCh = byte(0)
			d = 1
			continue
		}
	}

	return rt
}

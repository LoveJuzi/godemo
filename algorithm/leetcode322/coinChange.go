package leetcode322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	return f(coins, amount, map[int]int{})
}

func f(coins []int, amount int, T map[int]int) int {
	if L, ok := T[amount]; ok {
		return L
	}

	minL := -1
	for i := 0; i < len(coins); i++ {
		if amount-coins[i] < 0 {
			continue
		}
		tmp := 1
		if amount-coins[i] > 0 {
			tmp = f(coins, amount-coins[i], T) + 1
		}
		if tmp != 0 {
			if minL == -1 {
				minL = tmp
			} else if minL > tmp {
				minL = tmp
			}
		}
	}

	T[amount] = minL
	return T[amount]
}

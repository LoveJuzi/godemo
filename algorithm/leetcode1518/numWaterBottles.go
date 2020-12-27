package leetcode1518

func numWaterBottles(numBottles int, numExchange int) int {
	remainder := 0
	rt := 0
	for numBottles > 0 {
		rt += numBottles
		numBottles += remainder
		remainder = numBottles % numExchange
		numBottles = numBottles / numExchange
	}

	return rt
}

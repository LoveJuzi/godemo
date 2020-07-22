package leetcode605

// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	c := 1
// 	N := len(flowerbed)
// 	for i := 0; i < N && n > 0; i++ {
// 		if flowerbed[i] == 1 {
// 			c = 0
// 			continue
// 		}
// 		c++
// 		if c == 3 {
// 			c = 1
// 			n--
// 		}
// 	}

// 	if n != 0 && c == 2 {
// 		n--
// 	}

// 	return n == 0
// }

func canPlaceFlowers(flowerbed []int, n int) bool {
	c := 1
	N := len(flowerbed)
	m := 0
	for i := 0; i < N; i++ {
		if flowerbed[i] == 1 {
			c = 0
			continue
		}
		c++
		if c == 3 {
			c = 1
			m++
		}
	}

	if c == 2 {
		m++
	}

	return n <= m
}

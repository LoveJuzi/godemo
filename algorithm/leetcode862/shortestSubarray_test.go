package leetcode862

import "testing"

func Test_shortestSubarray(t *testing.T) {
	cases := []struct {
		A        []int
		K        int
		expected int
	}{
		// {
		// 	[]int{27, 20, 79, 87, -36, 78, 76, 72, 50, -26}, 453, 9,
		// },
		// {
		// 	[]int{77, 19, 35, 10, -14}, 19, 1,
		// },
		// {
		// 	[]int{86396, 74204, 24861, 72405, 30809, 40710, 47892, -48882, -9084, 59464, 29389, 1510, 16521, 38996, 98830, 15183, 38241, 90465, -10717, 81061, -40387, -23424, 74146, -24051, 56847, 44278, 41403, -763, 50836, 6482, 44225, 16178, -48529, -36193, 28857, -16654, 48188, 54971, -29822, 25959, 90144, -23182, -9464, 65609, 99248, -26248, 47993, -20085, 75072, 70400}, 209110, 4,
		// },
		{
			[]int{39353, 64606, -23508, 5678, -17612, 40217, 15351, -12613, -37037, 64183, 68965, -19778, -41764, -21512, 17700, -23100, 77370, 64076, 53385, 30915, 18025, 17577, 10658, 77805, 56466, -2947, 29423, 50001, 31803, 9888, 71251, -6466, 77254, -30515, 2903, 76974, -49661, -10089, 66626, -7065, -46652, 84755, -37843, -5067, 67963, 92475, 15340, 15212, 54320, -5286}, 207007, 4,
		},
	}

	for _, c := range cases {
		result := shortestSubarray(c.A, c.K)
		if result != c.expected {
			t.Fatal("error")
		}
	}
}
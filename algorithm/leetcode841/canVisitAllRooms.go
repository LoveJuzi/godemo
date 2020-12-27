package leetcode841

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	T := make([]int, len(rooms))
	T[0] = 0
	cnt := 0

	Q := []int{}
	Q = append(Q, 0)

	for {
		if len(Q) == 0 {
			break
		}
		n := Q[len(Q)-1]
		Q = Q[:len(Q)-1]
		if T[n] == 0 {
			Q = append(Q, rooms[n]...)
			T[n] = 1
			cnt++
		}
	}

	return cnt == len(T)
}

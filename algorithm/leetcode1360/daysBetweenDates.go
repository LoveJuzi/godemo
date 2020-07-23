package leetcode1360

import (
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", date1+" 12:00:00", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", date2+" 12:00:00", time.Local)
	if t1.UnixNano() < t2.UnixNano() {
		t1, t2 = t2, t1
	}

	return int(t1.Sub(t2).Hours()) / 24
}

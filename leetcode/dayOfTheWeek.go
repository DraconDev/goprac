package leetcode

import "time"

func dayOfTheWeek(day int, month int, year int) string {
	mon := time.Month(month)
	part_time := time.Date(year, mon, day, 0, 0, 0, 0, time.UTC)

	return part_time.Weekday().String()
}

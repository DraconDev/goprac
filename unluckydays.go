package main

import "time"

func UnluckyDays(year int) int {

	unlucky := 0
	for m := 1; m <= 12; m++ {
		if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
			unlucky++
		}
	}
	return unlucky
}

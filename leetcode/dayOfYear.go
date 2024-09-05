package leetcode

import "strconv"

func dayOfYear(date string) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	result := 0

	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	for i := 0; i < month-1; i++ {
		result += months[i]
	}
	day, _ := strconv.Atoi(date[8:10])

	if month > 2 && (year%4 == 0 && year%100 != 0 || year%400 == 0) {
		result++
	}
	return result + day

}

package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	diff := 0
	for i, v := range seats {
		diff += abs(v - students[i])
	}
	return diff
}

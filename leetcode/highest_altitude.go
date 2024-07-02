package leetcode

func largestAltitude(gain []int) int {
	highest, current := 0, 0
	for _, v := range gain {
		current += v
		if current > highest {
			highest = current
		}
	}
	return highest
}

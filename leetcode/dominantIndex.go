package leetcode

func dominantIndex(nums []int) int {
	highest := 0
	highestIndex := 0
	secondHighest := 0

	for i, v := range nums {
		if v > highest {
			secondHighest = highest
			highest = v
			highestIndex = i
		} else if v > secondHighest && v != highest {
			secondHighest = v
		}
	}

	if highest >= secondHighest*2 {
		return highestIndex
	}
	return -1
}

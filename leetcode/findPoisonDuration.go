package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	var poisonDuration int

	for i, v := range timeSeries[:len(timeSeries)-1] {
		poisonDuration += min(timeSeries[i+1]-v, duration)
	}
	poisonDuration += duration

	return poisonDuration
}

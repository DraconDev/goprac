package leetcode

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	for i, v := range temperatures {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > v {
				days[i] = j - i
				break
			}
		}
	}
	return days

}

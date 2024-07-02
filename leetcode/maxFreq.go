package leetcode

func MaxFrequencyElements(nums []int) int {

	var maxFreq int
	count := make(map[int]int)

	for _, v := range nums {
		count[v]++
		if count[v] > maxFreq {
			maxFreq = count[v]
		}
	}

	var result int

	for _, v := range count {
		if v == maxFreq {
			result += v
		}
	}

	return result
}

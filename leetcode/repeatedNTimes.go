package leetcode

func repeatedNTimes(nums []int) int {
	numberMap := make(map[int]int)

	for _, v := range nums {
		numberMap[v]++
		if numberMap[v] == len(nums)/2 {
			return v
		}
	}
	return 0
}

package leetcode

func FindErrorNums(nums []int) []int {

	exists := make(map[int]bool)

	result := make([]int, 2)

	for _, v := range nums {
		if exists[v] {
			result[0] = v
		}
		exists[v] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !exists[i] {
			result[1] = i
			break
		}
	}

	return result
}

package leetcode

func findDuplicate(nums []int) int {
	var exists = make(map[int]bool)

	for _, v := range nums {
		if exists[v] {
			return v
		}
		exists[v] = true
	}
	return 0
}
